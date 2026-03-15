package fees

import (
	"net/url"
	"time"
)

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

type Priority uint8

const (
	PriorityFast Priority = iota
	PriorityMedium
	PrioritySlow
)

type BlockchainFee struct {
	// BTC-like
	Fast   uint64    `json:"fast"`
	Medium uint64    `json:"medium"`
	Slow   uint64    `json:"slow"`
	Block  int       `json:"block"`
	Time   time.Time `json:"time"`
	Weight int       `json:"weight"`

	// EVM
	GasLimit string `json:"gasLimit"`
	GasPrice string `json:"gasPrice"` // wei (10^-18 ETH)
}

func (fee BlockchainFee) Get(kind Priority) uint64 {
	switch kind {
	case PriorityFast:
		return fee.Fast
	case PriorityMedium:
		return fee.Medium
	case PrioritySlow:
		return fee.Slow
	default:
		return 0
	}
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
