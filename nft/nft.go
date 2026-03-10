package nft

import (
	"context"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetCollectionTokens(ctx context.Context, req CollectionTokensRequest) ([]NFTToken, error)
	GetMetadata(ctx context.Context, req MetadataRequest) (*NFTToken, error)
	GetOwners(ctx context.Context, req OwnersRequest) (*OwnersResponse, error)
	GetWalletBalances(ctx context.Context, req WalletBalancesRequest) ([]NFTToken, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetCollectionTokens(ctx context.Context, req CollectionTokensRequest) ([]NFTToken, error) {
	var out []NFTToken
	return out, s.c.Get(ctx, "/v4/data/collections", req.toQuery(), &out)
}

func (s *service) GetMetadata(ctx context.Context, req MetadataRequest) (*NFTToken, error) {
	var out NFTToken
	return &out, s.c.Get(ctx, "/v4/data/metadata", req.toQuery(), &out)
}

func (s *service) GetOwners(ctx context.Context, req OwnersRequest) (*OwnersResponse, error) {
	var out OwnersResponse
	return &out, s.c.Get(ctx, "/v4/data/owners", req.toQuery(), &out)
}

func (s *service) GetWalletBalances(ctx context.Context, req WalletBalancesRequest) ([]NFTToken, error) {
	var out []NFTToken
	return out, s.c.Get(ctx, "/v4/data/nfts", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
