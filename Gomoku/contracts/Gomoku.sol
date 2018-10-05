pragma solidity ^0.4.23;
import "./lib/chain.pb.sol";

contract Gomoku {
    enum GamePhase {start, settle, play, over}
    GamePhase public game_phase;          // phase of the game

    address[2] public players;            // addresses of player1 (black) and player2 (white)

    uint8 constant board_dimmension = 15; // board dimmension is 15x15
    uint8[board_dimmension][board_dimmension] public board; // game board
    
    uint public nonce;                    // nonce of the board state
    uint8 public winner;                  // who winned? 0: nobody; 1: player1; 2: player2
    uint8 public turn;                    // whose turn? 0: nobody; 1: player1; 2: player2
 
    uint public finalized_time;           // block number when the game is over
    uint public move_timeout;             // timeout to make an on-chain move
    uint public move_deadline;            // block.number + move_timeout
    uint public settle_timeout;           // timeout to submit state proof during settle phase
    uint public settle_deadline;          // block.number + settle_deadline

    event IntendSettle(uint nonce);       // emit event when an off-chain state is submitted
    event ConfirmSettle();                // emit event when an off-chain state is settled
    event PlaceStone(uint8 x, uint8 y);   // emit event when a stone is placed on-chain
    event GameOver(uint winner);          // emit event when game is over

    uint16 private stone_num;             // number of stones
    uint16 private stone_num_onchain;     // number of stones placed on-chain
    uint8 private min_stone_offchain;     // minimal number of stones before go onchain
    uint8 private max_stone_onchain;      // maximal number of stones after go onchain

    /**
    @notice Gomoku constructor
    @param _player1 is address of player1 (black)
    @param _player2 is address of player2 (white)
    @param _move_timeout is timeout to make an on-chain move
    @param _settle_timeout is timeout to submit state proofs during settle phase
    @param _min_stone_offchain is minimal number of stones before go onchain
    @param _max_stone_onchain is maximal number of stones after go onchain
    */
    constructor(
        address _player1, 
        address _player2, 
        uint _move_timeout, 
        uint _settle_timeout,
        uint8 _min_stone_offchain,
        uint8 _max_stone_onchain) 
        public 
    {
        players[0] = _player1;
        players[1] = _player2;
        move_timeout = _move_timeout;
        settle_timeout = _settle_timeout;
        min_stone_offchain = _min_stone_offchain;
        max_stone_onchain = _max_stone_onchain;
        nonce = 0;
        winner = 0;
        turn = 1; // player1 moves first
        stone_num = 0;
        stone_num_onchain = 0;
        game_phase = GamePhase.start;
    }

    /**
    @notice Place a stone on the board
    @param _x is x coordinate on the board
    @param _y is y coordinate on the board
    */
    function placeStone(uint8 _x, uint8 _y) public {
        require(game_phase != GamePhase.over, "Game is over");
        require(game_phase != GamePhase.settle, "Settling off-chain state");
        require(msg.sender == players[turn-1], "Not your turn");
        require(checkBoundary(_x, _y), "Out of boundary");
        require(board[_x][_y] == 0, "The slot is already taken");

        // place the stone
        board[_x][_y] = turn;
        nonce++;
        stone_num++;
        stone_num_onchain++;
        emit PlaceStone(_x, _y);

        // check if there is five-in-a-row including this new stone
        if (checkFive(_x, _y, 1, 0) ||  // horizaontal bidirection
            checkFive(_x, _y, 0, 1) ||  // vertical bidirection
            checkFive(_x, _y, 1, 1) ||  // main-diagonal bidirection
            checkFive(_x, _y, 1, -1)) { // anti-diagonal bidirection
            winGame(turn); // we have a winner
            return;
        }

        if (stone_num == 225 || stone_num_onchain > max_stone_onchain) {
            // all slots occupied, game is over with no winner
            game_phase = GamePhase.over;
            turn = 0;
            emit GameOver(0);
        } else {
            // toggle turn and update game phase
            if (turn == 1) { turn = 2; } else { turn = 1; }
            move_deadline = block.number + move_timeout;
            game_phase = GamePhase.play;
        }
    }

    /**
    @notice Submit off-chain game state proof
    @param _stateproof is serialized off-chain state
    @param _signatures is serialized signatures
    */
    function intendSettle(bytes _stateproof, bytes _signatures) public {
        require(game_phase != GamePhase.over);
        pbRpcMultiSignature.Data memory sigs = pbRpcMultiSignature.decode(_signatures);
        require(checkSignature(_stateproof, sigs), "invalid signatures");
        if(game_phase == GamePhase.settle) {
            require(block.number <= settle_deadline);
        }
        loadGameState(_stateproof); 
        if(game_phase == GamePhase.start || game_phase == GamePhase.play) {
            game_phase = GamePhase.settle;
            settle_deadline = block.number + settle_timeout;
        }
    }

    /**
    @notice Confirm off-chain state is settled and update on-chain states
    */
    function confirmSettle() public {
        if (game_phase == GamePhase.settle && block.number >= settle_deadline) {
            game_phase = GamePhase.play; 
            move_deadline = block.number + move_timeout;
            emit ConfirmSettle();
        }
    }

    /**
    @notice Check if the game is finalized
    @param _query is query data (empty in Gomoku game) 
    @param _timeout is deadline (block number) for the game to be finalized
    @return true if game is finalized before given timeout
    */
    function isFinalized(bytes _query, uint _timeout) public view returns (bool) {
        return game_phase == GamePhase.over && finalized_time <= _timeout;
    }

    /**
    @notice Query the game result
    @param _query is query data (player address in Gomoku game) 
    @return true if given player wins
    */
    function queryResult(bytes _query) public view returns (bool) {
        require(_query.length == 1);
        return winner == uint8(_query[0]);
    }

    /**
    @notice Update game state in case of on-chain move timeout
    @return true if the game is over due to on-chain move timeout
    */
    function finalizeOnMoveTimeout() public returns (bool) {
        if (game_phase == GamePhase.play && block.number > move_deadline) {
            if (turn == 1) {
                winGame(2);
            } else if (turn == 2) {
                winGame(1);
            }
            return true;
        }
        return false;
    }

    /**
    @notice Query game state
    @return serialized game state
    */
    function queryState() public view returns (bytes) {
        bytes memory state = new bytes(227);
        state[0] = byte(winner);
        state[1] = byte(turn);
        uint i = 2;
        for (uint8 x = 0; x < board_dimmension; x++) {
            for (uint8 y = 0; y < board_dimmension; y++) {
                state[i] = byte(board[x][y]);
                i++;
            }
        }
        return state;
    }

    /**
    @notice Get the deadline of off-chain state settle
    @return block number of the settle deadline
    */
    function getSettleTime() public view returns (uint) {
        return settle_deadline;
    }

    /**
    @notice Check if there is five in a row in a given direction
    @param _x is x coordinate on the board
    @param _y is y coordinate on the board
    @param _xdir is direction (-1 or 0 or 1) in x axis
    @param _ydir is direction (-1 or 0 or 1) in y axis
    @return true if there is five in a row
    */
    function checkFive(uint8 _x, uint8 _y, int _xdir, int _ydir) private view returns (bool) {
        uint8 count = 0;
        count += countStone(_x, _y, _xdir, _ydir);
        count += countStone(_x, _y, -1 * _xdir, -1 * _ydir) - 1; // reverse direction
        if (count >= 5) {
            return true;
        }
        return false;
    }

    /**
    @notice Count the maximum consecutive stones in a given direction
    @param _x is x coordinate on the board
    @param _y is y coordinate on the board
    @param _xdir is direction (-1 or 0 or 1) in x axis
    @param _ydir is direction (-1 or 0 or 1) in y axis
    @return the number of maximum consecutive stones in a given direction
    */
    function countStone(uint8 _x, uint8 _y, int _xdir, int _ydir) private view returns (uint8) {
        uint8 count = 1;
        while (count <= 5) {
            uint8 x = uint8(int8(_x) + _xdir * count);
            uint8 y = uint8(int8(_y) + _ydir * count);
            if (checkBoundary(x, y) && board[x][y] == board[_x][_y]) {
                count += 1;
            } else {
                return count;
            }
        }
    }

    /**
    @notice Check if coordinate (x, y) is valid
    @param _x is x coordinate on the board
    @param _y is y coordinate on the board
    @return true is (x, y) is a valid coordinate
    */
    function checkBoundary(uint8 _x, uint8 _y) private pure returns (bool) {
        return (_x < board_dimmension && _y < board_dimmension);
    }

    /**
    @notice Verify state signatures from both players
    @param _stateproof is off-chain game state
    @param _sigs is signatures of the two players
    @return true is the signatures are both valid
    */
    function checkSignature(bytes _stateproof, pbRpcMultiSignature.Data _sigs) internal returns (bool) {
        require(_sigs.v.length == 2);
        bytes32 mhash = keccak256(_stateproof);
        bytes32 h = keccak256("\x19Ethereum Signed Message:\n32", mhash);
        address addr1 = ecrecover(h, _sigs.v[0], _sigs.r[0], _sigs.s[0]);
        address addr2 = ecrecover(h, _sigs.v[1], _sigs.r[1], _sigs.s[1]);
        if ((addr1 == players[0] && addr2 == players[1]) || (addr1 == players[1] && addr2 == players[0])) {
            return true;
        } else {
            return false;
        }
    }

    /**
    @notice Load off-chain game state
    @param _stateproof is off-chain game state
    */
    function loadGameState(bytes _stateproof) private {
        pbRpcStateProof.Data memory stateproof = pbRpcStateProof.decode(_stateproof);
        require(stateproof.nonce > nonce); // nonce is larger than current one
        nonce = stateproof.nonce;
        // stateproof.state: uint8 winner + uint8 turn + 15x15 bytes board
        require(stateproof.state.length == 227);
        winGame(uint8(stateproof.state[0]));
        // load other states only if winner is not specified
        if(winner == 0) {
            uint16 i = 1;
            turn = uint8(stateproof.state[i++]);
            require(0 <= turn && turn <= 2, "invalid turn state");
            stone_num = 0;
            for (uint8 x = 0; x < board_dimmension; x++) {
                for (uint8 y = 0; y < board_dimmension; y++) {
                    board[x][y] = uint8(stateproof.state[i++]);
                    require(0 <= board[x][y] && board[x][y] <= 2, "invalid board state");
                    if (board[x][y] != 0) {
                        stone_num++;
                    }
                }
            }
            require(stone_num >= min_stone_offchain); // need at least minimal number of stones
        }
        emit IntendSettle(nonce);
    }

    /**
    @notice Set game states when there winner
    @param _winner is id of the winner
    */
    function winGame(uint8 _winner) private {
        require(0 <= _winner && _winner <= 2, "invalid winner state");
        winner = _winner;
        // Game over
        if (winner != 0) {
            game_phase = GamePhase.over;
            turn = 0;
            finalized_time = block.number;
            emit GameOver(winner);
        }
    }

}
