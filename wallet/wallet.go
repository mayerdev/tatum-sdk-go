package wallet

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetPortfolio(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error)
	GetBalanceByTime(ctx context.Context, req BalanceByTimeRequest) (*BalanceByTimeResponse, error)
	GetReputation(ctx context.Context, req ReputationRequest) (*ReputationResponse, error)
	GetUTXOs(ctx context.Context, req UTXORequest) ([]UTXO, error)
	GetUTXOsBatch(ctx context.Context, req UTXOBatchRequest) ([]UTXO, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetPortfolio(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error) {
	var out []PortfolioItem
	return out, s.c.Get(ctx, "/v4/data/portfolio", req.toQuery(), &out)
}

func (s *service) GetBalanceByTime(ctx context.Context, req BalanceByTimeRequest) (*BalanceByTimeResponse, error) {
	var out BalanceByTimeResponse
	return &out, s.c.Get(ctx, "/v4/data/balance/historical", req.toQuery(), &out)
}

func (s *service) GetReputation(ctx context.Context, req ReputationRequest) (*ReputationResponse, error) {
	var out ReputationResponse
	return &out, s.c.Get(ctx, "/v4/data/reputation", req.toQuery(), &out)
}

func (s *service) GetUTXOs(ctx context.Context, req UTXORequest) ([]UTXO, error) {
	var out []UTXO
	return out, s.c.Get(ctx, "/v4/data/utxos", req.toQuery(), &out)
}

func (s *service) GetUTXOsBatch(ctx context.Context, req UTXOBatchRequest) ([]UTXO, error) {
	var out []UTXO
	return out, s.c.Get(ctx, "/v4/data/utxos/batch", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func itoa64(n int64) string {
	return fmt.Sprintf("%d", n)
}
