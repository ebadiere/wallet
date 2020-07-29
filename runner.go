package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	_ "github.com/ebadiere/wallet/contracts/arbitrage"
	"github.com/ebadiere/wallet/kyber"
	uni1 "github.com/ebadiere/wallet/uniswapone"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
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
	for {
		iterateResponse(exchanges, count, client, ctx, dai, daiAddr)
	}

}

func iterateResponse(exchanges map[string]uni1.Exchange, count int, client *ethclient.Client, ctx context.Context, dai ERC20, daiAddr string) {
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
		//Commented out but keep
		//To run later
		//	uni1.CalculateEthToTokenTrade(client,
		//		ctx,
		//		1,
		//		v.Id,
		//		k)
		//}

		uni1TokenAmount, tokenDecimals := uni1.CalculateTokenToTokenTrade(
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
			tokenDecimals,
			daiAddr)
		fmt.Println("DAI Amount back: ", kyberTokenAmount)

		if kyberTokenAmount > 10000 {
			fmt.Println("ARBITRAGE!!!!!!!")
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			privateKey, err := crypto.HexToECDSA(os.Getenv("privateKey"))
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			auth := bind.NewKeyedTransactor(privateKey)
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)     // in wei
			auth.GasLimit = uint64(800000) // in units
			auth.GasPrice = gasPrice

			//address := common.HexToAddress("0x1da69abFa9F4e6cf51B14c8794921F34bB90D03f")
			//arbi, err := arbitrage.NewArbitrage(address, client)
			//if err != nil {
			//	log.Fatal(err)
			//}

			//baseIn := big.NewFloat(10000000000000000000000)
			//decIn := utils.ToDecimal(baseIn, 18)
			//transaction, err := arbi.MakeArbitrage(auth, decimal.New(10000, 18).BigInt(), common.HexToAddress(k))
			//if err != nil {
			//	log.Fatal(err)
			//}
			//fmt.Println("Transaction: ", transaction.Hash().Hex())
		}
	}
}
