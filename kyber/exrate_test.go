package kyber

import (
	"fmt"
	//"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/stretchr/testify/assert"
	"log"
	_ "math/big"
	"os"
	"testing"
)

func TestTokenToToken(t *testing.T) {

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
		t.Error("Connection Failed")
	}

	rate, err := TokenToTokenRate(rpcUrl, "0x1F573D6Fb3F13d689FF844B4cE37794d79a7FF1C", 100, "0x6B175474E89094C44Da98b954EedeAC495271d0F")
	if chainID == nil {
		t.Error("Connection Failed")
	}

	fmt.Println("Expected Rate: ", rate.ExpectedRate)
	fmt.Println("Slippage: ", rate.SlippageRate)
}
