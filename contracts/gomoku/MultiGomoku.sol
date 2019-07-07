pragma solidity ^0.5.0;
import "../templates/MultiSessionBooleanOutcome.sol";

contract MultiGomoku is MultiSessionBooleanOutcome {

    enum StateKey {TurnColor, WinnerColor, FullState}

    uint8 constant BLACK = 1;
    uint8 constant WHITE = 2;
    uint8 constant BOARD_DIMENSION = 15; // board dimension is 15x15
    uint constant STATE_LENGTH = 228; // length of board state

    uint8 public minStoneOffchain; // minimal number of stones before go onchain
    uint8 public maxStoneOnchain;  // maximal number of stones after go onchain

    struct GameInfo {
        // 228 bytes: uint8 winner color + uint8 turn color + uint8 black id + 15x15 bytes board
        bytes boardState;
        // number of stones
        uint16 stoneNum;
        // number of stones placed on-chain
        uint16 stoneNumOnchain;
    }

    mapping(bytes32 => GameInfo) private gameInfoMap; // session id -> game info

    constructor(uint8 _minStoneOffchain, uint8 _maxStoneOnchain)
        public MultiSessionBooleanOutcome(2)
    {
        minStoneOffchain = _minStoneOffchain;
        maxStoneOnchain = _maxStoneOnchain;
    }

    /**
     * @notice Get the app outcome
     * @param _query Query arg, player id (1 or 2)
     * @return True if the player wins
     */
    function getOutcome(bytes32 _session, bytes memory _query)
        internal
        view
        returns (bool)
    {
        require(_query.length == 1, "invalid query length");
        return gameInfoMap[_session].boardState[0] == _query[0];
    }

    /**
     * @notice Get app state
     * @param _session App session ID
     * @param _key Query key, 0:TurnColor, 1:WinnerColor, 2:FullState
     * @return App session state: TurnColor (1 byte), WinnerColor (1 byte), FllState (228 bytes)
     */
    function getState(bytes32 _session, uint _key) external view returns (bytes memory) {
        if (_key == uint(StateKey.WinnerColor)) {
            bytes memory b = new bytes(1);
            b[0] = gameInfoMap[_session].boardState[0];
            return b;
        } else if (_key == uint(StateKey.TurnColor)) {
            bytes memory b = new bytes(1);
            b[0] = gameInfoMap[_session].boardState[1];
            return b;
        }
        return gameInfoMap[_session].boardState;
    }

    /**
     * @notice Update on-chain state according to off-chain state proof
     * @param _session Session ID
     * @param _state Signed off-chain state
     */
    function updateByState(bytes32 _session, bytes memory _state)
        internal
        returns (bool)
    {
        // uint8 winner color + uint8 turn color + uint8 black id + 15x15 bytes board
        require(_state.length == STATE_LENGTH, "invalid state length");
        GameInfo storage game = gameInfoMap[_session];
        if (game.boardState.length == 0) {
            game.boardState = new bytes(STATE_LENGTH);
        }
        if(uint8(_state[0]) != 0) {
            winGame(_session, uint8(_state[0]));
        } else {
            // load other states only if winner color is not BLACK or WHITE
            uint count = 0;
            for (uint i = 3; i < STATE_LENGTH; i++) {
                if (uint8(_state[i]) != 0) {
                    count++;
                }
            }
            game.stoneNum = uint16(count);
            require(game.stoneNum >= minStoneOffchain, "not enough offchan stones");
        }
        game.boardState = _state;
        return true;
    }

    /**
     * @notice Update state according to an on-chain action
     * @param _session App session
     * @param _action Action data to place stone, 2 bytes for coordinate (x, y)
     * @return True if update succeeds
     */
    function updateByAction(bytes32 _session, bytes memory _action)
        internal
        returns (bool)
    {
        uint8 turnColor = getTurnColor(_session);
        // black player index, smaller (=1) or larger (=2) addr
        uint8 blackID = getBlackPlayerID(_session);
        if(blackID == 1) {
            require(msg.sender == sessionInfoMap[_session].players[turnColor-1], "Not your turn");
        } else if(blackID == 2) {
            require(msg.sender == sessionInfoMap[_session].players[2-turnColor], "Not your turn");
        } else {
            assert(false);
        }
        require(_action.length == 2, "invalid action length");
        uint8 x = uint8(_action[0]);
        uint8 y = uint8(_action[1]);
        require(checkBoundary(x, y), "out of boundary");
        GameInfo storage game = gameInfoMap[_session];
        uint index = stateIndex(x, y);
        require(uint8(game.boardState[index]) == 0, "slot is occupied");

        // place the stone
        game.boardState[index] = byte(turnColor);
        game.stoneNum++;
        game.stoneNumOnchain++;
        bytes memory board = game.boardState;

        // check if there is five-in-a-row including this new stone
        if (checkFive(board, x, y, 1, 0) ||  // horizaontal bidirection
            checkFive(board, x, y, 0, 1) ||  // vertical bidirection
            checkFive(board, x, y, 1, 1) ||  // main-diagonal bidirection
            checkFive(board, x, y, 1, -1)) { // anti-diagonal bidirection
            winGame(_session, turnColor); // we have a winner
            return true;
        }

        if (game.stoneNum == 225 || game.stoneNumOnchain > maxStoneOnchain) {
            // all slots occupied, game is over with no winner
            sessionInfoMap[_session].status = SessionStatus.FINALIZED;
            setTurnColor(_session, 0);
        } else {
            // toggle turn and update game phase
            if (turnColor == BLACK) {
                setTurnColor(_session, WHITE);
            } else {
                setTurnColor(_session, BLACK);
            }
        }
        return true;
    }

    /**
     * @notice Finalize the session based on current state in case of on-chain action timeout
     * @param _session Session ID
     */
    function finalizeOnTimeout(bytes32 _session) internal  {
        if (getTurnColor(_session) == BLACK) {
            winGame(_session, WHITE);
        } else if (getTurnColor(_session) == WHITE) {
            winGame(_session, BLACK);
        }
    }

    /**
     * @notice Check if there is five in a row in a given direction
     * @param _x is x coordinate on the board
     * @param _y is y coordinate on the board
     * @param _xdir is direction (-1 or 0 or 1) in x axis
     * @param _ydir is direction (-1 or 0 or 1) in y axis
     * @return true if there is five in a row
     */
    function checkFive(bytes memory _board, uint8 _x, uint8 _y, int _xdir, int _ydir)
        private
        pure
        returns (bool)
    {
        uint8 count = 0;
        count += countStone(_board, _x, _y, _xdir, _ydir);
        count += countStone(_board, _x, _y, -1 * _xdir, -1 * _ydir) - 1; // reverse direction
        if (count >= 5) {
            return true;
        }
        return false;
    }

    /**
     * @notice Count the maximum consecutive stones in a given direction
     * @param _x is x coordinate on the board
     * @param _y is y coordinate on the board
     * @param _xdir is direction (-1 or 0 or 1) in x axis
     * @param _ydir is direction (-1 or 0 or 1) in y axis
     * @return the number of maximum consecutive stones in a given direction
     */
    function countStone(bytes memory _board, uint8 _x, uint8 _y, int _xdir, int _ydir)
        private
        pure
        returns (uint8)
    {
        uint8 count = 1;
        while (count <= 5) {
            uint8 x = uint8(int8(_x) + _xdir * count);
            uint8 y = uint8(int8(_y) + _ydir * count);
            if (checkBoundary(x, y) && (_board[stateIndex(x, y)] == _board[stateIndex(_x, _y)])) {
                count += 1;
            } else {
                return count;
            }
        }
    }

    /**
     * @notice Check if coordinate (x, y) is valid
     * @param _x is x coordinate on the board
     * @param _y is y coordinate on the board
     * @return true is (x, y) is a valid coordinate
     */
    function checkBoundary(uint8 _x, uint8 _y) private pure returns (bool) {
        return (_x < BOARD_DIMENSION && _y < BOARD_DIMENSION);
    }

    /**
     * @notice Check if coordinate (x, y) is valid
     * @param _x is x coordinate on the board
     * @param _y is y coordinate on the board
     * @return index of the state
     */
    function stateIndex(uint8 _x, uint8 _y) private pure returns (uint) {
        return 3 + BOARD_DIMENSION * _x + _y;
    }

    /**
     * @notice Set game states when there is a winner
     */
    function winGame(bytes32 _session, uint8 _winner) private {
        require(0 <= _winner && _winner <= 2, "invalid winner state");
        setWinnerColor(_session, _winner);
        // Game over
        if (_winner != 0) {
            setTurnColor(_session, 0);
            sessionInfoMap[_session].status = SessionStatus.FINALIZED;
        }
    }

    function setWinnerColor(bytes32 _session, uint8 _winner) private {
        gameInfoMap[_session].boardState[0] = byte(_winner);
    }

    function setTurnColor(bytes32 _session, uint8 _turnColor) private {
        gameInfoMap[_session].boardState[1] = byte(_turnColor);
    }

    function getTurnColor(bytes32 _session) private view returns (uint8){
        return uint8(gameInfoMap[_session].boardState[1]);
    }

    function getBlackPlayerID(bytes32 _session) private view returns (uint8){
        return uint8(gameInfoMap[_session].boardState[2]);
    }
}
