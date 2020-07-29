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

	kncAddr := "0xdd974D5C2e2928deA5F71b9825b8b646686BD200"
	knc := LISTING{
		Name:     "Kyber Network Crystal",
		Symbol:   "KNC",
		Id:       "0x49c4f9bc14884f6210F28342ceD592A633801a8b",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	ubtAddr := "0x8400D94A5cb0fa0D041a3788e395285d61c9ee5e"
	ubt := LISTING{
		Name:     "UniBright",
		Symbol:   "UBT",
		Id:       "0xfc96e234d4B31C63051E707105fCC4aba37807Fa",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	amplAddr := "0xD46bA6D942050d489DBd938a2C909A5d5039A161"
	ampl := LISTING{
		Name:     "Ampleforth",
		Symbol:   "AMPL",
		Id:       "0x042dBBDc27F75d277C3D99efE327DB21Bc4fde75",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	client, ctx := walletClient.Connect()
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		fmt.Println("Connection Failed")
		log.Fatal("Connection Failed")
	}
	tokenAmount, _ := CalculateTokenToTokenTrade(
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

	tokenAmount, _ = CalculateTokenToTokenTrade(
		client,
		ctx,
		knc.Id,
		kncAddr,
		10000,
		dgx.Id,
		dgxAddr)

	fmt.Println("KNC Amount: 10000")
	fmt.Println("Can buy", knc.Symbol, " DAI amount of: ", tokenAmount)
	fmt.Println("Whoohoo")

	tokenAmount, _ = CalculateTokenToTokenTrade(
		client,
		ctx,
		dai.Id,
		daiAddr,
		10000,
		ubt.Id,
		ubtAddr)

	fmt.Println("DAI Amount: 10000")
	fmt.Println("Can buy", ubt.Symbol, " DAI amount of: ", tokenAmount)
	fmt.Println("Whoohoo")

	tokenAmount, _ = CalculateTokenToTokenTrade(
		client,
		ctx,
		dai.Id,
		daiAddr,
		10000,
		ampl.Id,
		amplAddr)

	fmt.Println("DAI Amount: 10000")
	fmt.Println("Can buy", ampl.Symbol, " DAI amount of: ", tokenAmount)
	fmt.Println("Whoohoo")

}
