package wallet

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/pager"
)

type Service interface {
	GetPortfolio(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error)
	GetFullPortfolio(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error)
	GetPortfolioBatch(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error)

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
	chainID, _ := req.Chain.Split()
	if chainID == chain.Bitcoin {
		if req.TokenTypes == TokenNative {
			var out struct {
				Incoming decimal.Decimal `json:"incoming"`
				Outgoing decimal.Decimal `json:"outgoing"`
			}

			err := s.c.Get(ctx, "/v3/bitcoin/address/balance/"+req.Addresses[0], nil, &out)
			if err != nil {
				return nil, err
			}

			return []PortfolioItem{
				{
					Address: req.Addresses[0],
					Chain:   req.Chain,
					Balance: out.Incoming.Sub(out.Outgoing),
					Asset:   "",
					Price:   0,
				},
			}, nil
		}

		return nil, nil
	}

	var out struct {
		Result []PortfolioItem `json:"result"`
	}

	return out.Result, s.c.Get(ctx, "/v4/data/wallet/portfolio", req.toQuery(), &out)
}

func (s *service) GetFullPortfolio(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error) {
	return pager.FetchAll(&req.Paginated, func() ([]PortfolioItem, error) {
		return s.GetPortfolio(ctx, req)
	})
}

func (s *service) GetPortfolioBatch(ctx context.Context, req PortfolioRequest) ([]PortfolioItem, error) {
	chainID, _ := req.Chain.Split()
	if chainID == chain.Bitcoin {
		if req.TokenTypes == TokenNative {
			var out []struct {
				Incoming decimal.Decimal `json:"incoming"`
				Outgoing decimal.Decimal `json:"outgoing"`
			}

			query := url.Values{}
			query.Set("addresses", strings.Join(req.Addresses, ","))

			err := s.c.Get(ctx, "/v3/bitcoin/address/balance/batch", query, &out)
			if err != nil {
				return nil, err
			}

			result := make([]PortfolioItem, len(out))
			for i, item := range out {
				result[i].Address = req.Addresses[i]
				result[i].Chain = req.Chain
				result[i].Balance = item.Incoming.Sub(item.Outgoing)
				result[i].Asset = chainID.GetNativeCurrency()
				result[i].Price = 0
			}

			return result, nil
		}

		return nil, nil
	}

	req2 := req
	result := make([]PortfolioItem, 0)
	var part []PortfolioItem
	var err error

	for _, address := range req.Addresses {
		req2.Addresses = []string{address}
		part, err = s.GetFullPortfolio(ctx, req2)
		if err != nil {
			return nil, err
		}

		result = append(result, part...)
	}

	return result, nil
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
