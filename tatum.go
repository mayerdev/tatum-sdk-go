package tatum

import (
	"gitlab.com/mayerdev/tatum-sdk-go/blockchain"
	"gitlab.com/mayerdev/tatum-sdk-go/defi"
	"gitlab.com/mayerdev/tatum-sdk-go/fees"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
	"gitlab.com/mayerdev/tatum-sdk-go/nameservice"
	"gitlab.com/mayerdev/tatum-sdk-go/nft"
	"gitlab.com/mayerdev/tatum-sdk-go/notifications"
	"gitlab.com/mayerdev/tatum-sdk-go/prices"
	"gitlab.com/mayerdev/tatum-sdk-go/rpc"
	"gitlab.com/mayerdev/tatum-sdk-go/security"
	"gitlab.com/mayerdev/tatum-sdk-go/staking"
	"gitlab.com/mayerdev/tatum-sdk-go/token"
	"gitlab.com/mayerdev/tatum-sdk-go/transactions"
	"gitlab.com/mayerdev/tatum-sdk-go/wallet"
	"gitlab.com/mayerdev/tatum-sdk-go/walletgen"
)

const defaultBaseURL = "https://api.tatum.io"

type Client struct {
	Wallet        wallet.Service
	Transactions  transactions.Service
	Blockchain    blockchain.Service
	NFT           nft.Service
	Token         token.Service
	Prices        prices.Service
	Staking       staking.Service
	DeFi          defi.Service
	Fees          fees.Service
	NameService   nameservice.Service
	Security      security.Service
	Notifications notifications.Service
	RPC           rpc.Service
	WalletGen     walletgen.Service
}

func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	cfg := &clientConfig{
		baseURL:    defaultBaseURL,
		maxRetries: 3,
	}
	for _, o := range opts {
		o(cfg)
	}

	hc := httpclient.New(httpclient.Config{
		APIKey:     apiKey,
		BaseURL:    cfg.baseURL,
		HTTPClient: cfg.httpClient,
		MaxRetries: cfg.maxRetries,
		UserAgent:  cfg.userAgent,
	})

	gatewayTpl := cfg.rpcBaseURL
	if gatewayTpl == "" {
		gatewayTpl = "https://%s-%s.gateway.tatum.io/"
	}

	return &Client{
		Wallet:        wallet.NewService(hc),
		Transactions:  transactions.NewService(hc),
		Blockchain:    blockchain.NewService(hc),
		NFT:           nft.NewService(hc),
		Token:         token.NewService(hc),
		Prices:        prices.NewService(hc),
		Staking:       staking.NewService(hc),
		DeFi:          defi.NewService(hc),
		Fees:          fees.NewService(hc),
		NameService:   nameservice.NewService(hc),
		Security:      security.NewService(hc),
		Notifications: notifications.NewService(hc),
		RPC:           rpc.NewService(hc, gatewayTpl),
		WalletGen:     walletgen.NewService(hc),
	}, nil
}
