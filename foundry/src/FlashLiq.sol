// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.11;
pragma abicoder v2;

//import "../lib/ds-test/src/test.sol";
import { FlashLoanReceiverBase } from "../lib/contracts/FlashLoanReceiverBase.sol";
import { ILendingPool } from "../lib/interfaces/ILendingPool.sol";
import { ILendingPoolAddressesProvider } from "../lib/interfaces/ILendingPoolAddressesProvider.sol";
import { ISwapRouter } from "../lib/interfaces/ISwapRouter.sol";
import { IERC20 } from "../lib/interfaces/IERC20.sol";

// Flash liq will be passed into flashloan receiverAddress.
contract FlashLiq is FlashLoanReceiverBase {

    uint256 MAX_UINT256 = 2**256 - 1;
    address public _owner;
    ISwapRouter public swapRouter = ISwapRouter(0xE592427A0AEce92De3Edee1F18E0157C05861564);

    // Constructor from base class + init owner.
    constructor(address _addressProvider) FlashLoanReceiverBase(_addressProvider) {
        _owner = msg.sender;
    }

    // Called from flashloan(), contract has funds at this point.
    function executeOperation(
        address[] calldata assets, 
        uint256[] calldata amounts, 
        uint256[] calldata premiums, 
        address            initiator, 
        bytes     calldata params
    ) external override returns (bool) {

        // Decode parameters.
        (
            address collateralAddress,
            address target
        )
        = abi.decode(params,
        (
            address,
            address
        ));
        
        require(initiator == _owner, "Only owner can execute this operation.");

        IERC20(assets[0]).approve(address(LENDING_POOL), MAX_UINT256);
        // Execute liquidation.
        LENDING_POOL.liquidationCall(
            collateralAddress,  // collateral - address of the collateral reserve.
            assets[0],          // debt - address of the debt reserve.
            target,             // user - address of the borrower.
            amounts[0],         // debtToCover - amount of debt to cover.
            false               // receiveAsToken - if false receive underlying asset directly.    
        );

        IERC20(collateralAddress).approve(address(swapRouter), amounts[0] + premiums[0]);
        // Swap collateral for debt to repay.
        swapRouter.exactOutputSingle(
            ISwapRouter.ExactOutputSingleParams(
                collateralAddress,          // tokenIn - address
                assets[0],                  // tokenOut - address
                3000,                       // fee - uint24
                address(this),              // recipient - address
                10**18,                     // deadline - uint256
                amounts[0] + premiums[0],   // amountOut - uint256
                0,                          // amountInMaximum - uint256
                0                           // sqrtPriceLimitX96 - uint160
            )
        );
        return true;
    }
}
