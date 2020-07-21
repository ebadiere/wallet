package uniswapone

import (
	"fmt"
	"testing"
)

type ERC20 struct {
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
	dai := ERC20{
		Name:     "ERC20 Stablecoin",
		Symbol:   "DAI",
		Id:       "0x2a1530C4C41db0B0b2bB646CB5Eb1A67b7158667",
		MakerFee: "0",
		TakerFee: "0.003",
	}

	nmrAddr := "0x1776e1F26f98b1A5dF9cD347953a26dd3Cb46671"
	nmr := ERC20{
		Name:     "Numeraire",
		Symbol:   "NMR",
		Id:       "0x2Bf5A5bA29E60682fC56B2Fcf9cE07Bef4F6196f",
		MakerFee: "0",
		TakerFee: "0.003",
	}
	fmt.Println("Whoohoo")
}
