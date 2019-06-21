pragma solidity ^0.5.0;
import "./MultiSessionApp.sol";
import "./IBooleanOutcome.sol";

contract MultiSessionBooleanOutcome is MultiSessionApp, IBooleanOutcome {

    constructor(uint _playerNum) public MultiSessionApp(_playerNum) {}

    /**
     * @notice Get the app outcome
     * @param _session Session ID
     * @param _query Query args
     * @return True if query satisfied
     */
    function getOutcome(bytes32 _session, bytes memory _query) internal view returns (bool);

    /**
     * @notice Get the app outcome
     * @param _query Query arg starting with session ID
     * @return True if query satisfied
     */
    function getOutcome(bytes calldata _query) external view returns (bool) {
        PbApp.SessionQuery memory req = PbApp.decSessionQuery(_query);
        return getOutcome(req.session, req.query);
    }

    /**
     * @notice Check if the app session is finalized
     * @param _query Query arg that encodes session ID
     * @return True if app session is finalized
     */
    function isFinalized(bytes calldata _query) external view returns (bool) {
        bytes32 session = Pb._bytes32(_query);
        return sessionInfoMap[session].status == SessionStatus.FINALIZED;
    }
}
