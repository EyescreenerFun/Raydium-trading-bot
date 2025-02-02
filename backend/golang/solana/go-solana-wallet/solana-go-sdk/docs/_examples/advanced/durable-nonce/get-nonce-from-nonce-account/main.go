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

	nonceAccountAddr := "DJyNpXgggw1WGgjTVzFsNjb3fuQZVMqhoakvSBfX9LYx"
	nonce, err := c.GetNonceFromNonceAccount(context.Background(), nonceAccountAddr)
	if err != nil {
		log.Fatalf("failed to get nonce account, err: %v", err)
	}

	fmt.Println("nonce", nonce)
}
