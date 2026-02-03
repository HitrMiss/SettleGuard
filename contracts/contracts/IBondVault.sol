// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IBondVault {
    /**
     * @notice Returns the active, locked bond amount for any address
     * @param party The address to lookup
     * @return uint256 The amount of tokens currently locked as bond.
     */
    function getActiveBond(address party) external view returns (uint256);

    /**
     * @notice Arbiter function to slash a bond in case of fraud.
     * @param party the address to slash.
     * @param amount the amount to refund.
     * @param recipient where to issue the refund.
     */
    function slashBond(address party, uint256 amount, address recipient) external;
}