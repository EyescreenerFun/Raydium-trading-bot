package main

import (
	"context"
	"fmt"
	"log"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/common"
	"go-solana-wallet/solana-go-sdk/program/address_lookup_table"
	"go-solana-wallet/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	lookupTablePubkey := common.PublicKeyFromString("3LZbwptsCkv5R5uu1GNZKiX9SoC6egNG8NXg9zH5ZVM9")

	accountInfo, err := c.GetAccountInfo(context.Background(), lookupTablePubkey.ToBase58())
	if err != nil {
		log.Fatalf("failed to get account info, err: %v", err)
	}
	fmt.Println(accountInfo.Data)

	addressLookupTable, err := address_lookup_table.DeserializeLookupTable(accountInfo.Data, accountInfo.Owner)
	if err != nil {
		log.Fatalf("failed to deserialized account, err: %v", err)
	}

	fmt.Printf("%+v\n", addressLookupTable)
}
