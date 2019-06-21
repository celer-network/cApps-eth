pragma solidity ^0.5.0;
import "../templates/SingleSessionBooleanOutcome.sol";

contract SingleGomoku is SingleSessionBooleanOutcome {

    enum StateKey {Turn, Winner, FullState}

    uint8 constant BOARD_DIMENSION = 15; // board dimension is 15x15

    bytes private boardState;       // 227 bytes: uint8 winner + uint8 turn + 15x15 bytes board
    uint16 private stoneNum;        // number of stones
    uint16 private stoneNumOnchain; // number of stones placed on-chain
    uint8 private minStoneOffchain; // minimal number of stones before go onchain
    uint8 private maxStoneOnchain;  // maximal number of stones after go onchain

    constructor(
        address[] memory _players,
        uint _nonce,
        uint _timeout,
        uint8 _minStoneOffchain,
        uint8 _maxStoneOnchain)
        public
        SingleSessionBooleanOutcome(_players, _nonce, _timeout)
    {
        require(_players.length == 2, "invalid number of players");
        minStoneOffchain = _minStoneOffchain;
        maxStoneOnchain = _maxStoneOnchain;
        boardState = new bytes(227);
    }

    /**
     * @notice Get the app outcome
     * @param _query Query arg, player id (1 or 2)
     * @return True if the player wins
     */
    function getOutcome(bytes calldata _query) external view returns (bool) {
        require(_query.length == 1, "invalid query length");
        return boardState[0] == _query[0];
    }

    /**
     * @notice Get app state
     * @param _key Query key 0:Turn, 1:Winner, 2:FullState
     * @return App session state: Turn (1 byte), Winner (1 byte), FllState (227 bytes)
     */
    function getState(uint _key) external view returns (bytes memory) {
        if (_key == uint(StateKey.Winner)) {
            bytes memory b = new bytes(1);
            b[0] = boardState[0];
            return b;
        } else if (_key == uint(StateKey.Turn)) {
            bytes memory b = new bytes(1);
            b[0] =boardState[1];
            return b;
        }
        return boardState;
    }

    /**
     * @notice Update state according to an off-chain state proof
     * @dev state co-signed by all players, so skeip thorough check for state correctness
     * @param _state Signed off-chain app state
     * @return True if update succeeds
     */
    function updateByState(bytes memory _state) internal returns (bool) {
        // uint8 winner + uint8 turn + 15x15 bytes board
        require(_state.length == 227, "invalid state length");
        if(uint8(_state[0]) != 0) {
            winGame(uint8(_state[0]));
        } else {
            // load other states only if winner is not specified
            uint count = 0;
            for (uint i = 2; i < 227; i++) {
                if (uint8(_state[i]) != 0) {
                    count++;
                }
            }
            stoneNum = uint16(count);
            require(stoneNum >= minStoneOffchain, "not enough offchan stones");
        }
        boardState = _state;
        return true;
    }

    /**
     * @notice Update state according to an on-chain action
     * @param _action Action data to place stone, 2 bytes for coordinate (x, y)
     * @return True if update succeeds
     */
    function updateByAction(bytes memory _action) internal returns (bool) {
        uint8 turn = uint8(boardState[1]);
        require(msg.sender == appInfo.players[turn-1], "not your turn");
        require(_action.length == 2, "invalid action length");
        uint8 x = uint8(_action[0]);
        uint8 y = uint8(_action[1]);
        require(checkBoundary(x, y), "out of boundary");
        uint index = stateIndex(x, y);
        require(uint8(boardState[index]) == 0, "slot is occupied");

        // place the stone
        boardState[index] = byte(turn);
        stoneNum++;
        stoneNumOnchain++;
        bytes memory board = boardState;

        // check if there is five-in-a-row including this new stone
        if (checkFive(board, x, y, 1, 0) ||  // horizaontal bidirection
            checkFive(board, x, y, 0, 1) ||  // vertical bidirection
            checkFive(board, x, y, 1, 1) ||  // main-diagonal bidirection
            checkFive(board, x, y, 1, -1)) { // anti-diagonal bidirection
            winGame(turn); // we have a winner
            return true;
        }

        if (stoneNum == 225 || stoneNumOnchain > maxStoneOnchain) {
            // all slots occupied, game is over with no winner
            appInfo.status = AppStatus.FINALIZED;
            setTurn(0);
        } else {
            // toggle turn and update game phase
            if (turn == 1) {
                setTurn(2);
            } else {
                setTurn(1);
            }
        }
        return true;
    }

    /**
     * @notice Finalized based on current state in case of on-chain action timeout
     */
    function finalizeOnTimeout() internal {
        if (uint8(boardState[1]) == 1) {
            winGame(2);
        } else if (uint8(boardState[1]) == 2) {
            winGame(1);
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
        return 2 + BOARD_DIMENSION * _x + _y;
    }

    /**
     * @notice Set game states when there is a winner
     * @param _winner is id of the winner
     */
    function winGame(uint8 _winner) private {
        require(0 <= _winner && _winner <= 2, "invalid winner state");
        setWinner(_winner);
        // Game over
        if (_winner != 0) {
            setTurn(0);
            appInfo.status = AppStatus.FINALIZED;
        }
    }

    function setWinner(uint8 _winner) private {
        boardState[0] = byte(_winner);
    }

    function setTurn(uint8 _turn) private {
        boardState[1] = byte(_turn);
    }

}
