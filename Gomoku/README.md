cGomoku on-chain contract that impletements Celer Network cOS API.

#### Public functions - on-chain game API

```solidity
/**
@notice Place a stone on the board
@param _x is x coordinate on the board
@param _y is y coordinate on the board
*/
function placeStone(uint8 _x, uint8 _y)

/**
@notice Update game state in case of on-chain move timeout
@return true if the game is over due to on-chain move timeout
*/
function finalizeOnMoveTimeout() public returns (bool)

/**
@notice Query game state
@return serialized game state
*/
function queryState() public view returns (bytes)

```

#### Public functions - cOS required API

```solidity
/**
@notice Submit off-chain game state proof
@param _stateproof is serialized off-chain state
@param _signatures is serialized signatures
*/
function intendSettle(bytes _stateproof, bytes _signatures)

/**
@notice Confirm off-chain state is settled and update on-chain states
*/
function confirmSettle()

/**
@notice Check if the game is finalized
@param _query is query data (empty in Gomoku game) 
@param _timeout is deadline (block number) for the game to be finalized
@return true if game is finalized before given timeout
*/
function isFinalized(bytes _query, uint _timeout) public view returns (bool)

/**
@notice Query the game result
@param _query is query data (player address in Gomoku game) 
@return true if given player wins
*/
function queryResult(bytes _query) public view returns (bool)

/**
@notice Get the deadline of off-chain state settle
@return block number of the settle deadline
*/
function getSettleTime() public view returns (uint)
```