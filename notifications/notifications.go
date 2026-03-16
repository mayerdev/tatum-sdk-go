package notifications

import (
	"context"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/pager"
)

type Service interface {
	CreateV3(ctx context.Context, req CreateRequest) (*Notification, error)
	SubscribeAddressTransactions(ctx context.Context, attr CreateAddressTransactionsSubscription) (*Notification, error)

	GetActiveV3(ctx context.Context, filter SubscriptionsFilter) ([]SubscriptionInfo, error)
	GetAllActiveV3(ctx context.Context, address *string) ([]SubscriptionInfo, error)

	CancelV3(ctx context.Context, id string) error

	List(ctx context.Context, req ListRequest) ([]Notification, error)
	Cancel(ctx context.Context, req CancelRequest) error
	GetWebhookLogs(ctx context.Context, req WebhookLogsRequest) ([]WebhookLog, error)
	EnableHMAC(ctx context.Context, req EnableHMACRequest) error
	DisableHMAC(ctx context.Context) error
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) CreateV3(ctx context.Context, req CreateRequest) (*Notification, error) {
	var out Notification
	return &out, s.c.Post(ctx, "/v3/subscription", nil, req, &out)
}

func (s *service) SubscribeAddressTransactions(ctx context.Context, attr CreateAddressTransactionsSubscription) (*Notification, error) {
	return s.CreateV3(ctx, CreateRequest{
		Type: "ADDRESS_TRANSACTION",
		Attr: attr,
	})
}

func (s *service) GetActiveV3(ctx context.Context, filter SubscriptionsFilter) ([]SubscriptionInfo, error) {
	out := make([]SubscriptionInfo, 0)
	return out, s.c.Get(ctx, "/v3/subscription", filter.toQuery(), &out)
}

func (s *service) GetAllActiveV3(ctx context.Context, address *string) ([]SubscriptionInfo, error) {
	filter := SubscriptionsFilter{
		Address: address,
	}

	return pager.FetchAll(&filter.Paginated, func() ([]SubscriptionInfo, error) {
		return s.GetActiveV3(ctx, filter)
	})
}

func (s *service) CancelV3(ctx context.Context, id string) error {
	return s.c.Delete(ctx, "/v3/subscription/"+id, nil, nil)
}

func (s *service) List(ctx context.Context, req ListRequest) ([]Notification, error) {
	var out []Notification
	return out, s.c.Get(ctx, "/v4/notifications/subscriptions", req.toQuery(), &out)
}

func (s *service) Cancel(ctx context.Context, req CancelRequest) error {
	return s.c.Delete(ctx, "/v4/notifications/subscriptions/"+req.ID, nil, nil)
}

func (s *service) GetWebhookLogs(ctx context.Context, req WebhookLogsRequest) ([]WebhookLog, error) {
	var out []WebhookLog
	return out, s.c.Get(ctx, "/v4/notifications/webhooks", req.toQuery(), &out)
}

func (s *service) EnableHMAC(ctx context.Context, req EnableHMACRequest) error {
	return s.c.Put(ctx, "/v4/subscription", nil, req, nil)
}

func (s *service) DisableHMAC(ctx context.Context) error {
	return s.c.Delete(ctx, "/v4/subscription", nil, nil)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
