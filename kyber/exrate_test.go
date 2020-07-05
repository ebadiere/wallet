package kyber

import (
	"fmt"
	 client "github.com/ebadiere/wallet/client"
	"github.com/joho/godotenv"
	_ "github.com/stretchr/testify/assert"
	"log"
	_ "math/big"
	"os"
	"testing"
)

func TestTokenToToken(t *testing.T){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	rpcUrl := os.Getenv("rpcUrl")
	client, ctx := client.Connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}

	res := TokenToToken("0x6B175474E89094C44Da98b954EedeAC495271d0F", 10000, "0xdd974D5C2e2928deA5F71b9825b8b646686BD200");
	fmt.Println(res)
	fmt.Println("rpcUrl: ", rpcUrl)
}
