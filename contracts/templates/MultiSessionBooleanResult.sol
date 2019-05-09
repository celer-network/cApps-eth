pragma solidity ^0.5.0;
import "./MultiSessionApp.sol";
import "./IBooleanResult.sol";

contract MultiSessionBooleanResult is MultiSessionApp, IBooleanResult {

    constructor(uint _timeout, uint _playerNum) public MultiSessionApp(_timeout, _playerNum) {}

    /**
     * @notice Get the app result
     * @param _session Session ID
     * @param _query Query args
     * @return True if query satisfied
     */
    function getResult(bytes32 _session, bytes memory _query) internal view returns (bool);

    /**
     * @notice Get the app result
     * @param _query Query arg starting with session ID
     * @return True if query satisfied
     */
    function getResult(bytes calldata _query) external view returns (bool) {
        PbSession.SessionQuery memory req = PbSession.decSessionQuery(_query);
        return getResult(req.session, req.query);
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
