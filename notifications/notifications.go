package notifications

import (
	"context"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	Create(ctx context.Context, req CreateRequest) (*Notification, error)
	List(ctx context.Context, req ListRequest) ([]Notification, error)
	Cancel(ctx context.Context, req CancelRequest) error
	GetWebhookLogs(ctx context.Context, req WebhookLogsRequest) ([]WebhookLog, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*Notification, error) {
	var out Notification
	return &out, s.c.Post(ctx, "/v4/notifications/subscribe", nil, req, &out)
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

func itoa(n int) string {
	return strconv.Itoa(n)
}
