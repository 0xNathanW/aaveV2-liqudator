// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "../../lib/ds-test/src/test.sol";
import { FlashLiq } from "../FlashLiq.sol";
import { ISwapRouter } from "../../lib/interfaces/ISwapRouter.sol";
import { ILendingPool } from "../../lib/interfaces/ILendingPool.sol";
import { IERC20 } from "../../lib/interfaces/IERC20.sol";
import { IWETH } from "../../lib/interfaces/IWETH.sol";
import { SafeERC20 } from "../../lib/contracts/SafeERC20.sol";
import { TestWithdrawl } from "./SafeTransfer.sol";
import { ILendingPoolAddressesProvider } from "../../lib/interfaces/ILendingPoolAddressesProvider.sol";
import { IPriceOracle } from "../../lib/interfaces/IPriceOracle.sol";
import { IChainlinkAggregator } from "../../lib/interfaces/IChainlinkAggregator.sol";

// Cheatcodes defined by DappTools.
interface CheatCodes {
    function deal(address, uint256) external;
    function prank(address) external;
    function startPrank(address) external;
    function stopPrank() external;
    function mockCall(address where, bytes calldata data, bytes calldata retdata) external;
}

interface DataProvider {
    function getAllReserveTokens() external;
}

contract ContractTest is DSTest {

    using SafeERC20 for IERC20;
    
    address _contract;

    FlashLiq flashLiq;  // This is the contract we are testing.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);  // Access cheatcodes.

    address fakeOwner   = 0xF831de50F3884Cf0f4550Bb129032a80CB5a26B7;
    address target      = 0xabc9b5632AB9176926283F25eF252a71603B8214;     

    address collateral  = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2; //WETH
    address debt        = 0x6B175474E89094C44Da98b954EedeAC495271d0F; //DAI

    ILendingPool lendingPool = ILendingPool(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9);
    IPriceOracle priceOracle = IPriceOracle(0xA50ba011c48153De246E5192C8f9258A2ba79Ca9);

    function setUp() public {
        cheats.startPrank(fakeOwner);
        flashLiq = new FlashLiq(0xB53C1a33016B2DD2fF3653530bfF1848a515c8c5);
        _contract = address(flashLiq);
        cheats.stopPrank();
    }

    // Debugging error on safeTransfer.
    function testSafeTransfer() public {

        TestWithdrawl taker = new TestWithdrawl();

        cheats.deal(target, 10 ether);
        cheats.startPrank(target);
        
        IWETH(collateral).deposit{value: 5 ether}();
        emit log_named_uint("WETH Balance", IERC20(collateral).balanceOf(target)); 
        IERC20(collateral).approve(address(taker), 5 ether);
        
        uint256 amount = 3 ether;
        taker.take(collateral, amount);
        
        emit log_named_uint("WETH Balance", IERC20(collateral).balanceOf(target));
        cheats.stopPrank();
    }

    // Main test function.
    function testMain() public {
        setupEnv(); // Setup the environment.
        
        cheats.startPrank(fakeOwner);
        emit log_named_address("\n  Contract address", address(flashLiq));

        emit log("\n  Before execution");
        printContractBalance();
    
        address[] memory assets = new address[](1);
        assets[0] = debt;
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 0.05 ether;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 0;

        bytes memory params = abi.encode(
            collateral,
            target
        );
        
        lendingPool.flashLoan(
            _contract, 
            assets, 
            amounts, 
            modes, 
            _contract, 
            params, 
            0
            );

        emit log("\n  After execution");
        printContractBalance();
        cheats.stopPrank();
    }

    // Setup environment where which target has a liquidatable position.
    function setupEnv() public {
        /* 
            HealthFactor = (total collat in ETH * liq threshold) / (total borrowed in ETH)
        */
        cheats.deal(target, 10 ether);
        cheats.startPrank(target);

        IWETH(collateral).deposit{value: 10 ether}();   // Get 10 WETH.
        assertEq(IERC20(collateral).balanceOf(target), 10 ether); // Balance 10 WETH.
        IERC20(collateral).approve(0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9, 5 ether);
        emit log_named_uint("Allowed to withdraw", IERC20(collateral).allowance(target, 0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9));
        emit log_named_address("Sender", msg.sender);

        lendingPool.deposit(collateral, 0.15 ether, target, 0);
        lendingPool.setUserUseReserveAsCollateral(collateral, true);
        emit log("Borrowing...");
        lendingPool.borrow(debt, 300 ether, 1, 0, target);
        printUserData();
        
        emit log_named_uint("Asset price", priceOracle.getAssetPrice(debt));
        cheats.mockCall(
            0x773616E4d11A78F511299002da57A0a94577F1f4,
            abi.encodeWithSelector(IChainlinkAggregator.latestAnswer.selector),
            abi.encode(474254113057743)
        );
        emit log_named_uint("New asset price", priceOracle.getAssetPrice(debt));

        printUserData();
        cheats.stopPrank();
    }

    function printUserData() public {
        (   
            uint256 totalCollateral,
            uint256 totalDebt,
            uint256 availableBorrow,
            uint256 currentLiquidationThreshold,
            uint256 ltv,
            uint256 healthFactor
        ) = lendingPool.getUserAccountData(target);

        string memory healthy = (healthFactor > 1 ether) ? "No" : "Yes";
        emit log("----------------------------------");
        emit log_named_decimal_uint("Total collateral", totalCollateral, 18);
        emit log_named_decimal_uint("Total debt", totalDebt, 18);
        emit log_named_decimal_uint("Available borrow", availableBorrow, 18);
        emit log_named_decimal_uint("AWETH", IERC20(0x030bA81f1c18d280636F32af80b9AAd02Cf0854e).balanceOf(target), 18);
        emit log_named_uint("Current liquidation threshold", currentLiquidationThreshold);
        emit log_named_uint("LTV", ltv);
        emit log_named_decimal_uint("Health factor", healthFactor, 18);
        emit log_named_string("Can liquidate?", healthy);
        emit log_named_decimal_uint("Liquidatable amount", totalDebt/2, 18);
        emit log("----------------------------------");
    }

    function printContractBalance() public {
        emit log("--- Contract balance ---");
        emit log_named_address("Owner", _contract);
        emit log_named_decimal_uint(
            "WETH", IERC20(collateral).balanceOf(_contract), 18
            );
        emit log_named_decimal_uint(
            "DAI", IERC20(debt).balanceOf(_contract), 18
            );
        emit log_named_decimal_uint(
            "Ether", _contract.balance, 18
            );
        emit log("-------------------");
    }

    function testEncode() public {
    bytes memory data = abi.encode(
        collateral,
        target
    );
    emit log_named_bytes("Encoded data", data);
    }
    
}

