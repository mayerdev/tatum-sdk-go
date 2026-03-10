package transactions_test

import (
	"context"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/testutil"
	"gitlab.com/mayerdev/tatum-sdk-go/transactions"
)

func TestGetHistory(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/transactions", 200, `[{"hash":"0xabc","blockNumber":100,"from":"0x1","to":"0x2","value":"1000","chain":"ethereum"}]`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := transactions.NewService(c)

	txs, err := svc.GetHistory(context.Background(), transactions.HistoryRequest{Address: "0x1", Chain: "ethereum"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(txs) != 1 || txs[0].Hash != "0xabc" {
		t.Errorf("unexpected result: %+v", txs)
	}
}

func TestGetByHash_RateLimit(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/transactions/hash", 429, `{"errorCode":"RATE_LIMIT","message":"rate limit exceeded"}`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL, MaxRetries: 0})
	svc := transactions.NewService(c)

	_, err := svc.GetByHash(context.Background(), transactions.ByHashRequest{Hash: "0xabc", Chain: "ethereum"})
	if err == nil {
		t.Fatal("expected error")
	}
}
