package notifications

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
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

	if digest != xPayloadHash {
		return nil, false, nil
	}

	if payload.SubscriptionType == "ADDRESS_TRANSACTION" || payload.SubscriptionType == "ADDRESS_EVENT" {
		parseHookTransfers(&payload)
	}

	return &payload, true, nil
}

func parseHookTransfers(payload *WebhookPayload) {
	chainID, network := payload.Chain.Split()
	if chainID.IsEVM() {
		// https://docs.tatum.io/docs/notifications-understanding-address-and-counteraddress
		if payload.Type == "native" {
			payload.Transfers = append(payload.Transfers, TransferMeta{
				Native:   true,
				Asset:    payload.Asset,
				Sender:   payload.CounterAddress,
				Receiver: payload.Address,
			})
		} else if payload.Type == "token" {
			payload.Transfers = append(payload.Transfers, TransferMeta{
				Native:       false,
				AssetAddress: payload.Asset,
				Sender:       payload.Address,
				Receiver:     payload.CounterAddress,
			})
		}
	} else if chainID == chain.Tron {
		if payload.Type == "native" || payload.Type == "trc10" || payload.Type == "trc20" || payload.Type == "token" {
			native := payload.Type == "native"
			asset := ""
			assetAddress := ""

			// https://docs.tatum.io/docs/about-contractaddress-value-and-symbol
			if !native {
				assetAddress = payload.ContractAddress

				if network == chain.Mainnet {
					if assetAddress == "INRT_TRON" {
						asset = "INRT"
						assetAddress = "TX66VmiV1txm45vVLvcHYEqPXXLoREyAXm"
					} else if assetAddress == "USDT_TRON" {
						asset = "USDT"
						assetAddress = "TX66VmiV1txm45vVLvcHYEqPXXLoREyAXm"
					}
				}
			}

			payload.Transfers = append(payload.Transfers, TransferMeta{
				Native:       native,
				Asset:        asset,
				AssetAddress: assetAddress,
				Sender:       payload.CounterAddress,
				Receiver:     payload.Address,
			})
		}
	} else {
		if payload.Type == "native" {
			payload.Transfers = append(payload.Transfers, TransferMeta{
				Native:   true,
				Asset:    payload.Asset,
				Receiver: payload.Address,
			})
		} else if payload.Type == "token" {
			payload.Transfers = append(payload.Transfers, TransferMeta{
				Native:       false,
				AssetAddress: payload.Asset,
				Receiver:     payload.Address,
			})
		}
	}
}
