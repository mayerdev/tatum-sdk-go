package notifications

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func ParseAndVerifyWebhook(body []byte, xPayloadHash string, hmacSecret string) (*WebhookPayload, bool, error) {
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, false, fmt.Errorf("parse webhook body: %w", err)
	}

	if hmacSecret == "" {
		return &payload, true, nil
	}

	mac := hmac.New(sha512.New, []byte(hmacSecret))
	mac.Write(body)
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return &payload, digest == xPayloadHash, nil
}
