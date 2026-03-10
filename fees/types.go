package fees

import "net/url"

type EstimateGasRequest struct {
	Chain string
	From  string
	To    string
	Value string
	Data  string
}

func (r EstimateGasRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	q.Set("from", r.From)
	q.Set("to", r.To)
	if r.Value != "" {
		q.Set("value", r.Value)
	}
	if r.Data != "" {
		q.Set("data", r.Data)
	}
	return q
}

type GasEstimate struct {
	GasLimit string `json:"gasLimit"`
	GasPrice string `json:"gasPrice"`
}

type BlockchainFeeRequest struct {
	Chain string
}

func (r BlockchainFeeRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type BlockchainFee struct {
	Chain  string `json:"chain"`
	Fast   string `json:"fast"`
	Medium string `json:"medium"`
	Slow   string `json:"slow"`
}

type EstimateFeeRequest struct {
	Chain string
	Type  string
}

func (r EstimateFeeRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	if r.Type != "" {
		q.Set("type", r.Type)
	}
	return q
}

type FeeEstimate struct {
	Fee      string `json:"fee"`
	Currency string `json:"currency"`
}
