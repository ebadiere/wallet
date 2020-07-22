package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	walletClient "github.com/ebadiere/wallet/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"

	arb "github.com/ebadiere/wallet/contracts/arbitrage"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	rpcUrl := os.Getenv("rpcUrl")
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, depErr := arb.DeployArbitrage(auth, client)
	if depErr != nil {
		log.Fatal(depErr)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
