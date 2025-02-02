package client

import (
	"context"
	"testing"

	"go-solana-wallet/solana-go-sdk/internal/client_test"
	"go-solana-wallet/solana-go-sdk/pkg/pointer"
	"go-solana-wallet/solana-go-sdk/rpc"
)

func TestClient_GetSlotLeader(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader"}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlotLeader(context.TODO())
				},
				ExpectedValue: "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				ExpectedError: nil,
			},
		},
	)
}

func TestClient_GetSlotLeaderWithConfig(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"commitment": "processed"}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							Commitment: pointer.Get(rpc.CommitmentProcessed),
						},
					)
				},
				ExpectedValue: "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"minContextSlot":0}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							MinContextSlot: pointer.Get[uint64](0),
						},
					)
				},
				ExpectedValue: "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getSlotLeader","params":[{"commitment": "confirmed","minContextSlot":10}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":"q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot","id":1}`,
				F: func(url string) (any, error) {
					c := NewClient(url)
					return c.GetSlotLeaderWithConfig(
						context.TODO(),
						GetSlotLeaderConfig{
							Commitment:     pointer.Get(rpc.CommitmentConfirmed),
							MinContextSlot: pointer.Get[uint64](10),
						},
					)
				},
				ExpectedValue: "q9XWcZ7T1wP4bW9SB4XgNNwjnFEJ982nE8aVbbNuwot",
				ExpectedError: nil,
			},
		},
	)
}
