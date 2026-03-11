package walletgen_test

import (
	"context"
	"errors"
	"testing"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/testutil"
	"gitlab.com/mayerdev/tatum-sdk-go/walletgen"
)

func TestGenerateWallet(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/ethereum/wallet", 200, `{"mnemonic":"word1 word2","xpub":"xpubABC"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	w, err := svc.GenerateWallet(context.Background(), chain.Ethereum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Mnemonic != "word1 word2" {
		t.Errorf("unexpected mnemonic: %s", w.Mnemonic)
	}
	if w.XPub != "xpubABC" {
		t.Errorf("unexpected xpub: %s", w.XPub)
	}
}

func TestGenerateWallet_Solana(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/solana/wallet", 200, `{"address":"SolAddr123","privateKey":"privSol"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	w, err := svc.GenerateWallet(context.Background(), chain.Solana)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Address != "SolAddr123" {
		t.Errorf("unexpected address: %s", w.Address)
	}
	if w.PrivateKey != "privSol" {
		t.Errorf("unexpected privateKey: %s", w.PrivateKey)
	}
}

func TestGenerateWallet_Algorand(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/algorand/wallet", 200, `{"address":"AlgoAddr","secret":"algoSecret"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	w, err := svc.GenerateWallet(context.Background(), chain.Algorand)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.Address != "AlgoAddr" {
		t.Errorf("unexpected address: %s", w.Address)
	}
	if w.Secret != "algoSecret" {
		t.Errorf("unexpected secret: %s", w.Secret)
	}
}

func TestDeriveAddress(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/bitcoin/address/xpubXXX/0", 200, `{"address":"1BTCAddr"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	addr, err := svc.DeriveAddress(context.Background(), chain.Bitcoin, "xpubXXX", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if addr != "1BTCAddr" {
		t.Errorf("unexpected address: %s", addr)
	}
}

func TestDeriveAddress_Unsupported(t *testing.T) {
	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: "http://unused"}))

	_, err := svc.DeriveAddress(context.Background(), chain.Solana, "xpubXXX", 0)
	if !errors.Is(err, walletgen.ErrNotSupported) {
		t.Errorf("expected ErrNotSupported, got %v", err)
	}
}

func TestDerivePrivateKey(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/ethereum/wallet/priv", 200, `{"key":"0xprivkey"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	key, err := svc.DerivePrivateKey(context.Background(), chain.Ethereum, "word1 word2", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key != "0xprivkey" {
		t.Errorf("unexpected key: %s", key)
	}
}

func TestDerivePrivateKey_Unsupported(t *testing.T) {
	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: "http://unused"}))

	_, err := svc.DerivePrivateKey(context.Background(), chain.Algorand, "word1 word2", 0)
	if !errors.Is(err, walletgen.ErrNotSupported) {
		t.Errorf("expected ErrNotSupported, got %v", err)
	}
}

func TestGenerateWallet_NotFound(t *testing.T) {
	srv := testutil.NewServer(t)
	srv.Handle("/v3/ethereum/wallet", 404, `{"errorCode":"NOT_FOUND","message":"not found"}`)

	svc := walletgen.NewService(httpclient.New(httpclient.Config{BaseURL: srv.URL}))

	_, err := svc.GenerateWallet(context.Background(), chain.Ethereum)
	if err == nil {
		t.Fatal("expected error")
	}
}
