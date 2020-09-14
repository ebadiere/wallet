package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	arb "github.com/ebadiere/wallet/contracts/uniswap2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	client, _ := walletClient.Connect()

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

	gasPrice, e := client.SuggestGasPrice(context.Background())
	if e != nil {
		log.Fatal(e)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1900000) // in units
	//auth.GasPrice = gasPrice
	gas := big.NewInt(101000000000)
	auth.GasPrice = gas

	fmt.Println("Gas Price: ", gasPrice)

	address, tx, instance, depErr := arb.DeployUniswap2(
		auth,
		client,
		common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x80fB784B7eD66730e8b1DBd9820aFD29931aab03"))

	if depErr != nil {
		log.Fatal(depErr)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
