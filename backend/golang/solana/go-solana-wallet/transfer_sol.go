package main

import (
	"context"
	"log"

	"go-solana-wallet/solana-go-sdk/client"
	"go-solana-wallet/solana-go-sdk/common"
	"go-solana-wallet/solana-go-sdk/program/system"
	"go-solana-wallet/solana-go-sdk/rpc"
	"go-solana-wallet/solana-go-sdk/types"

	"fmt"
)

// 3mBuMkK2JiFXfqkUZRwDzRJ1nkiHkBFizMBTsyc2vd2V
var feePayer, _ = types.AccountFromBase58("3Ug3X8gr2NAaf1mh6Uftexs28XwPdYWyVUsy9bZaKPnUeevCcNosbc65fzyP4N7RZ4VfKpc8PGXdALnmKeCVu9sj")

// FCTKi654HAgC8Ht5pEAGQfnr2tJKWbQqvTafACbrsgg7
var alice, _ = types.AccountFromBase58("4HVQb1EX8RzWZLFxfyitgCDC1mjMdtoQrxreEa5YnZzvibo646harbtVtRWP68zDKrmvxHThqjTpz1WP2o86D7wM")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// to fetch recent blockhash
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	balance, err := c.GetBalance(
		context.TODO(), // request context
		"3mBuMkK2JiFXfqkUZRwDzRJ1nkiHkBFizMBTsyc2vd2V", // wallet to fetch balance for
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Println("balance:", balance)

	balance, err = c.GetBalance(
		context.TODO(), // request context
		"FCTKi654HAgC8Ht5pEAGQfnr2tJKWbQqvTafACbrsgg7", // wallet to fetch balance for
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Println("balance:", balance)

	// create a transfer tx
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   feePayer.PublicKey,
					To:     common.PublicKeyFromString("44NR6CrhUEhh7eiVKxF6nawBaokz1ATTM5s31k1vGjcy"),
					Amount: 1e8, // 0.1 SOL
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	// send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
}
