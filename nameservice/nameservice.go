package nameservice

import (
	"context"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	Resolve(ctx context.Context, req ResolveRequest) (*ResolveResponse, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) Resolve(ctx context.Context, req ResolveRequest) (*ResolveResponse, error) {
	var out ResolveResponse
	return &out, s.c.Get(ctx, "/v4/tns/resolve", req.toQuery(), &out)
}
