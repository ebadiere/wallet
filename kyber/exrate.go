package kyber

import (
	_ "crypto/ecdsa"
	"fmt"
	kyber "github.com/ebadiere/wallet/contracts/Kyber"
	"github.com/ebadiere/wallet/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
)

type rate struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}

func TokenToTokenRate(
	client *ethclient.Client,
	sourceToken string,
	amount float64,
	sourceDecimals uint8,
	destToken string) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {

	kyberAddress := common.HexToAddress(os.Getenv("kyberNetworkProxyAddress"))
	kyberProxy, err := kyber.NewKyber(kyberAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	amountBigIn := utils.ToWei(amount, int(sourceDecimals))
	// Call Get ExpectedRate here
	fmt.Println("Amount BigIn: ", amountBigIn)

	auth := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	return kyberProxy.GetExpectedRate(&auth, common.HexToAddress(sourceToken), common.HexToAddress(destToken), amountBigIn)
}

func TokenToTokenCalc(
	client *ethclient.Client,
	sourceTokenAddress string,
	sourceTokenAmount float64,
	sourceTokenDecimals uint8,
	destTokenAddress string,
) float64 {

	rate, err := TokenToTokenRate(
		client,
		sourceTokenAddress,
		sourceTokenAmount,
		sourceTokenDecimals,
		destTokenAddress)

	if err != nil {
		log.Fatal("Connection Failed", err)
	}

	// Get the decimal numbers from erc20

	fmt.Println("Expected Rate: ", rate.ExpectedRate)

	sourceAmount := decimal.NewFromFloat(sourceTokenAmount)
	//sourceAmount := utils.ToWei(sourceTokenAmount, int(sourceTokenDecimals))
	// use slippage rate for now
	slipRate := utils.ToDecimal(rate.SlippageRate, 18)
	fmt.Println("Slippage: ", slipRate)

	tokens, _ := sourceAmount.Mul(slipRate).Float64()

	return tokens

}
