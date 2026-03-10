package wallet

import "net/url"

type PortfolioRequest struct {
	Addresses []string
	Chain     string
	PageSize  int
	Cursor    string
}

func (r PortfolioRequest) toQuery() url.Values {
	q := url.Values{}
	for _, a := range r.Addresses {
		q.Add("addresses", a)
	}
	if r.Chain != "" {
		q.Set("chain", r.Chain)
	}
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type PortfolioItem struct {
	Address string  `json:"address"`
	Chain   string  `json:"chain"`
	Balance string  `json:"balance"`
	Asset   string  `json:"asset"`
	Price   float64 `json:"price"`
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
