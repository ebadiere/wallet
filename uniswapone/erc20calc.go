package uniswapone

import (
	"context"
	"encoding/json"
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	"github.com/ebadiere/wallet/contracts/token/ERC20"
	_ "github.com/ebadiere/wallet/contracts/token/ERC20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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

	ethReserve := walletClient.BigIntToFloat(ethReserveBal)
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

	tokenReserve := walletClient.BigIntToFloat(tokenReserveBal)

	tokenAmount := sellEthForTokenAmount(ethAmount, ethReserve, tokenReserve)

	fmt.Println("Amount: ", ethAmount)
	fmt.Println("Eth reserve Amount: ", ethReserveBal)
	fmt.Println("Token reserve Amount: ", tokenReserveBal)
	fmt.Println("Can buy token amount of: ", tokenAmount)
}

func sellEthForTokenAmount(ethAmount float64, ethReserve float64, tokenReserveBal float64) float64 {

	numerator := ethAmount * tokenReserveBal * 997
	denominator := (ethReserve * 1000) + (ethAmount * 997)
	tokenAmount := numerator / denominator

	return tokenAmount
}
