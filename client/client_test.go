package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestSend(t *testing.T) {

	fmt.Println("Hello")

	client, ctx := getClientConnection(t)
	privateKey, err := crypto.HexToECDSA("2dd5ee8e6cd77d432f803612b819dbb0e0447ff98a5d27996b6ce7e60dc09d3b")
	if err != nil {
		log.Fatal(err)
	}

	send(client, ctx, privateKey, 10)

}

func getClientConnection(t *testing.T) (*ethclient.Client, context.Context) {
	rpcUrl := "http://127.0.0.1:7545"

	client, ctx := connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}

	return client, ctx
}
