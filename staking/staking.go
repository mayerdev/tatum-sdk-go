package staking

import (
	"context"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetStakedAssets(ctx context.Context, req StakedAssetsRequest) ([]StakedAsset, error)
	GetValidators(ctx context.Context, req ValidatorsRequest) ([]Validator, error)
	GetRewards(ctx context.Context, req RewardsRequest) (*RewardResponse, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetStakedAssets(ctx context.Context, req StakedAssetsRequest) ([]StakedAsset, error) {
	var out []StakedAsset
	return out, s.c.Get(ctx, "/v4/staking/positions", req.toQuery(), &out)
}

func (s *service) GetValidators(ctx context.Context, req ValidatorsRequest) ([]Validator, error) {
	var out []Validator
	return out, s.c.Get(ctx, "/v4/staking/validators", req.toQuery(), &out)
}

func (s *service) GetRewards(ctx context.Context, req RewardsRequest) (*RewardResponse, error) {
	var out RewardResponse
	return &out, s.c.Get(ctx, "/v4/staking/rewards", req.toQuery(), &out)
}
