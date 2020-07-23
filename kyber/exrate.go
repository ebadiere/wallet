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
	destToken string) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {

	kyberAddress := common.HexToAddress(os.Getenv("kyberNetworkProxyAddress"))
	kyberProxy, err := kyber.NewKyber(kyberAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	amountBigIn := utils.ToWei(amount, 18)
	// Call Get ExpectedRate here

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
	destTokenAddress string) float64 {

	rate, err := TokenToTokenRate(
		client,
		sourceTokenAddress,
		sourceTokenAmount,
		destTokenAddress)

	if err != nil {
		log.Fatal("Connection Failed", err)
	}

	fmt.Println("Expected Rate: ", rate.ExpectedRate)
	fmt.Println("Slippage: ", rate.SlippageRate)

	sourceAmount := utils.ToDecimal(sourceTokenAmount, 18)
	// use slippage rate for now
	slipRate := utils.ToDecimal(rate.SlippageRate, 18)
	tokens, _ := sourceAmount.Mul(slipRate).Float64()

	return tokens

}
