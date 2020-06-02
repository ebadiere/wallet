package client

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

}

func connect(rpcUrl string) (*ethclient.Client, context.Context) {

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal("Connection to node failed! ", err)
	}

	ctx := context.Background()

	return client, ctx
}
