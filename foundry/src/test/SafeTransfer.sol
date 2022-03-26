pragma solidity ^0.8.10;

import { IERC20 } from "../../lib/interfaces/IERC20.sol";
import { SafeERC20 } from "../../lib/contracts/SafeERC20.sol";

contract TestWithdrawl {

    using SafeERC20 for IERC20;

    function take(address asset, uint _amount) public {
        IERC20(asset).safeTransferFrom(msg.sender, address(this), _amount);
    }
}