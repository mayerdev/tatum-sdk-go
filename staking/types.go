package staking

import "net/url"

type StakedAssetsRequest struct {
	Address string
	Chain   string
}

func (r StakedAssetsRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type StakedAsset struct {
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	Amount    string `json:"amount"`
	Validator string `json:"validator"`
	Reward    string `json:"reward"`
}

type ValidatorsRequest struct {
	Chain string
}

func (r ValidatorsRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type Validator struct {
	Address    string  `json:"address"`
	Commission float64 `json:"commission"`
	Stake      string  `json:"stake"`
	Status     string  `json:"status"`
}

type RewardsRequest struct {
	Address string
	Chain   string
}

func (r RewardsRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type RewardResponse struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
	Reward  string `json:"reward"`
}
