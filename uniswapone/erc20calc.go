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

	amount := walletClient.FloatToBigInt(ethAmount)
	//exchangeAddress := common.HexToAddress(tokenExchange)

	ethReserveBal, err := client.BalanceAt(ctx, common.HexToAddress(tokenExchange), nil)
	if err != nil {
		log.Fatal(err)
	}

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

	//tokenReserveBal, err := client.BalanceAt(ctx, common.HexToAddress(tokenExchange), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("Amount: ", amount)
	fmt.Println("Eth reserve Amount: ", ethReserveBal)
	fmt.Println("Token reserve Amount: ", tokenReserveBal)
	//inputReserve := use client here
}
