package main

import (
	"context"
	"fmt"
	"log"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, err := c.GetBalance(
		context.TODO(),
		"9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an",
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}
	fmt.Println(balance)
}
