package nft_test

import (
	"context"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/testutil"
	"gitlab.com/mayerdev/tatum-sdk-go/nft"
)

func TestGetMetadata(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/metadata", 200, `{"tokenId":"1","contract":"0xabc","chain":"ethereum","owner":"0x1"}`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := nft.NewService(c)

	meta, err := svc.GetMetadata(context.Background(), nft.MetadataRequest{TokenID: "1", Contract: "0xabc", Chain: "ethereum"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.TokenID != "1" {
		t.Errorf("unexpected token id: %s", meta.TokenID)
	}
}

func TestGetMetadata_MalformedJSON(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v4/data/metadata", 200, `{invalid json`)

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := nft.NewService(c)

	_, err := svc.GetMetadata(context.Background(), nft.MetadataRequest{})
	if err == nil {
		t.Fatal("expected error for malformed JSON")
	}
}
