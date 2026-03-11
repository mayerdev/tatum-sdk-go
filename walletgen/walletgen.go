package walletgen

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

var ErrNotSupported = errors.New("walletgen: operation not supported for this chain")

type Service interface {
	GenerateWallet(ctx context.Context, c chain.Chain) (*WalletResponse, error)
	DeriveAddress(ctx context.Context, c chain.Chain, xpub string, index uint32) (string, error)
	DerivePrivateKey(ctx context.Context, c chain.Chain, mnemonic string, index uint32) (string, error)
}

type service struct{ c *httpclient.Client }

func NewService(c *httpclient.Client) Service { return &service{c: c} }

func (s *service) GenerateWallet(ctx context.Context, c chain.Chain) (*WalletResponse, error) {
	path, err := chainPath(c)
	if err != nil {
		return nil, err
	}
	var out WalletResponse
	return &out, s.c.Get(ctx, fmt.Sprintf("/v3/%s/wallet", path), nil, &out)
}

func (s *service) DeriveAddress(ctx context.Context, c chain.Chain, xpub string, index uint32) (string, error) {
	if !supportsAddressDerivation(c) {
		return "", ErrNotSupported
	}
	path, _ := chainPath(c)
	var out addressResponse
	err := s.c.Get(ctx, fmt.Sprintf("/v3/%s/address/%s/%d", path, xpub, index), nil, &out)
	return out.Address, err
}

func (s *service) DerivePrivateKey(ctx context.Context, c chain.Chain, mnemonic string, index uint32) (string, error) {
	if !supportsPrivKeyDerivation(c) {
		return "", ErrNotSupported
	}
	path, _ := chainPath(c)
	body := privKeyRequest{Mnemonic: mnemonic, Index: index}
	var out privKeyResponse
	err := s.c.Post(ctx, fmt.Sprintf("/v3/%s/wallet/priv", path), nil, body, &out)
	return out.Key, err
}
