pragma solidity ^0.5.0;

interface INumericResult {
    function isFinalized(bytes calldata _query) external view returns (bool);
    function getResult(bytes calldata _query) external view returns (uint);
}
