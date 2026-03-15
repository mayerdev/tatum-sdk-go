package transactions

import (
	"context"
	"errors"
	"strconv"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetHistory(ctx context.Context, req HistoryRequest) ([]Transaction, error)
	GetByHash(ctx context.Context, req ByHashRequest) (*Transaction, error)
	SendSimpleUTXO(ctx context.Context, chainID chain.Chain, req SimpleUTXORequest) (string, error)
	SendNative(ctx context.Context, chainID chain.Chain, req SendNativeRequest) (string, error)
	SendTokens(ctx context.Context, chainID chain.Chain, req SendTokensRequest) (string, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetHistory(ctx context.Context, req HistoryRequest) ([]Transaction, error) {
	var out []Transaction
	return out, s.c.Get(ctx, "/v4/data/transactions", req.toQuery(), &out)
}

func (s *service) GetByHash(ctx context.Context, req ByHashRequest) (*Transaction, error) {
	var out Transaction
	return &out, s.c.Get(ctx, "/v4/data/transactions/hash", req.toQuery(), &out)
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func (s *service) SendSimpleUTXO(ctx context.Context, chainID chain.Chain, req SimpleUTXORequest) (string, error) {
	var out struct {
		TxID string `json:"txId"`
	}
	return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &req, &out)
}

func (s *service) SendNative(ctx context.Context, chainID chain.Chain, req SendNativeRequest) (string, error) {
	var out struct {
		TxID string `json:"txId"`
	}

	if chainID.IsEVM() {
		localReq := map[string]any{
			"fromPrivateKey": req.FromPrivateKey,
			"to":             req.To,
			"currency":       chainID.GetTatumNativeCurrency(),
			"amount":         req.Amount,
		}
		return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &localReq, &out)
	} else if chainID == chain.Ripple {
		localReq := map[string]any{
			"fromAccount": req.From,
			"fromSecret":  req.FromPrivateKey,
			"to":          req.To,
			"amount":      req.Amount,
		}
		return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &localReq, &out)
	}

	localReq := map[string]any{
		"fromPrivateKey": req.FromPrivateKey,
		"to":             req.To,
		"amount":         req.Amount,
	}
	return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &localReq, &out)
}

// SendTokens unsupported: Solana, Avalanche, Fantom
func (s *service) SendTokens(ctx context.Context, chainID chain.Chain, req SendTokensRequest) (string, error) {
	var out struct {
		TxID string `json:"txId"`
	}

	if chainID.IsEVM() {
		localReq := map[string]any{
			"fromPrivateKey": req.FromPrivateKey,
			"to":             req.To,
			"currency":       chainID.RemapCurrency(req.Currency),
			"amount":         req.Amount,
		}
		return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &localReq, &out)
	} else if chainID == chain.Tron {
		localReq := map[string]any{
			"fromPrivateKey": req.FromPrivateKey,
			"to":             req.To,
			"tokenAddress":   req.TokenAddress,
			"amount":         req.Amount,
			"feeLimit":       req.FeeLimit,
		}
		return out.TxID, s.c.Post(ctx, "/v3/"+string(chainID)+"/transaction", nil, &localReq, &out)
	}

	// ethereum: ETH, USDT, LEO, LINK, UNI, MATIC_ETH, POL_ETH, BUSD, SAND, REVV, LATOKEN, COIIN, FREE, XCON, MKR, USDC, BAT, GMC, TUSD, PAX, PLTC, MMY, PAXG, WBTC
	// bsc: BETH, BBTC, RMD, USDC_BSC, USDT_BSC, B2U_BSC, BADA, WBNB, GMC_BSC, BDOT, BXRP, BLTC, BBCH, HAG, CAKE, BUSD_BSC
	// arb: USDC_ARB, USDT_ARB
	// base: USDC_BASE, USDT_BASE
	// polygon: USDC_MATIC, USDT_MATIC, GAMEE, INTENT, EURTENT, GOLDAX
	// optimism: USDC_OP, USDT_OP

	return "", errors.New("unsupported chain")
}
