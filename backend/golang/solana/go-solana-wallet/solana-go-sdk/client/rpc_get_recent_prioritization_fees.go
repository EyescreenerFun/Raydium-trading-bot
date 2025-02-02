package client

import (
	"context"

	"go-solana-wallet/solana-go-sdk/common"
	"go-solana-wallet/solana-go-sdk/rpc"
)

// GetRecentPrioritizationFees returns a list of prioritization fees from recent blocks.
func (c *Client) GetRecentPrioritizationFees(ctx context.Context, addresses []common.PublicKey) (rpc.PrioritizationFees, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.PrioritizationFees], error) {
			a := make([]string, 0, len(addresses))
			for _, address := range addresses {
				a = append(a, address.ToBase58())
			}
			return c.RpcClient.GetRecentPrioritizationFees(ctx, a)
		},
		forward[rpc.PrioritizationFees],
	)
}
