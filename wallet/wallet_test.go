package wallet_test

import (
	"context"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/testutil"
	"gitlab.com/mayerdev/tatum-sdk-go/wallet"
)

func TestGetPortfolio(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/portfolio", 200, `[{"address":"0x1","chain":"ethereum","balance":"1.0","asset":"ETH","price":2000}]`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := wallet.NewService(c)

	items, err := svc.GetPortfolio(context.Background(), wallet.PortfolioRequest{Addresses: []string{"0x1"}, Chain: "ethereum"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].Address != "0x1" {
		t.Errorf("unexpected address: %s", items[0].Address)
	}
}

func TestGetPortfolio_NotFound(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/portfolio", 404, `{"errorCode":"NOT_FOUND","message":"not found"}`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := wallet.NewService(c)

	_, err := svc.GetPortfolio(context.Background(), wallet.PortfolioRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
}
