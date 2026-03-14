package notifications

import (
	"net/url"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
)

type EnableHMACRequest struct {
	HMACSecret string `json:"hmacSecret"`
}

type TransferMeta struct {
	Native       bool
	Asset        string
	AssetAddress string
	Sender       string // can be empty
	Receiver     string
}

type WebhookPayload struct {
	Address          string             `json:"address"`
	Amount           string             `json:"amount"`
	CounterAddress   string             `json:"counterAddress"`
	Asset            string             `json:"asset"`
	BlockNumber      int64              `json:"blockNumber"`
	TxID             string             `json:"txId"`
	Type             string             `json:"type"`
	TokenID          *string            `json:"tokenId"`
	ContractAddress  string             `json:"contractAddress"`
	Chain            chain.ChainNetwork `json:"chain"`
	SubscriptionType string             `json:"subscriptionType"`

	Transfers []TransferMeta
}

type CreateRequest struct {
	Type string `json:"type"`
	Attr any    `json:"attr"`
}

type CreateAddressTransactionsSubscription struct {
	Address      string             `json:"address"`
	ChainNetwork chain.ChainNetwork `json:"chain"`
	URL          string             `json:"url"`
}

type Notification struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Chain   string `json:"chain"`
	URL     string `json:"url"`
	Active  bool   `json:"active"`
}

type ListRequest struct {
	PageSize int
	Cursor   string
}

func (r ListRequest) toQuery() url.Values {
	q := url.Values{}
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}

type CancelRequest struct {
	ID string
}

type WebhookLog struct {
	ID             string `json:"id"`
	NotificationID string `json:"notificationId"`
	StatusCode     int    `json:"statusCode"`
	Timestamp      int64  `json:"timestamp"`
	Failed         bool   `json:"failed"`
}

type WebhookLogsRequest struct {
	NotificationID string
	PageSize       int
	Cursor         string
}

func (r WebhookLogsRequest) toQuery() url.Values {
	q := url.Values{}
	if r.NotificationID != "" {
		q.Set("notificationId", r.NotificationID)
	}
	if r.PageSize > 0 {
		q.Set("pageSize", itoa(r.PageSize))
	}
	if r.Cursor != "" {
		q.Set("cursor", r.Cursor)
	}
	return q
}
