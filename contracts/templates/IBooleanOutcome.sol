pragma solidity ^0.5.0;

interface IBooleanOutcome {
    function isFinalized(bytes calldata _query) external view returns (bool);
    function getOutcome(bytes calldata _query) external view returns (bool);
}
