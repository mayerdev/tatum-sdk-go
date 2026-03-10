package token

import "net/url"

type TransfersRequest struct {
	Address  string
	Contract string
	Chain    string
	PageSize int
	Cursor   string
}

func (r TransfersRequest) toQuery() url.Values {
	q := url.Values{}
	if r.Address != "" {
		q.Set("address", r.Address)
	}
	if r.Contract != "" {
		q.Set("contract", r.Contract)
	}
	q.Set("chain", r.Chain)
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type TokenTransfer struct {
	TxHash   string `json:"txHash"`
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Contract string `json:"contract"`
	Chain    string `json:"chain"`
}

type TrendingRequest struct {
	Chain    string
	PageSize int
}

func (r TrendingRequest) toQuery() url.Values {
	q := url.Values{}
	if r.Chain != "" {
		q.Set("chain", r.Chain)
	}
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	return q
}

type TokenInfo struct {
	Contract string  `json:"contract"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Chain    string  `json:"chain"`
	Price    float64 `json:"price"`
}

type ExchangeRatesRequest struct {
	Contract string
	Chain    string
}

func (r ExchangeRatesRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("contract", r.Contract)
	q.Set("chain", r.Chain)
	return q
}

type ExchangeRate struct {
	Contract string  `json:"contract"`
	Rate     float64 `json:"rate"`
	Currency string  `json:"currency"`
}
