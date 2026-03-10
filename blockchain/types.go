package blockchain

import "net/url"

type CurrentBlockRequest struct {
	Chain string
}

func (r CurrentBlockRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type CurrentBlockResponse struct {
	Chain  string `json:"chain"`
	Number int64  `json:"number"`
	Hash   string `json:"hash"`
}

type BlockByHashRequest struct {
	Hash  string
	Chain string
}

func (r BlockByHashRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("hash", r.Hash)
	q.Set("chain", r.Chain)
	return q
}

type BlockByHeightRequest struct {
	Height int64
	Chain  string
}

func (r BlockByHeightRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("height", itoa64(r.Height))
	q.Set("chain", r.Chain)
	return q
}

type Block struct {
	Hash         string   `json:"hash"`
	Number       int64    `json:"number"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

type TxByHashRequest struct {
	Hash  string
	Chain string
}

func (r TxByHashRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("hash", r.Hash)
	q.Set("chain", r.Chain)
	return q
}

type BlockchainTx struct {
	Hash        string `json:"hash"`
	BlockNumber int64  `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
}

type NativeBalanceRequest struct {
	Address string
	Chain   string
}

func (r NativeBalanceRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type NativeBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type BatchBalanceRequest struct {
	Addresses []string
	Chain     string
}

func (r BatchBalanceRequest) toQuery() url.Values {
	q := url.Values{}
	for _, a := range r.Addresses {
		q.Add("addresses", a)
	}
	q.Set("chain", r.Chain)
	return q
}

type MempoolRequest struct {
	Chain string
}

func (r MempoolRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type MempoolResponse struct {
	Transactions []string `json:"transactions"`
}
