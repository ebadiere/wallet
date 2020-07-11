package main

import (
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	uni1 "github.com/ebadiere/wallet/uniswapone"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Dai struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Id       string `json:"id"`
	MakerFee string `json:"maker_fee"`
	TakerFee string `json:"taker_fee"`
}

func main() {

	dai := Dai{
		Name:     "Dai Stablecoin",
		Symbol:   "DAI",
		Id:       "0x2a1530C4C41db0B0b2bB646CB5Eb1A67b7158667",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpcUrl := os.Getenv("rpcUrl")
	if err != nil {
		log.Fatal(err)
	}

	client, ctx := walletClient.Connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		fmt.Println("Connection Failed")
		log.Fatal("Connection Failed")
	}

	fmt.Println("dai: ", dai)
	exchanges := uni1.LoadResponse()
	count := 0
	for k, v := range exchanges {
		count++
		//if v.Name == dai.Name{
		//	fmt.Println("Found stable dai! ")
		//	continue
		//}
		//fmt.Println("token address:", k)
		//fmt.Println("token Name:", v.Name)
		//fmt.Println("token Symbol:", v.Symbol)
		//fmt.Println("Exchange Address :", v.Id)
		//fmt.Println("maker fee:", v.MakerFee)
		//fmt.Println("taker fee:", v.TakerFee)
		//fmt.Println("Exchange count: ", count)
		uni1.CalculateEthToTokenTrade(client,
			ctx,
			10000,
			v.Id,
			k)
	}

}
