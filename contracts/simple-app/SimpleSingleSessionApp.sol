pragma solidity ^0.5.0;
import "../templates/SingleSessionBooleanResult.sol";

contract SimpleSingleSessionApp is SingleSessionBooleanResult {
    uint8 private state;

    constructor(
        address[] memory _players,
        uint _nonce,
        uint _timeout)
        public
        SingleSessionBooleanResult(_players, _nonce, _timeout)
    {
    }

    /**
     * @notice Get the app result
     * @param _query Query args
     * @return True if query satisfied
     */
    function getResult(bytes calldata _query) external view returns (bool) {
        require(_query.length == 1, "invalid query length");
        return state == uint8(_query[0]);
    }

    /**
     * @notice Get app state
     * @param _key Query key
     * @return App state associated with the key
     */
    function getState(uint _key) external view returns (bytes memory) {
        bytes memory b = new bytes(32);
        uint8 x = state;
        assembly { mstore(add(b, 32), x) }
        return b;
    }

    /**
     * @notice Update state according to an off-chain state proof
     * @param _state Signed off-chain app state
     * @return True if update succeeds
     */
    function updateByState(bytes memory _state) internal returns (bool) {
        require(_state.length == 1, "invalid state length");
        state = uint8(_state[0]);
        if (state == 1 || state == 2) {
            appInfo.status = AppStatus.FINALIZED;
        }
        return true;
    }

    /**
     * @notice Update state according to an on-chain action
     * @param _action Action data
     * @return True if update succeeds
     */
    function updateByAction(bytes memory _action) internal returns (bool) {
        require(_action.length == 1, "invalid action length");
        state = uint8(_action[0]);
        if (state == 1 || state == 2) {
            appInfo.status = AppStatus.FINALIZED;
        }
        return true;
    }

    /**
     * @notice Finalize in case of on-chain action timeout
     */
    function finalizeOnTimeout() internal {
        appInfo.status = AppStatus.FINALIZED;
    }
}
