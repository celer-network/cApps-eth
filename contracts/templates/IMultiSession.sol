pragma solidity ^0.5.0;

interface IMultiSession {

    enum SessionStatus {IDLE, SETTLE, ACTION, FINALIZED}

    event IntendSettle(bytes32 session, uint seq);

    /**
     * @notice Submit off-chain state and start to dispute
     */
    function intendSettle(bytes calldata _stateProof) external;

    /**
     * @notice Get app session state settle finalized time
     * @param _session App session ID
     */
    function getSettleFinalizedTime(bytes32 _session) external view returns (uint);

    /**
     * @notice Apply an action to the on-chain state
     */
    function applyAction(bytes32 _session, bytes calldata _action) external;

    /**
     * @notice Finalize the session in case of on-chain action timeout
     */
    function finalizeOnActionTimeout(bytes32 _session) external;

    /**
     * @notice Get app session action deadline
     */
    function getActionDeadline(bytes32 _session) external view returns (uint);

    /**
     * @notice Get app session status
     */
    function getStatus(bytes32 _session) external view returns (SessionStatus);

    /**
     * @notice Get app session sequence number
     */
    function getSeqNum(bytes32 _session) external view returns (uint);

    /**
     * @notice Get app session state associated with the given key
     */
    function getState(bytes32 _session, uint _key) external view returns (bytes memory);

    /**
     * @notice Compute session ID
     * @param _nonce app session nonce
     * @param _signers signers required for the offchain state to be valid
     * @return Session ID that encodes nonce and singers
     */
    function getSessionID(uint _nonce, address[] calldata _signers) external pure returns (bytes32);
}
