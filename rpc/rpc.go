package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	Call(ctx context.Context, c chain.Chain, n chain.Network, method string, params any) (json.RawMessage, error)
}

type service struct {
	c          *httpclient.Client
	gatewayTpl string
}

func NewService(c *httpclient.Client, gatewayTpl string) Service {
	if gatewayTpl == "" {
		gatewayTpl = "https://%s-%s.gateway.tatum.io/"
	}
	return &service{c: c, gatewayTpl: gatewayTpl}
}

func (s *service) Call(ctx context.Context, c chain.Chain, n chain.Network, method string, params any) (json.RawMessage, error) {
	gatewayURL := fmt.Sprintf(s.gatewayTpl, c, n)
	body := jsonRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
	var resp jsonRPCResponse
	if err := s.c.PostAbsolute(ctx, gatewayURL, body, &resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", resp.Error.Code, resp.Error.Message)
	}
	return resp.Result, nil
}
