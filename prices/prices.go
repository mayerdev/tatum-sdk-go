package prices

import (
	"context"
	"fmt"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetRate(ctx context.Context, req RateRequest) (*RateResponse, error)
	GetOHLCV(ctx context.Context, req OHLCVRequest) ([]OHLCVCandle, error)
	GetPriceChange(ctx context.Context, req PriceChangeRequest) (*PriceChangeResponse, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetRate(ctx context.Context, req RateRequest) (*RateResponse, error) {
	var out RateResponse
	return &out, s.c.Get(ctx, "/v4/rates/currency", req.toQuery(), &out)
}

func (s *service) GetOHLCV(ctx context.Context, req OHLCVRequest) ([]OHLCVCandle, error) {
	var out []OHLCVCandle
	return out, s.c.Get(ctx, "/v4/rates/ohlcv", req.toQuery(), &out)
}

func (s *service) GetPriceChange(ctx context.Context, req PriceChangeRequest) (*PriceChangeResponse, error) {
	var out PriceChangeResponse
	return &out, s.c.Get(ctx, "/v4/rates/change", req.toQuery(), &out)
}

func itoa64(n int64) string {
	return fmt.Sprintf("%d", n)
}
