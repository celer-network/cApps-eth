pragma solidity ^0.5.0;
import "../lib/Pb.sol";
import "../lib/PbApp.sol";
import "../lib/LibSignature.sol";
import "./IMultiSession.sol";

contract MultiSessionApp is IMultiSession {

    struct AppInfo {
        uint playerNum;
    }

    struct SessionInfo {
        address[] players;
        uint seqNum;
        uint timeout;
        uint deadline;
        SessionStatus status;
    }

    AppInfo internal appInfo;
    mapping(bytes32 => SessionInfo) internal sessionInfoMap; // session id -> session info

    constructor(uint _playerNum) public {
        appInfo.playerNum =_playerNum;
    }

    /**
     * @notice Update on-chain state according to off-chain state proof
     * @param _session Session ID
     * @param _state Signed off-chain state
     */
    function updateByState(bytes32 _session, bytes memory _state) internal returns (bool);

    /**
     * @notice Update state according to an on-chain action
     * @param _session Session ID
     * @param _action Action data
     * @return True if update succeeds
     */
    function updateByAction(bytes32 _session, bytes memory _action) internal returns (bool);

    /**
     * @notice Finalize the session based on current state in case of on-chain action timeout
     * @param _session Session ID
     */
    function finalizeOnTimeout(bytes32 _session) internal;

    /**
     * @notice Get app state associated with the given key
     */
    function getState(bytes32 _session, uint _key) external view returns (bytes memory);

    /**
     * @notice Submit and settle off-chain state
     * @param _stateProof Serialized off-chain state with signatures
     */
    function intendSettle(bytes calldata _stateProof) external {
        PbApp.StateProof memory stateProof = PbApp.decStateProof(_stateProof);
        address[] memory signers =
            LibSignature.recoverSigners(stateProof.appState, stateProof.sigs, true);
        require(signers.length == appInfo.playerNum, "invalid number of players");
        PbApp.AppState memory appstate = PbApp.decAppState(stateProof.appState);
        bytes32 session = keccak256(abi.encode(appstate.nonce, signers));
        SessionInfo storage sessionInfo = sessionInfoMap[session];
        if (sessionInfo.status == SessionStatus.IDLE) {
            sessionInfo.players = signers;
        }
        require(sessionInfo.status != SessionStatus.FINALIZED, "state is finalized");
        require(sessionInfo.seqNum < appstate.seqNum, "invalid sequence number");
        sessionInfo.seqNum = appstate.seqNum;
        sessionInfo.status = SessionStatus.SETTLE;
        sessionInfo.timeout = appstate.timeout;
        sessionInfo.deadline = block.number + appstate.timeout;

        require(updateByState(session, appstate.state), "state update failed");
        emit IntendSettle(session, sessionInfo.seqNum);
    }

    /**
     * @notice Apply an action to the on-chain state
     * @param _session App session ID
     * @param _action Action request
     */
    function applyAction(bytes32 _session, bytes calldata _action) external {
        SessionInfo storage sessionInfo = sessionInfoMap[_session];
        require(sessionInfo.status != SessionStatus.FINALIZED, "app state is finalized");
        if (sessionInfo.status == SessionStatus.SETTLE && sessionInfo.deadline < block.number) {
            sessionInfo.status = SessionStatus.ACTION;
        }
        require(sessionInfo.status == SessionStatus.ACTION, "app not in action mode");
        sessionInfo.seqNum++;
        sessionInfo.deadline = block.number + sessionInfo.timeout;

        require(updateByAction(_session, _action));
    }

    /**
     * @notice Finalize in case of on-chain action timeout
     */
    function finalizeOnActionTimeout(bytes32 _session) external {
        SessionInfo storage sessionInfo = sessionInfoMap[_session];
        SessionStatus status = sessionInfo.status;
        uint deadline = sessionInfo.deadline;
        if (status == SessionStatus.ACTION) {
            require(block.number > deadline);
        } else if (status == SessionStatus.SETTLE) {
            require(block.number > deadline + sessionInfo.timeout);
        } else {
            return;
        }
        finalizeOnTimeout(_session);
    }

    /**
     * @notice get app session status
     */
    function getStatus(bytes32 _session) external view returns (SessionStatus) {
        return sessionInfoMap[_session].status;
    }

    /**
     * @notice get app session state settle finalized time
     */
    function getSettleFinalizedTime(bytes32 _session) external view returns (uint) {
        if(sessionInfoMap[_session].status == SessionStatus.SETTLE) {
            return sessionInfoMap[_session].deadline;
        }
        return 0;
    }

    /**
     * @notice get app session action deadline
     */
    function getActionDeadline(bytes32 _session) external view returns (uint) {
        if(sessionInfoMap[_session].status == SessionStatus.ACTION) {
            return sessionInfoMap[_session].deadline;
        } else if(sessionInfoMap[_session].status == SessionStatus.SETTLE) {
            return sessionInfoMap[_session].deadline + sessionInfoMap[_session].timeout;
        }
        return 0;
    }

    /**
     * @notice get app session sequence number
     */
    function getSeqNum(bytes32 _session) external view returns (uint) {
        return sessionInfoMap[_session].seqNum;
    }

    /**
     * @notice Compute session ID
     * @param _nonce app session nonce
     * @param _signers signers of the state
     * @return Session ID that encodes nonce and singers
     */
    function getSessionID(uint _nonce, address[] calldata _signers)
        external
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(_nonce, _signers));
    }
}
