pragma solidity ^0.5.0;

interface IBooleanResult {
    function isFinalized(bytes calldata _query) external view returns (bool);
    function getResult(bytes calldata _query) external view returns (bool);
}
