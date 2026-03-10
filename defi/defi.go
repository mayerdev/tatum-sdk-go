package defi

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetEvents(ctx context.Context, req EventsRequest) ([]DefiEvent, error)
	GetBlocks(ctx context.Context, req BlocksRequest) ([]DefiBlock, error)
	GetLatestBlock(ctx context.Context, req LatestBlockRequest) (*DefiBlock, error)
	GetBlockByTimestamp(ctx context.Context, req BlockByTimestampRequest) (*DefiBlock, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetEvents(ctx context.Context, req EventsRequest) ([]DefiEvent, error) {
	var out []DefiEvent
	return out, s.c.Get(ctx, "/v4/defi/events", req.toQuery(), &out)
}

func (s *service) GetBlocks(ctx context.Context, req BlocksRequest) ([]DefiBlock, error) {
	var out []DefiBlock
	return out, s.c.Get(ctx, "/v4/defi/blocks", req.toQuery(), &out)
}

func (s *service) GetLatestBlock(ctx context.Context, req LatestBlockRequest) (*DefiBlock, error) {
	var out DefiBlock
	return &out, s.c.Get(ctx, "/v4/defi/blocks/latest", req.toQuery(), &out)
}

func (s *service) GetBlockByTimestamp(ctx context.Context, req BlockByTimestampRequest) (*DefiBlock, error) {
	var out DefiBlock
	return &out, s.c.Get(ctx, "/v4/defi/blocks/timestamp", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func itoa64(n int64) string {
	return fmt.Sprintf("%d", n)
}
