package blockchain

import (
	"context"
	"fmt"

	"gitlab.com/mayerdev/tatum-sdk-go/internal/httpclient"
)

type Service interface {
	GetCurrentBlock(ctx context.Context, req CurrentBlockRequest) (*CurrentBlockResponse, error)
	GetBlockByHash(ctx context.Context, req BlockByHashRequest) (*Block, error)
	GetBlockByHeight(ctx context.Context, req BlockByHeightRequest) (*Block, error)
	GetTxByHash(ctx context.Context, req TxByHashRequest) (*BlockchainTx, error)
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

func (s *service) GetCurrentBlock(ctx context.Context, req CurrentBlockRequest) (*CurrentBlockResponse, error) {
	var out CurrentBlockResponse
	return &out, s.c.Get(ctx, "/v4/blockchain/block/current", req.toQuery(), &out)
}

func (s *service) GetBlockByHash(ctx context.Context, req BlockByHashRequest) (*Block, error) {
	var out Block
	return &out, s.c.Get(ctx, "/v4/blockchain/block/hash", req.toQuery(), &out)
}

func (s *service) GetBlockByHeight(ctx context.Context, req BlockByHeightRequest) (*Block, error) {
	var out Block
	return &out, s.c.Get(ctx, "/v4/blockchain/block/height", req.toQuery(), &out)
}

func (s *service) GetTxByHash(ctx context.Context, req TxByHashRequest) (*BlockchainTx, error) {
	var out BlockchainTx
	return &out, s.c.Get(ctx, "/v4/blockchain/transaction/hash", req.toQuery(), &out)
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
