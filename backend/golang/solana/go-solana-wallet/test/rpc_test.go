package test

import (
	"context"
	"fmt"
	"testing"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/rpc"
)

	
func TestClient_GetClusterNodes(t *testing.T) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	nodes, err := c.GetClusterNodes(context.Background())
	if err != nil {
		t.Fatalf("failed to get cluster nodes, err: %v", err)
	}
	fmt.Println(nodes)
	fmt.Println(len(nodes))
}

func TestClient_GetMultipleAccounts(t *testing.T) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	accounts, err := c.GetMultipleAccounts(context.Background(),
		[]string{
			"3mBuMkK2JiFXfqkUZRwDzRJ1nkiHkBFizMBTsyc2vd2V",
			"FCTKi654HAgC8Ht5pEAGQfnr2tJKWbQqvTafACbrsgg7",
		},
	)
	if err != nil {
		t.Fatalf("failed to get multiple accounts, err: %v", err)
	}
	fmt.Println(accounts)
	for _, account := range accounts {
		fmt.Println("lamports:", account.Lamports)
		fmt.Println("owner:", account.Owner)
		fmt.Println("executable:", account.Executable)
		fmt.Println("rentEpoch:", account.RentEpoch)
	}
	fmt.Println(len(accounts))
}