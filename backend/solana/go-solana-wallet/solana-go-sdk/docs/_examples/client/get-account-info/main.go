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

	// get account info
	accountInfo, err := c.GetAccountInfo(
		context.TODO(),
		"F5RYi7FMPefkc7okJNh21Hcsch7RUaLVr8Rzc8SQqxUb",
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("accountInfo: %v\n", accountInfo)
}
