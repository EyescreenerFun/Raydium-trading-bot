package main

import (
        "context"
        "fmt"
        "go-solana-wallet/solana-go-sdk/client"
        "go-solana-wallet/solana-go-sdk/rpc"
        "go-solana-wallet/solana-go-sdk/types"
        "github.com/mr-tron/base58"
        "go-solana-wallet/solana-go-sdk/common"
        "go-solana-wallet/solana-go-sdk/program/system"
)

func TransferSol(c *client.Client, wallet *types.Account, PublicKey string ) string {
        // 转账 从网络上检索最新的区块哈希值
        response2, err := c.GetLatestBlockhash(context.TODO())
        fmt.Println("response2", response2.Blockhash)
        // 用检索3到的区块哈希值和交易签字人的公钥制作一个转账信息
        tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        wallet.PublicKey,
			RecentBlockhash: response2.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   wallet.PublicKey,
					To:     common.PublicKeyFromString(PublicKey),
					Amount: 1e8, // 0.1 SOL
				}),
			},
		}),
		Signers: []types.Account{*wallet},
	})
        //将交易发送到网络
        // send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println("failed to send tx, err: %v", err)
	}
        return txhash
}

// acount info
func GetAccountInfo(c *client.Client, wallet *types.Account, PublicKey string ) {
        response, err := c.GetAccountInfo(context.TODO(), PublicKey)
        if err != nil {
                fmt.Println("err", err)
        }
        fmt.Println("response", response)
}

func main() {
        // create a RPC client
        c := client.NewClient(rpc.DevnetRPCEndpoint)

        // get the current running Solana version
        response, err := c.GetVersion(context.TODO())
        if err != nil {
                panic(err)
        }

        fmt.Println("version", response.SolanaCore)

        // create a new wallet using types.NewAccount()
        var wallet types.Account
        if false {
                wallet = types.NewAccount()
        } else {
                // 3mBuMkK2JiFXfqkUZRwDzRJ1nkiHkBFizMBTsyc2vd2V
                wallet, err = types.AccountFromBase58("3Ug3X8gr2NAaf1mh6Uftexs28XwPdYWyVUsy9bZaKPnUeevCcNosbc65fzyP4N7RZ4VfKpc8PGXdALnmKeCVu9sj")
                if err != nil {
                        panic(err)
                }
        }

        // display the wallet public and private keys
        fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
        fmt.Println("Private Key:", wallet.PrivateKey)

        // 直接将私钥（[]byte）进行 base58 编码
        encoded := base58.Encode(wallet.PrivateKey)
        fmt.Println("Private Key Encoded:", encoded)

        // 请求androip
        if false {
                // 请求androip
                // request for 1 SOL airdrop using RequestAirdrop()
                txhash, err := c.RequestAirdrop(
                        context.TODO(), // request context
                        wallet.PublicKey.ToBase58(), // wallet address requesting airdrop
                        1e9, // amount of SOL in lamport
                )
                fmt.Println("txhash", txhash)
                // check for errors
                if err != nil {
                        panic(err)
                }
        }

        // 余额
        // fetch the balance using GetBalance()
        balance, err := c.GetBalance(
                context.TODO(), // request context
                "3mBuMkK2JiFXfqkUZRwDzRJ1nkiHkBFizMBTsyc2vd2V", // wallet to fetch balance for
        )

        // check for errors
        if err != nil {
                panic(err)
        }

        fmt.Println("Wallet Balance in Lamport:", balance)
        fmt.Println("Wallet Balance in SOL:", balance/1e9)


        // 发送SOL
        // txhash := TransferSol(c, &wallet, "44NR6CrhUEhh7eiVKxF6nawBaokz1ATTM5s31k1vGjcy")
	// fmt.Println("txhash:", txhash)

        GetAccountInfo(c, &wallet, "44NR6CrhUEhh7eiVKxF6nawBaokz1ATTM5s31k1vGjcy")
}


