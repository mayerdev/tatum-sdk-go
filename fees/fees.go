package fees

import (
	"context"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	EstimateGas(ctx context.Context, req EstimateGasRequest) (*GasEstimate, error)
	GetBlockchainFee(ctx context.Context, chainID chain.Chain) (*BlockchainFee, error)
	EstimateFee(ctx context.Context, req EstimateFeeRequest) (*FeeEstimate, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) EstimateGas(ctx context.Context, req EstimateGasRequest) (*GasEstimate, error) {
	var out GasEstimate
	return &out, s.c.Get(ctx, "/v4/fees/gas", req.toQuery(), &out)
}

func (s *service) GetBlockchainFee(ctx context.Context, chainID chain.Chain) (*BlockchainFee, error) {
	var chainCode string
	var dummyAddress string
	useGasEndpoint := false

	switch chainID {
	case chain.Ethereum:
		chainCode = "ETH"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.BNBSmartChain:
		chainCode = "BSC"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Celo:
		chainCode = "CELO"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.EGLD:
		chainCode = "EGLD"
		dummyAddress = "erd1l9wp4d4q2p5m6emr9l786f4calsm02emuxehmq34pscez2gvfxlqxr5j9s"
		useGasEndpoint = true
	case chain.Harmony:
		chainCode = "ONE"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Klaytn:
		chainCode = "KLAY"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.KuCoin:
		chainCode = "KCS"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Flare:
		chainCode = "FLR"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Cronos:
		chainCode = "CRO"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Avalanche:
		chainCode = "AVAX"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Base:
		chainCode = "ETH_BASE"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Polygon:
		chainCode = "POL_ETH"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Optimism:
		chainCode = "ETH_OP"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Fantom:
		chainCode = "FTM"
		dummyAddress = "0x0000000000000000000000000000000000000000"
		useGasEndpoint = true
	case chain.Bitcoin:
		chainCode = "BTC"
	case chain.Litecoin:
		chainCode = "LTC"
	case chain.Dogecoin:
		chainCode = "DOGE"
	}

	var out BlockchainFee
	if useGasEndpoint {
		req := map[string]string{
			"chain":  chainCode,
			"from":   dummyAddress,
			"to":     dummyAddress,
			"amount": "0",
		}

		return &out, s.c.Post(ctx, "/v4/blockchainOperations/gas", nil, &req, &out)
	}

	return &out, s.c.Get(ctx, "/v3/blockchain/fee/"+chainCode, nil, &out)
}

func (s *service) EstimateFee(ctx context.Context, req EstimateFeeRequest) (*FeeEstimate, error) {
	var out FeeEstimate
	return &out, s.c.Get(ctx, "/v4/fees/estimate", req.toQuery(), &out)
}
