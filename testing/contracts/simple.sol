/**
 *Submitted for verification at Etherscan.io on 2021-09-15
 */

pragma solidity 0.8.10;

contract SimpleStorage {
    string a;
    string b;
    event StorageSetA(string NewString);
    event StorageSetB(string NewString);

    function setA(string memory x) public {
        a = x;
        emit StorageSetA(x);
    }

    function setB(string memory x) public {
        b = x;
        emit StorageSetB(x);
    }

    function getA() public view returns (string memory) {
        return a;
    }

    function getB() public view returns (string memory) {
        return b;
    }
}
