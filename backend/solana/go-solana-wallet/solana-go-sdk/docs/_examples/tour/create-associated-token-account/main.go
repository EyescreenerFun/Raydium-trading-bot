package main

import (
	"context"
	"fmt"
	"log"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/common"
	"go-solana-wallet/solana-go-sdk/program/associated_token_account"
	"go-solana-wallet/solana-go-sdk/rpc"
	"go-solana-wallet/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

var mintPubkey = common.PublicKeyFromString("F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	ata, _, err := common.FindAssociatedTokenAddress(alice.PublicKey, mintPubkey)
	if err != nil {
		log.Fatalf("find ata error, err: %v", err)
	}
	fmt.Println("ata:", ata.ToBase58())

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				associated_token_account.Create(associated_token_account.CreateParam{
					Funder:                 feePayer.PublicKey,
					Owner:                  alice.PublicKey,
					Mint:                   mintPubkey,
					AssociatedTokenAccount: ata,
				}),
			},
		}),
		Signers: []types.Account{feePayer},
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}
