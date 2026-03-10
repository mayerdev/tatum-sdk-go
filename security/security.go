package security

import (
	"context"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	CheckAddress(ctx context.Context, req CheckAddressRequest) (*CheckAddressResponse, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) CheckAddress(ctx context.Context, req CheckAddressRequest) (*CheckAddressResponse, error) {
	var out CheckAddressResponse
	return &out, s.c.Get(ctx, "/v4/security/address", req.toQuery(), &out)
}
