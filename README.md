# Celer dApp Contract

Celer dApps are highly interactive, secure and low-cost state-channel applications running on [Celer Network](https://www.celer.network) together with Celer [generic payment channel](https://github.com/celer-network/cChannel-eth).

This repo provides templates and examples for developing the on-chain contract parts of dApps that can smoothly run on Celer mobile and web SDK [CelerX](https://celerx.app/). Note that most of the app nteractions happen off-chain. On-chain operations are only needed when players cannot reach consensus off-chain and want to dispute.

- **Multi-Session App:** initially deployed once by the developer and can be repeatedly shared by all players. No additional code needs to be deployed when players want to dispute on-chain.

- **Single-Session App:** mostly used as a one-time virtual contract for fixed players without initial deployment. The player who wants to bring the off-chain game to on-chain dispute needs first to deploy the contract.


## Latest Deployments

### Ropsten Testnet

#### MultiGomoku
- Contract address: [0xa5cb3fcb9fc2c7a285615455735530ed0d1dfeb6](https://ropsten.etherscan.io/address/0xb352b23620ab8d75a05012aec0e0f5ce1015d743)
- Constructor: minStoneOffchain=5， maxStoneOnchain=3
- Deployed code: [MultiGomoku.sol](https://github.com/celer-network/cApps-eth/blob/26e26e1f152cc6ad11675b22383c868839055358/contracts/gomoku/MultiGomoku.sol)

### Mainnet

#### MultiGomoku
- Contract address: [0x8521714a9ef8f9e63e5adb4246207a04815099b9](https://etherscan.io/address/0x8521714a9ef8f9e63e5adb4246207a04815099b9)
- Deployed code: [MultiGomoku.sol](https://github.com/celer-network/cApps-eth/blob/88391d4953/contracts/gomoku/MultiGomoku.sol)

## External API

#### API required by [CelerChannel](https://github.com/celer-network/cChannel-eth)

- [App with Boolean Outcome](https://github.com/celer-network/cApps-eth/blob/master/contracts/templates/IBooleanOutcome.sol)

- [App with Numeric Outcome](https://github.com/celer-network/cApps-eth/blob/master/contracts/templates/INumericOutcome.sol)

#### API required by [CelerX](https://celerx.app/).

- [Multi-session Apps](https://github.com/celer-network/cApps-eth/blob/master/contracts/templates/IMultiSession.sol)

- [Single-session Apps](https://github.com/celer-network/cApps-eth/blob/master/contracts/templates/ISingleSession.sol)


## Template Interface

We provide [templates](https://github.com/celer-network/cApps-eth/tree/master/contracts/templates) to implement the common state-channel logics of external APIs, so that the developers can focus on the app-specific logic.

Developers using the provided templates **only need to implement the following interfaces**. For detailed usages, please refer to these [simplest example contracts](https://github.com/celer-network/cApps-eth/tree/master/contracts/simple-app) and [tests](https://github.com/celer-network/cApps-eth/tree/master/test/simple-app)

#### MultiSessionApp template interface

```javascript
/**
 * @notice Get the app outcome
 * @param _session Session ID
 * @param _query Query arg
 * @return True if query satisfied
 */
function getOutcome(bytes32 _session, bytes memory _query) internal view returns (bool) {}

/**
 * @notice Update on-chain state according to off-chain state proof
 * @param _session Session ID
 * @param _state Signed off-chain state
 */
function updateByState(bytes32 _session, bytes memory _state) internal returns (bool) {}

/**
 * @notice Update state according to an on-chain action
 * @param _session Session ID
 * @param _action Action data
 * @return True if update succeeds
 */
function updateByAction(bytes32 _session, bytes memory _action) internal returns (bool) {}

/**
 * @notice Finalize the session based on current state in case of on-chain action timeout
 * @param _session Session ID
 */
function finalizeOnTimeout(bytes32 _session) internal {}

/**
 * @notice Get app state associated with the given key
 */
function getState(bytes32 _session, uint _key) external view returns (bytes memory);
```

#### SingleSessionApp template interface

```javascript
/**
 * @notice Get the app outcome
 * @param _query Query args
 * @return True if query satisfied
 */
function getOutcome(bytes memory _query) public view returns (bool) {}

/**
 * @notice Update state according to an off-chain state proof
 * @param _state Signed off-chain app state
 * @return True if update succeeds
 */
function updateByState(bytes memory _state) internal returns (bool) {}

/**
 * @notice Update state according to an on-chain action
 * @param _action Action data
 * @return True if update succeeds
 */
function updateByAction(bytes memory _action) internal returns (bool) {}

/**
 * @notice Finalize based on current state in case of on-chain action timeout
 */
function finalizeOnTimeout() internal {}

/**
 * @notice Get app state associated with the given key
 */
function getState(uint _key) external view returns (bytes memory);
```

## Protobuf

We leverage Protocol Buffers to define a series of blockchain-neutral generalized data structures, which can be seamlessly used in off-chain communication protocols and instantly extended to other blockchains that we plan to support. We also developed and open sourced a Solidity library generator for decoding proto3 called [pb3-gen-sol](https://github.com/celer-network/pb3-gen-sol).

Below are the protos used by Celer dApps. [CelerX](https://celerx.app/) takes care of the protobuf encode and decode for app developers.

```protobuf
message AppState {
  // nonce should be unique for each app session among the same signers
  uint64 nonce = 1 [(soltype) = "uint"];
  // for each nonce, new state has higher sequence number
  uint64 seq_num = 2 [(soltype) = "uint"];
  // app specific state
  bytes state = 3;
  // on-chain response (settle, action) timeout
  uint64 timeout = 4 [(soltype) = "uint"];
}

message StateProof {
  // serialized AppState
  bytes app_state = 1;
  repeated bytes sigs = 2;
}

// used for multi-session app
message SessionQuery {
  // session ID
  bytes session = 1 [(soltype) = "bytes32"];
  // query related to the specified session
  bytes query = 2;
}
```
