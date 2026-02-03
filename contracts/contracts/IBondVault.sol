// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IBondVault {
    function getActiveBond(address party) external view returns (uint256);

    function lockRisk(address _user, uint256 _amount) external;
    function unlockRisk(address _user, uint256 _amount) external;

    function getAvailableBond(address _merchant) external view returns (uint256);

    function slashBond(address party, uint256 amount, address recipient) external;
}