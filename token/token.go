package token

import (
	"context"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetTransfers(ctx context.Context, req TransfersRequest) ([]TokenTransfer, error)
	GetTrending(ctx context.Context, req TrendingRequest) ([]TokenInfo, error)
	GetNewest(ctx context.Context, req TrendingRequest) ([]TokenInfo, error)
	GetPopular(ctx context.Context, req TrendingRequest) ([]TokenInfo, error)
	GetExchangeRates(ctx context.Context, req ExchangeRatesRequest) (*ExchangeRate, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetTransfers(ctx context.Context, req TransfersRequest) ([]TokenTransfer, error) {
	var out []TokenTransfer
	return out, s.c.Get(ctx, "/v4/data/token/transfers", req.toQuery(), &out)
}

func (s *service) GetTrending(ctx context.Context, req TrendingRequest) ([]TokenInfo, error) {
	var out []TokenInfo
	return out, s.c.Get(ctx, "/v4/data/token/trending", req.toQuery(), &out)
}

func (s *service) GetNewest(ctx context.Context, req TrendingRequest) ([]TokenInfo, error) {
	var out []TokenInfo
	return out, s.c.Get(ctx, "/v4/data/token/newest", req.toQuery(), &out)
}

func (s *service) GetPopular(ctx context.Context, req TrendingRequest) ([]TokenInfo, error) {
	var out []TokenInfo
	return out, s.c.Get(ctx, "/v4/data/token/popular", req.toQuery(), &out)
}

func (s *service) GetExchangeRates(ctx context.Context, req ExchangeRatesRequest) (*ExchangeRate, error) {
	var out ExchangeRate
	return &out, s.c.Get(ctx, "/v4/data/token/rates", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
