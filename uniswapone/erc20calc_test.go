package uniswapone

import (
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	"log"
	"testing"
)

type LISTING struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Id       string `json:"id"`
	MakerFee string `json:"maker_fee"`
	TakerFee string `json:"taker_fee"`
}

func TestIterateThruExchanges(t *testing.T) {

	LoadResponse()

}

func TestCalculateTokenToTokenTrade(t *testing.T) {

	//LoadResponse()
	daiAddr := "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	dai := LISTING{
		Name:     "LISTING Stablecoin",
		Symbol:   "DAI",
		Id:       "0x2a1530C4C41db0B0b2bB646CB5Eb1A67b7158667",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	dgxAddr := "0x4f3AfEC4E5a3F2A6a1A411DEF7D7dFe50eE057bF"
	dgx := LISTING{
		Name:     "Digix Gold Token",
		Symbol:   "DGX",
		Id:       "0xb92dE8B30584392Af27726D5ce04Ef3c4e5c9924",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	client, ctx := walletClient.Connect()
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		fmt.Println("Connection Failed")
		log.Fatal("Connection Failed")
	}
	tokenAmount := CalculateTokenToTokenTrade(
		client,
		ctx,
		dai.Id,
		daiAddr,
		10000,
		dgx.Id,
		dgxAddr)

	fmt.Println("DAI Amount: 10000")
	fmt.Println("Can buy", dgx.Symbol, "amount of: ", tokenAmount)
	fmt.Println("Whoohoo")
}
