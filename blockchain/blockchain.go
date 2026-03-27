package blockchain

import (
	"context"
	"fmt"
	"net/url"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetCurrentBlock(ctx context.Context, cn chain.ChainNetwork) (int64, error)
	GetBlockByHash(ctx context.Context, req BlockByHashRequest) (*Block, error)
	GetBlockByHeight(ctx context.Context, req BlockByHeightRequest) (*Block, error)
	GetTxByHash(ctx context.Context, req TxByHashRequest) (*TransactionInfo, error)
	GetNativeBalance(ctx context.Context, req NativeBalanceRequest) (*NativeBalanceResponse, error)
	GetBatchBalance(ctx context.Context, req BatchBalanceRequest) ([]NativeBalanceResponse, error)
	GetMempool(ctx context.Context, req MempoolRequest) (*MempoolResponse, error)
}

type service struct {
	c *httpclient.Client
}

func NewService(c *httpclient.Client) Service {
	return &service{c: c}
}

func (s *service) GetCurrentBlock(ctx context.Context, cn chain.ChainNetwork) (int64, error) {
	query := url.Values{}
	query.Set("chain", string(cn))

	var result int64
	return result, s.c.Get(ctx, "/v4/data/blockchains/block/current", query, &result)
}

func (s *service) GetBlockByHash(ctx context.Context, req BlockByHashRequest) (*Block, error) {
	var out Block
	return &out, s.c.Get(ctx, "/v4/blockchain/block/hash", req.toQuery(), &out)
}

func (s *service) GetBlockByHeight(ctx context.Context, req BlockByHeightRequest) (*Block, error) {
	var out Block
	return &out, s.c.Get(ctx, "/v4/blockchain/block/height", req.toQuery(), &out)
}

func (s *service) GetTxByHash(ctx context.Context, req TxByHashRequest) (*TransactionInfo, error) {
	var out TransactionInfo
	var err error

	ch, _ := req.Chain.Split()
	if ch.IsUTXO() {
		out.UTXO = &UTXOTransactionInfo{}
		err = s.c.Get(ctx, "/v4/data/blockchains/transaction", req.toQuery(), out.UTXO)
	} else if ch.IsEVM() {
		out.EVM = &EVMTransactionInfo{}
		err = s.c.Get(ctx, "/v4/data/blockchains/transaction", req.toQuery(), out.EVM)
	} else if ch == chain.Solana {
		out.Solana = &SolanaTransactionInfo{}
		err = s.c.Get(ctx, "/v4/data/blockchains/transaction", req.toQuery(), out.Solana)
	} else if ch == chain.Tron {
		out.Tron = &TronTransactionInfo{}
		err = s.c.Get(ctx, "/v4/data/blockchains/transaction", req.toQuery(), out.Tron)
	} else {
		return nil, fmt.Errorf("unsupported chain: %s", ch)
	}

	if err != nil {
		return nil, err
	}

	out.Normalize()
	return &out, err
}

func (s *service) GetNativeBalance(ctx context.Context, req NativeBalanceRequest) (*NativeBalanceResponse, error) {
	var out NativeBalanceResponse
	return &out, s.c.Get(ctx, "/v4/blockchain/balance", req.toQuery(), &out)
}

func (s *service) GetBatchBalance(ctx context.Context, req BatchBalanceRequest) ([]NativeBalanceResponse, error) {
	var out []NativeBalanceResponse
	return out, s.c.Get(ctx, "/v4/blockchain/balance/batch", req.toQuery(), &out)
}

func (s *service) GetMempool(ctx context.Context, req MempoolRequest) (*MempoolResponse, error) {
	var out MempoolResponse
	return &out, s.c.Get(ctx, "/v4/blockchain/mempool", req.toQuery(), &out)
}

func itoa64(n int64) string {
	return fmt.Sprintf("%d", n)
}
