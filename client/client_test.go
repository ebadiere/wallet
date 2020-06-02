package client

import (
	"testing"
)

func TestClientConnection(t *testing.T) {

	rpcUrl := "http://127.0.0.1:7545"

	client, ctx := connect(rpcUrl)
	chainID, _ := client.ChainID(ctx)
	if chainID == nil {
		t.Error("Connection Failed")
	}
}
