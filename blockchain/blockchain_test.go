package blockchain_test

import (
	"context"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/blockchain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/testutil"
)

func TestGetCurrentBlock(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/blockchain/block/current", 200, `{"chain":"ethereum","number":18000000,"hash":"0xabc"}`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := blockchain.NewService(c)

	block, err := svc.GetCurrentBlock(context.Background(), blockchain.CurrentBlockRequest{Chain: "ethereum"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if block.Number != 18000000 {
		t.Errorf("unexpected block number: %d", block.Number)
	}
}

func TestGetNativeBalance(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/blockchain/balance", 200, `{"address":"0x1","balance":"5.0"}`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := blockchain.NewService(c)

	bal, err := svc.GetNativeBalance(context.Background(), blockchain.NativeBalanceRequest{Address: "0x1", Chain: "ethereum"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bal.Balance != "5.0" {
		t.Errorf("unexpected balance: %s", bal.Balance)
	}
}
