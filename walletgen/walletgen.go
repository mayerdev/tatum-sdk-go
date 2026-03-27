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
	var out WalletResponse
	return &out, s.c.Get(ctx, fmt.Sprintf("/v3/%s/wallet", chainPath(c)), nil, &out)
}

func (s *service) DeriveAddress(ctx context.Context, ch chain.Chain, xpub string, index uint32) (string, error) {
	caps := ch.GetCapabilities()
	if caps == nil {
		return "", ErrNotSupported
	}

	var apiPath string
	if tpl, ok := customAddressPath[ch]; ok {
		apiPath = fmt.Sprintf(tpl, xpub, index)
	} else {
		apiPath = fmt.Sprintf("/v3/%s/address/%s/%d", chainPath(ch), xpub, index)
	}

	var out addressResponse
	err := s.c.Get(ctx, apiPath, nil, &out)
	return out.Address, err
}

func (s *service) DerivePrivateKey(ctx context.Context, ch chain.Chain, mnemonic string, index uint32) (string, error) {
	caps := ch.GetCapabilities()
	if caps == nil {
		return "", ErrNotSupported
	}

	body := privKeyRequest{Mnemonic: mnemonic, Index: index}
	var out privKeyResponse
	err := s.c.Post(ctx, fmt.Sprintf("/v3/%s/wallet/priv", chainPath(ch)), nil, body, &out)
	return out.value(), err
}
