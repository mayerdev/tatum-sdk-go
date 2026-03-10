package fees

import (
	"context"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	EstimateGas(ctx context.Context, req EstimateGasRequest) (*GasEstimate, error)
	GetBlockchainFee(ctx context.Context, req BlockchainFeeRequest) (*BlockchainFee, error)
	EstimateFee(ctx context.Context, req EstimateFeeRequest) (*FeeEstimate, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) EstimateGas(ctx context.Context, req EstimateGasRequest) (*GasEstimate, error) {
	var out GasEstimate
	return &out, s.c.Get(ctx, "/v4/fees/gas", req.toQuery(), &out)
}

func (s *service) GetBlockchainFee(ctx context.Context, req BlockchainFeeRequest) (*BlockchainFee, error) {
	var out BlockchainFee
	return &out, s.c.Get(ctx, "/v4/fees/blockchain", req.toQuery(), &out)
}

func (s *service) EstimateFee(ctx context.Context, req EstimateFeeRequest) (*FeeEstimate, error) {
	var out FeeEstimate
	return &out, s.c.Get(ctx, "/v4/fees/estimate", req.toQuery(), &out)
}
