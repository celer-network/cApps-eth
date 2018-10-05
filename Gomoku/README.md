Gomoku on-chain contract that impletement cOS API so that it can work with cGomoku App running on Celer Network.

### Public functions

```solidity
/**
@notice Check if the game is finalized (cOS API)
@param _query is query data (empty in Gomoku game) 
@param _timeout is deadline (block number) for the game to be finalized
@return true if game is finalized before given timeout
*/
function isFinalized(bytes _query, uint _timeout) public view returns (bool)

/**
@notice Query the game result (cOS API)
@param _query is query data (player address in Gomoku game) 
@return true if given player wins
*/
function queryResult(bytes _query) public view returns (bool)
    
/**
@notice Submit off-chain game state proof (cOS API)
@param _stateproof is serialized off-chain state
@param _signatures is serialized signatures
*/
function intendSettle(bytes _stateproof, bytes _signatures)

/**
@notice Confirm off-chain state is settled and update on-chain states (cOS API)
*/
function confirmSettle()

/**
@notice Query game state
@return serialized game state
*/
function queryState() public view returns (bytes)

/**
@notice Get the deadline of off-chain state settle
@return block number of the settle deadline
*/
function getSettleTime() public view returns (uint)

/**
@notice Update game state in case of on-chain move timeout
@return true if the game is over due to on-chain move timeout
*/
function isTimeout() public returns (bool)

/**
@notice Place a stone on the board
@param _x is x coordinate on the board
@param _y is y coordinate on the board
*/
function placeStone(uint8 _x, uint8 _y)
```