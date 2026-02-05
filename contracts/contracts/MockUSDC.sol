// SPDX-License-Identifier: MIT
// Clearly not written by us ;)
pragma solidity ^0.8.33;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockUSDC is ERC20 {
    constructor() ERC20("Mock USDC", "mUSDC") {
    }

    /**
     * @dev USDC uses 6 decimals on Ethereum/Sepolia.
     */
    function decimals() public pure override returns (uint8) {
        return 6;
    }

    /**
     * @notice Simple mint function for testing.
     * @param to The address to receive the tokens.
     * @param amount The amount in whole tokens (will be multiplied by 10^6).
     */
    function mint(address to, uint256 amount) public {
        _mint(to, amount * 10 ** decimals());
    }
}