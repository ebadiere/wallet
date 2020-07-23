package kyber

import (
	"fmt"
	//"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	_ "github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/stretchr/testify/assert"
	_ "math/big"
	"testing"
)

func TestTokenToToken(t *testing.T) {

	client, ctx := walletClient.Connect()
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}

	rate, err := TokenToTokenRate(client, "0x1F573D6Fb3F13d689FF844B4cE37794d79a7FF1C", 100, "0x6B175474E89094C44Da98b954EedeAC495271d0F")
	if err != nil {
		t.Error("Contract failed")
	}

	fmt.Println("Expected Rate: ", rate.ExpectedRate)
	fmt.Println("Slippage: ", rate.SlippageRate)
}
