pragma solidity ^0.5.0;
import "./SingleSessionApp.sol";
import "./IBooleanOutcome.sol";

contract SingleSessionBooleanOutcome is SingleSessionApp, IBooleanOutcome {

    constructor(
        address[] memory _players,
        uint _nonce,
        uint _timeout)
        public
        SingleSessionApp(_players, _nonce, _timeout)
    {
    }

    /**
     * @notice Check if the app is finalized
     * @return True if app is finalized
     */
    function isFinalized(bytes calldata _query) external view returns (bool) {
        require(_query.length == 0);
        return appInfo.status == AppStatus.FINALIZED;
    }

}
