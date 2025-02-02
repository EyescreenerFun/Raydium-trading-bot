package client

import (
	"context"
	"testing"

	"go-solana-wallet/solana-go-sdk/internal/client_test"
	"go-solana-wallet/solana-go-sdk/pkg/pointer"
	"go-solana-wallet/solana-go-sdk/rpc"
)

func TestClient_GetVersion(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getVersion"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"feature-set":1824749018,"solana-core":"1.7.14"},"id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetVersion(
						context.TODO(),
					)
				},
				ExpectedValue: rpc.GetVersion{
					SolanaCore: "1.7.14",
					FeatureSet: pointer.Get[uint32](1824749018),
				},
				ExpectedError: nil,
			},
		},
	)
}
