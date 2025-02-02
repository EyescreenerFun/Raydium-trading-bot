package client

import (
	"context"

	"go-solana-wallet/solana-go-sdk/rpc"
)

// GetBlockTime returns the estimated production time of a block.
func (c *Client) GetBlockTime(ctx context.Context, slot uint64) (*int64, error) {
	return process(
		func() (rpc.JsonRpcResponse[*int64], error) {
			return c.RpcClient.GetBlockTime(ctx, slot)
		},
		forward[*int64],
	)
}
