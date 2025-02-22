package rpc

import (
	"context"
	"testing"

	"go-solana-wallet/solana-go-sdk/internal/client_test"
)

func TestGetProgramAccounts(t *testing.T) {
	client_test.TestAll(
		t,
		[]client_test.Param{
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"},{"account":{"data":["7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q"}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetProgramAccounts(
						context.TODO(),
						"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
					)
				},
				ExpectedValue: JsonRpcResponse[GetProgramAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetProgramAccounts{
						{
							Pubkey: "9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
							Account: AccountInfo{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []any{
									"0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
									"base64",
								},
								Executable: false,
							},
						},
						{
							Pubkey: "4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q",
							Account: AccountInfo{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []any{
									"7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
									"base64",
								},
								Executable: false,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"memcmp": {"offset": 0, "bytes": "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq"}}]}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetProgramAccountsWithConfig(
						context.TODO(),
						"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						GetProgramAccountsConfig{
							Encoding: AccountEncodingBase64Zstd,
							Filters: []GetProgramAccountsConfigFilter{
								{
									MemCmp: &GetProgramAccountsConfigFilterMemCmp{
										Offset: 0,
										Bytes:  "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq",
									},
								},
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetProgramAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetProgramAccounts{
						{
							Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
							Account: AccountInfo{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []any{
									"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
									"base64+zstd",
								},
								Executable: false,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"dataSize": 165}]}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetProgramAccountsWithConfig(
						context.TODO(),
						"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						GetProgramAccountsConfig{
							Encoding: AccountEncodingBase64Zstd,
							Filters: []GetProgramAccountsConfigFilter{
								{
									DataSize: 165,
								},
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetProgramAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetProgramAccounts{
						{
							Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
							Account: AccountInfo{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []any{
									"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
									"base64+zstd",
								},
								Executable: false,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"encoding": "base64+zstd", "filters":[{"memcmp": {"offset": 0, "bytes": "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq"}}, {"dataSize": 165}]}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":[{"account":{"data":["KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=","base64+zstd"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6"}],"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetProgramAccountsWithConfig(
						context.TODO(),
						"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
						GetProgramAccountsConfig{
							Encoding: AccountEncodingBase64Zstd,
							Filters: []GetProgramAccountsConfigFilter{
								{
									MemCmp: &GetProgramAccountsConfigFilterMemCmp{
										Offset: 0,
										Bytes:  "GwktiL5dEA4mCGeVdGyUPDeWKzGC486KkPVTbwwnMGYq",
									},
								},
								{
									DataSize: 165,
								},
							},
						},
					)
				},
				ExpectedValue: JsonRpcResponse[GetProgramAccounts]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetProgramAccounts{
						{
							Pubkey: "Dh4w3Pn6HqCDbEDhZdcDY8bHydeqNAhYY6EktLiWxFf6",
							Account: AccountInfo{
								Lamports:  2039280,
								Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
								RentEpoch: 181,
								Data: []any{
									"KLUv/QBYbQIANATs5kvExA7d9DwekfxJB78GIvCdh5i2DIdwRkIs0WfLghCWWRdefGQzIaXtRkKgJ7Cr2XuN2XrRvMbcZHE4bM3cAAEAAgAEJwaY4Aw=",
									"base64+zstd",
								},
								Executable: false,
							},
						},
					},
				},
				ExpectedError: nil,
			},
			{
				RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getProgramAccounts", "params":["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", {"withContext": true}]}`,
				ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":78496860},"value":[{"account":{"data":["0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D"},{"account":{"data":["7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","base64"],"executable":false,"lamports":2039280,"owner":"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA","rentEpoch":181},"pubkey":"4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q"}]},"id":1}`,
				F: func(url string) (any, error) {
					c := NewRpcClient(url)
					return c.GetProgramAccountsWithContext(
						context.TODO(),
						"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
					)
				},
				ExpectedValue: JsonRpcResponse[GetProgramAccountsWithContext]{
					JsonRpc: "2.0",
					Id:      1,
					Error:   nil,
					Result: GetProgramAccountsWithContext{
						Context: Context{
							Slot: 78496860,
						},
						Value: GetProgramAccounts{
							{
								Pubkey: "9ywX3U33UZC1HThhoBR2Ys7SiouXDkkDoH6brJApFh5D",
								Account: AccountInfo{
									Lamports:  2039280,
									Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch: 181,
									Data: []any{
										"0SWx++406Gemp6iqPyXxkCKUGHf4+wT/Ycjdf33fscwGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
										"base64",
									},
									Executable: false,
								},
							},
							{
								Pubkey: "4MdGG2EnnAp4dhbMCDww1qLEEiW5SHfUAe9U9RtTqS8q",
								Account: AccountInfo{
									Lamports:  2039280,
									Owner:     "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
									RentEpoch: 181,
									Data: []any{
										"7OZLxMQO3fQ8HpH8SQe/BiLwnYeYtgyHcEZCLNFny4IGPnDZkVeY3KPu25E3z2THD/r7Ys3TgoEFJQvXax1s3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
										"base64",
									},
									Executable: false,
								},
							},
						},
					},
				},
				ExpectedError: nil,
			},
		},
	)
}
