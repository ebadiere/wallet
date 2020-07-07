package kyber

import (
	_ "crypto/ecdsa"
	_ "github.com/ebadiere/wallet/client"
	walletClient "github.com/ebadiere/wallet/client"
	kyber "github.com/ebadiere/wallet/contracts"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

func TokenToToken(rpcUrl string,
	sourceToken string,
	amount float64,
	destToken string) (*big.Float, *big.Float) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, ctx := walletClient.Connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}

	kyberAddress := common.HexToAddress(os.Getenv("kyberNetworkProxyAddress"))
	kyberProxy, err := kyber.NewKyber(kyberAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	// Call Get ExpectedRate here
	return kyberProxy.GetExpectedRate(common.HexToAddress(sourceToken), common.HexToAddress(destToken), walletClient.FloatToBigInt(amount)), nil

}
