package uniswapone

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ebadiere/wallet/contracts/token/ERC20"
	_ "github.com/ebadiere/wallet/contracts/token/ERC20"
	"github.com/ebadiere/wallet/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"os"
)

type Exchange struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Id       string `json:"id"`
	MakerFee string `json:"maker_fee"`
	TakerFee string `json:"taker_fee"`
}

func LoadResponse() map[string]Exchange {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(os.Getenv("uniswapOneResp"))
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	exchanges := make(map[string]Exchange)
	json.Unmarshal([]byte(byteValue), &exchanges)

	return exchanges
}

func CalculateEthToTokenTrade(client *ethclient.Client,
	ctx context.Context,
	ethAmount float64,
	tokenExchange string,
	tokenAddr string) {

	ethReserveBal, err := client.BalanceAt(ctx, common.HexToAddress(tokenExchange), nil)
	if err != nil {
		log.Fatal(err)
	}

	//ethReserve := walletClient.BigIntToFloat(ethReserveBal)
	ethReserve := utils.ToDecimal(ethReserveBal, 18)
	fmt.Println("ethReserve: ", ethReserve)

	erc20, err := ERC20.NewERC20(common.HexToAddress(tokenAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	tokenReserveBal, err := erc20.BalanceOf(&auth, common.HexToAddress(tokenExchange))
	if err != nil {
		log.Fatal(err)
	}

	//tokenReserve := walletClient.BigIntToFloat(tokenReserveBal)
	tokenReserve := utils.ToDecimal(tokenReserveBal, 18)

	tokenAmount := sellEthForTokenAmount(decimal.NewFromFloat(ethAmount), ethReserve, tokenReserve)

	fmt.Println("Amount: ", ethAmount)
	fmt.Println("Eth reserve Amount: ", ethReserveBal)
	fmt.Println("Token reserve Amount: ", tokenReserveBal)
	fmt.Println("Can buy token amount of: ", tokenAmount)
}

func CalculateTokenToTokenTrade(
	client *ethclient.Client,
	ctx context.Context,
	inputTokenExchange string,
	inputTokenAddress string,
	inputTokenAmount float64,
	outputTokenExchange string,
	outputTokenAddress string) float64 {

	auth := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	erc20Input, err := ERC20.NewERC20(common.HexToAddress(inputTokenAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	erc20Output, err := ERC20.NewERC20(common.HexToAddress(outputTokenAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	inputTokenReserveBal, err := erc20Input.BalanceOf(&auth, common.HexToAddress(inputTokenExchange))
	if err != nil {
		log.Fatal(err)
	}

	//inputReserve := walletClient.BigIntToFloat(inputTokenReserveBal)
	//wei := new(big.Int)
	//inputReserve, _ := wei.SetString(inputTokenReserveBal.String(), 18)
	inputReserve := utils.ToDecimal(inputTokenReserveBal, 18)
	fmt.Println("DAI tokenReserve: ", inputReserve)

	inputEthReserveBal, err := client.BalanceAt(ctx, common.HexToAddress(inputTokenExchange), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Eth Reserve: ", inputEthReserveBal)

	//ethReserve := walletClient.BigIntToFloat(inputEthReserveBal)
	ethReserve := utils.ToDecimal(inputEthReserveBal, 18)

	fmt.Println("ethReserve: ", ethReserve)

	combinedFee := inputTokenAmount * 0.00591
	amountMinusFee := inputTokenAmount - combinedFee
	fmt.Println("Fee: ", amountMinusFee)

	ethAmount := sellTokenForEth(decimal.NewFromFloat(amountMinusFee), ethReserve, inputReserve)
	fmt.Println("ethAmount: ", ethAmount)

	// next steps Eth to output token
	outputEthReserveBal, err := client.BalanceAt(ctx, common.HexToAddress(outputTokenExchange), nil)
	if err != nil {
		log.Fatal(err)
	}

	//outputEthReserve := walletClient.BigIntToFloat(outputEthReserveBal)
	outputEthReserve := utils.ToDecimal(outputEthReserveBal, 18)
	fmt.Println("ethReserve: ", outputEthReserve)

	outputTokenReserveBal, err := erc20Output.BalanceOf(&auth, common.HexToAddress(outputTokenExchange))
	if err != nil {
		log.Fatal(err)
	}

	//outputReserve := walletClient.BigIntToFloat(outputTokenReserveBal)
	outputReserve := utils.ToDecimal(outputTokenReserveBal, 18)
	fmt.Println("Output tokenReserve: ", outputReserve)

	tokenAmount, _ := sellEthForTokenAmount(ethAmount, outputEthReserve, outputReserve).Float64()

	return tokenAmount

}

func sellEthForTokenAmount(ethAmount decimal.Decimal, ethReserve decimal.Decimal, tokenReserveBal decimal.Decimal) decimal.Decimal {

	numerator := ethAmount.Mul(tokenReserveBal).Mul(decimal.NewFromInt(997))
	denominator := ethReserve.Mul(decimal.NewFromInt(1000)).Add(ethAmount.Mul(decimal.NewFromInt(997)))
	tokenAmount := numerator.Div(denominator)

	return tokenAmount
}

func sellTokenForEth(tokenAmount decimal.Decimal, ethReserve decimal.Decimal, inputReserve decimal.Decimal) decimal.Decimal {

	numerator := tokenAmount.Mul(ethReserve).Mul(decimal.NewFromInt(997))
	denominator := inputReserve.Mul(decimal.NewFromInt(1000)).Add(tokenAmount.Mul(decimal.NewFromInt(997)))
	outputEth := numerator.Div(denominator)

	return outputEth

}
