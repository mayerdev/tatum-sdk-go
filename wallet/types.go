package wallet

import (
	"net/url"

	"github.com/shopspring/decimal"
	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/pager"
)

type PortfolioTokenType string

const (
	TokenNative      PortfolioTokenType = "native"
	TokenFungible    PortfolioTokenType = "fungible"
	TokenMultiAndNFT PortfolioTokenType = "nft,multitoken"
)

type PortfolioRequest struct {
	Chain           chain.ChainNetwork
	Addresses       []string // only 1 item allowed
	TokenTypes      PortfolioTokenType
	ExcludeMetadata bool
	pager.Paginated
}

func (req PortfolioRequest) toQuery() url.Values {
	query := url.Values{}
	query.Set("chain", string(req.Chain))
	query.Set("tokenTypes", string(req.TokenTypes))

	for _, a := range req.Addresses {
		query.Add("addresses", a)
	}

	if req.ExcludeMetadata {
		query.Set("excludeMetadata", "true")
	}

	req.Paginated.AssignTo(&query)
	return query
}

type PortfolioItem struct {
	Address string             `json:"address"`
	Chain   chain.ChainNetwork `json:"chain"`
	Balance decimal.Decimal    `json:"balance"`
	Asset   string             `json:"asset"`
	Price   float64            `json:"price"`
}

type BalanceByTimeRequest struct {
	Address string
	Chain   string
	Time    int64
}

func (r BalanceByTimeRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	if r.Time > 0 {
		q.Set("time", itoa64(r.Time))
	}
	return q
}

type BalanceByTimeResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Time    int64  `json:"time"`
}

type ReputationRequest struct {
	Address string
	Chain   string
}

func (r ReputationRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type ReputationResponse struct {
	Address    string `json:"address"`
	Reputation string `json:"reputation"`
}

type UTXORequest struct {
	Address string
	Chain   string
}

func (r UTXORequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type UTXO struct {
	TxHash  string `json:"txHash"`
	Index   int    `json:"index"`
	Value   string `json:"value"`
	Address string `json:"address"`
}

type UTXOBatchRequest struct {
	Addresses []string
	Chain     string
}

func (r UTXOBatchRequest) toQuery() url.Values {
	q := url.Values{}
	for _, a := range r.Addresses {
		q.Add("addresses", a)
	}
	q.Set("chain", r.Chain)
	return q
}
