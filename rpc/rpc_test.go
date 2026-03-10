package rpc_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/rpc"
)

func TestCall(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ethereum/mainnet/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	gatewayTpl := srv.URL + "/%s/%s/"
	svc := rpc.NewService(c, gatewayTpl)

	result, err := svc.Call(context.Background(), chain.Ethereum, chain.Mainnet, "eth_blockNumber", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var val string
	if err := json.Unmarshal(result, &val); err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}
	if val != "0x1" {
		t.Errorf("unexpected result: %s", val)
	}
}

func TestCall_RPCError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ethereum/mainnet/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"error":{"code":-32600,"message":"invalid request"}}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := httpclient.New(httpclient.Config{BaseURL: srv.URL})
	svc := rpc.NewService(c, srv.URL+"/%s/%s/")

	_, err := svc.Call(context.Background(), chain.Ethereum, chain.Mainnet, "eth_blockNumber", nil)
	if err == nil {
		t.Fatal("expected error")
	}
}
