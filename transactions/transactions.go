package transactions

import (
	"context"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetHistory(ctx context.Context, req HistoryRequest) ([]Transaction, error)
	GetByHash(ctx context.Context, req ByHashRequest) (*Transaction, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetHistory(ctx context.Context, req HistoryRequest) ([]Transaction, error) {
	var out []Transaction
	return out, s.c.Get(ctx, "/v4/data/transactions", req.toQuery(), &out)
}

func (s *service) GetByHash(ctx context.Context, req ByHashRequest) (*Transaction, error) {
	var out Transaction
	return &out, s.c.Get(ctx, "/v4/data/transactions/hash", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
