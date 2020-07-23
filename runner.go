package main

import (
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	"github.com/ebadiere/wallet/kyber"
	uni1 "github.com/ebadiere/wallet/uniswapone"
	"log"
)

type ERC20 struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Id       string `json:"id"`
	MakerFee string `json:"maker_fee"`
	TakerFee string `json:"taker_fee"`
}

func main() {

	daiAddr := "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	dai := ERC20{
		Name:     "ERC20 Stablecoin",
		Symbol:   "DAI",
		Id:       "0x2a1530C4C41db0B0b2bB646CB5Eb1A67b7158667",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	wethAddr := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	weth := ERC20{
		Name:     "Wrapped Ether",
		Symbol:   "WETH",
		Id:       "0xA2881A90Bf33F03E7a3f803765Cd2ED5c8928dFb",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	client, ctx := walletClient.Connect()
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		fmt.Println("Connection Failed")
		log.Fatal("Connection Failed")
	}

	fmt.Println("dai: ", dai)
	fmt.Println("weth: ", weth)
	fmt.Println("wethAddr: ", wethAddr)

	exchanges := uni1.LoadResponse()
	count := 0
	for k, v := range exchanges {
		count++
		//if v.Name == dai.Name{
		//	fmt.Println("Found stable dai! ")
		//	continue
		//}
		fmt.Println("token address:", k)
		fmt.Println("token Name:", v.Name)
		fmt.Println("token Symbol:", v.Symbol)
		fmt.Println("Exchange Address :", v.Id)
		fmt.Println("maker fee:", v.MakerFee)
		fmt.Println("taker fee:", v.TakerFee)
		fmt.Println("Exchange count: ", count)
		//Commented out but keep
		//To run later
		//	uni1.CalculateEthToTokenTrade(client,
		//		ctx,
		//		1,
		//		v.Id,
		//		k)
		//}

		uni1TokenAmount := uni1.CalculateTokenToTokenTrade(
			client,
			ctx,
			dai.Id,
			daiAddr,
			10000,
			v.Id,
			k)

		fmt.Println("DAI Amount: 10000")
		fmt.Println("Can buy", v.Symbol, "amount of: ", uni1TokenAmount)

		kyberTokenAmount := kyber.TokenToTokenCalc(
			client,
			k,
			uni1TokenAmount,
			daiAddr)
		fmt.Println("DAI Amount back: ", kyberTokenAmount)
	}

}
