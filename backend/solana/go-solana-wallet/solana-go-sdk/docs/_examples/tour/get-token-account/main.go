package main

import (
	"context"
	"fmt"
	"log"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/program/token"
	"go-solana-wallet/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// token account address
	getAccountInfoResponse, err := c.GetAccountInfo(context.TODO(), "HeCBh32JJ8DxcjTyc6q46tirHR8hd2xj3mGoAcQ7eduL")
	if err != nil {
		log.Fatalf("failed to get account info, err: %v", err)
	}

	tokenAccount, err := token.TokenAccountFromData(getAccountInfoResponse.Data)
	if err != nil {
		log.Fatalf("failed to parse data to a token account, err: %v", err)
	}

	fmt.Printf("%+v\n", tokenAccount)
	// {Mint:F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo Owner:9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde Amount:100000000 Delegate:<nil> State:1 IsNative:<nil> DelegatedAmount:0 CloseAuthority:<nil>}
}
