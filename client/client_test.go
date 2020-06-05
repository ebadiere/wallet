package client

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestSend(t *testing.T) {

	client, ctx := getClientConnection(t)
	privateKey, err := crypto.HexToECDSA("dc52646c3b180710fee14723f72197d94344db31790f085398d2cf61e29a8102")
	if err != nil {
		log.Fatal(err)
	}

	Send(client, ctx, privateKey, "0x0c7338Bf194fb0ea44afFe608974288EE55E869C", 10)

	// Send Back

}

func getClientConnection(t *testing.T) (*ethclient.Client, context.Context) {
	rpcUrl := "http://127.0.0.1:7545"

	client, ctx := Connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}

	return client, ctx
}
