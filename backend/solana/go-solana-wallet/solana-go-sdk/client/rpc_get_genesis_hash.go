package client

import (
	"context"

	"go-solana-wallet/solana-go-sdk/rpc"
)

// GetGenesisHash returns the genesis hash
func (c *Client) GetGenesisHash(ctx context.Context) (string, error) {
	return process(
		func() (rpc.JsonRpcResponse[string], error) {
			return c.RpcClient.GetGenesisHash(ctx)
		},
		forward[string],
	)
}
