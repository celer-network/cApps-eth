pragma solidity ^0.5.0;

interface ISingleSession {

    enum AppStatus {IDLE, SETTLE, ACTION, FINALIZED}

    event IntendSettle(uint seq);

    /**
     * @notice Submit off-chain state and start to dispute
     */
    function intendSettle(bytes calldata _stateProof) external;

    /**
     * @notice get app state settle finalized time
     */
    function getSettleFinalizedTime() external view returns (uint);

    /**
     * @notice Apply an action to the on-chain state
     */
    function applyAction(bytes calldata _action) external;

    /**
     * @notice get action deadline
     */
    function getActionDeadline() external view returns (uint);

    /**
     * @notice Finalize in case of on-chain action timeout
     */
    function finalizeOnActionTimeout() external;

    /**
     * @notice get app status
     */
    function getStatus() external view returns (AppStatus);

    /**
     * @notice get app sequence number
     */
    function getSeqNum() external view returns (uint);

    /**
     * @notice Get app state associated with the given key
     */
    function getState(uint _key) external view returns (bytes memory);
}
