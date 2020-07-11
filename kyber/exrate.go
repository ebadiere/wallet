package kyber

import (
	_ "crypto/ecdsa"
	walletClient "github.com/ebadiere/wallet/client"
	kyber "github.com/ebadiere/wallet/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

type rate struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}

func TokenToTokenRate(rpcUrl string, sourceToken string, amount float64, destToken string) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	client, ctx := walletClient.Connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		log.Fatal("Connection Failed")
	}

	kyberAddress := common.HexToAddress(os.Getenv("kyberNetworkProxyAddress"))
	kyberProxy, err := kyber.NewKyber(kyberAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	amountBigIn := walletClient.FloatToBigInt(amount)
	// Call Get ExpectedRate here

	auth := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	return kyberProxy.GetExpectedRate(&auth, common.HexToAddress(sourceToken), common.HexToAddress(destToken), amountBigIn)
}
