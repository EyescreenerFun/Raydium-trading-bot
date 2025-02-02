go test -v ./solana-go-sdk/program/memoprog $TestBuildMemo

go test -v ./test $TestClient_SendTransaction

go test -v ./test $TestClient_GetClusterNodes