// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "./IGovernance.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
//TODO: DO NOT USE, I REALLY NEED SLEEP
contract BondVault is ReentrancyGuard {
    IGovernance public immutable governance;

    mapping(address => uint256) public bonds;

    error Unauthorized();
    error TransferFailed();

    constructor(address _gov) {
        governance = IGovernance(_gov);
    }

    function deposit() external payable {
        bonds[msg.sender] += msg.value;
    }

    function withdraw(uint256 _amount) external nonReentrant {
        if (bonds[msg.sender] < _amount) revert Unauthorized();
        bonds[msg.sender] -= _amount;
        (bool success, ) = payable(msg.sender).call{value: _amount}("");
        if (!success) revert TransferFailed();
    }

    // Slash Dot Dash
    function slash(address _user, uint256 _amount, address _recipient) external {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) revert Unauthorized();
        uint256 balance = bonds[_user];
        uint256 amountToSlash = _amount > balance ? balance : _amount;
        bonds[_user] -= amountToSlash;
        payable(_recipient).transfer(amountToSlash);
    }
}