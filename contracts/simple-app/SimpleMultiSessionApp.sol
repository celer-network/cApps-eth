pragma solidity ^0.5.0;
import "../templates/MultiSessionBooleanResult.sol";

contract SimpleMultiSessionApp is MultiSessionBooleanResult {
    struct GameInfo {
        uint8 state;
    }

    mapping(bytes32 => GameInfo) private gameInfoMap; // session ID -> app data

    constructor(uint _timeout, uint _playerNum)
        public MultiSessionBooleanResult(_timeout, _playerNum)
    {
    }

    /**
     * @notice Get the app result
     * @param _session Session ID
     * @param _query Query args
     * @return True if given player wins
     */
    function getResult(bytes32 _session, bytes memory _query)
        internal
        view
        returns (bool)
    {
        require(_query.length == 1, "invalid query length");
        return gameInfoMap[_session].state == uint8(_query[0]);
    }

    /**
     * @notice Get app state
     * @param _session App session ID
     * @param _key Query key
     * @return App state associated with the key
     */
    function getState(bytes32 _session, uint _key) external view returns (bytes memory) {
        bytes memory b = new bytes(32);
        uint8 x = gameInfoMap[_session].state;
        assembly { mstore(add(b, 32), x) }
        return b;
    }

    /**
     * @notice Update on-chain state according to off-chain state proof
     * @param _session Session ID
     * @param _state Signed off-chain state
     */
    function updateByState(bytes32 _session, bytes memory _state)
        internal
        returns (bool)
    {
        require(_state.length == 1, "invalid state length");
        GameInfo storage game = gameInfoMap[_session];
        game.state = uint8(_state[0]);
        if (game.state == 1 || game.state == 2) {
            sessionInfoMap[_session].status = SessionStatus.FINALIZED;
        }
        return true;
    }

    /**
     * @notice Update state according to an on-chain action
     * @param _session Session ID
     * @param _action Action data
     * @return True if update succeeds
     */
    function updateByAction(bytes32 _session, bytes memory _action)
        internal
        returns (bool)
    {
        require(_action.length == 1, "invalid action length");
        GameInfo storage app = gameInfoMap[_session];
        app.state = uint8(_action[0]);
        if (app.state == 1 || app.state == 2) {
            sessionInfoMap[_session].status = SessionStatus.FINALIZED;
        }
        return true;
    }

    /**
     * @notice Finalize the session based on current state in case of on-chain action timeout
     * @param _session Session ID
     */
    function finalizeOnTimeout(bytes32 _session) internal  {
        sessionInfoMap[_session].status = SessionStatus.FINALIZED;
    }
}
