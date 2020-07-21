/**
*"SPDX-License-Identifier: copyleft-next 0.3.1
*/
pragma solidity ^0.6.6;

import "../zeppelin/contracts/math/SafeMath.sol";
import "../zeppelin/contracts/token/ERC20/IERC20.sol";
import "../zeppelin/contracts/token/ERC20/SafeERC20.sol";
import "../aave/FlashLoanReceiverBase.sol";
import "../aave/ILendingPool.sol";
import "./interfaces/IUniswap.sol";
import "../Kyber/KyberNetworkProxy.sol";


//1 DAI = 1000000000000000000 (18 decimals)
/*
 * Arbitrageur is a contract to simulate the usage of flashloans
 * to make profit out of a market inbalacement
 *
 * For this example we deployed 2 Uniswap instances which we'll
 * call by ExchangeA and ExchangeB
 *
 * The steps happens as following:
 * 1. Borrow DAI from Aave
 * 2. Buy BAT with DAI on ExchangeA
 * 3. Sell BAT for DAI on ExchangeB
 * 4. Repay Aave loan
 * 5. Keep the profits

  * Kovan addresses:
 * aave lending provider: 0x506B0B2CF20FAA8f38a4E2B524EE43e1f4458Cc5
 * uniswap 1 factory: 0xD3E51Ef092B2845f10401a0159B2B96e8B6c3D30
 * kyberswap  factory: 0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D
 * Dai: 0xC4375B7De8af5a38a93548eb8453a498222C4fF2
 *
 * Mainnet addresses:
 * aave lending provider: 0x24a42fD28C976A61Df5D00D0599C34c4f90748c8
 * DAI_ADDRESS = 0x6B175474E89094C44Da98b954EedeAC495271d0F;
 * uniswap 1 factory: 0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95
 * kyberswap  factory: 0x818E6FECD516Ecc3849DAf6845e3EC868087B755
 */
contract Uni1EthToKyberArbitrageur is
FlashLoanReceiverBase(address(0x24a42fD28C976A61Df5D00D0599C34c4f90748c8))
{
    event Swap(address indexed sender, KyberERC20 destToken, uint amount);

    // still kovan addresses for now
    address public constant DAI_ADDRESS = 0x6B175474E89094C44Da98b954EedeAC495271d0F;
    address public constant UNISWAP_1_FACTORY = 0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95;
    address public constant KYBERSWAP_PROXY = 0x818E6FECD516Ecc3849DAf6845e3EC868087B755;

    address public tokenAddress;
    KyberERC20 token;

    KyberERC20 constant internal ETH_TOKEN_ADDRESS = KyberERC20(0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee);

    ILendingPool public lendingPool;
    IUniswapExchange public uniSwap1Exchange;
    IUniswapExchange public exchangeB;
    IUniswapFactory public uniswapOneFactory;

    KyberNetworkProxy public proxy;

    constructor() public {
        // Instantiate Uniswap Factory A
        uniswapOneFactory = IUniswapFactory(UNISWAP_1_FACTORY);
        // get Exchange A Address
        address uni1Exchange_address = uniswapOneFactory.getExchange(DAI_ADDRESS);
        // Instantiate Exchange A
        uniSwap1Exchange = IUniswapExchange(uni1Exchange_address);

        //Instantiate KyberSwap proxy
        proxy = KyberNetworkProxy(KYBERSWAP_PROXY);

        // get lendingPoolAddress
        address lendingPoolAddress = addressesProvider.getLendingPool();
        //Instantiate Aaave Lending Pool B
        lendingPool = ILendingPool(lendingPoolAddress);
    }

    /*
     * Start the arbitrage
     */
    function makeArbitrage(uint256 amount, KyberERC20 _token) public onlyOwner {
        bytes memory data = "";
        tokenAddress = address (_token);
        token = _token;

        // get Exchange uniswap1 Address
        address uni1Exchange_address = uniswapOneFactory.getExchange(tokenAddress);

        // Instantiate Exchange A
        uniSwap1Exchange = IUniswapExchange(uni1Exchange_address);

        lendingPool.flashLoan(address(this), ethAddress, amount, data);

        // Any left amount of Eth is considered profit
        msg.sender.transfer(address(this).balance);

    }

    function executeOperation(
        address _reserve,
        uint256 _amount,
        uint256 _fee,
        bytes calldata _params
    ) override external {
        // If transactions are not mined until deadline the transaction is reverted
        uint256 deadline = getDeadline();

//        KyberERC20 eth = KyberERC20(ETH_TOKEN_ADDRESS);
        token = KyberERC20(tokenAddress);

        // Selling ETH at Exchange A
        require(
            token.approve(address(uniSwap1Exchange), _amount),
            "Could not approve Eth sell"
        );

        uint256 tokenBought = sellEth(_amount, deadline);
        require(
            token.approve(address(proxy), tokenBought),
            "Could not approve token sell"
        );

        execSwap(token, tokenBought, address(this));

        uint256 daiBought = address(this).balance;

        // Repay loan
        uint256 totalDebt = _amount.add(_fee);

        require(daiBought > totalDebt, "Did not profit");

        transferFundsBackToPoolInternal(_reserve, totalDebt);
    }

    function getDeadline() internal view returns (uint256) {
        return now + 3000;
    }

    function sellEth(uint256 _amount, uint256 deadline)public returns (uint256) {
        // Selling ETH at Exchange A
        require(
            token.approve(address(uniSwap1Exchange), _amount),
            "Could not approve Eth sell"
        );

        uint256 tokenBought = uniSwap1Exchange.ethToTokenSwapInput(
            _amount,
            deadline
        );

        return tokenBought;
    }

    //@dev Swap the user's ETH to ERC20 token
    //@param token destination token contract address
    //@param destAddress address to send swapped tokens to
    function execSwap(KyberERC20 _token, uint256 tokenBought, address destAddress) public payable {
        uint minConversionRate;

        // Get the minimum conversion rate
        (minConversionRate,) = proxy.getExpectedRate(KyberERC20(destAddress), ETH_TOKEN_ADDRESS, msg.value);

        // Swap the ERC20 token to ETH
        uint destAmount = proxy.swapTokenToEther(_token, tokenBought, minConversionRate);

        // Send the swapped tokens to the destination address
        require(token.transfer(destAddress, destAmount));

        // Log the event
        Swap(msg.sender, _token, destAmount);
    }
}
