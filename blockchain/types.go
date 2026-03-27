package blockchain

import (
	"net/url"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
)

type CurrentBlockRequest struct {
	Chain string
}

func (r CurrentBlockRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type CurrentBlockResponse struct {
	Chain  string `json:"chain"`
	Number int64  `json:"number"`
	Hash   string `json:"hash"`
}

type BlockByHashRequest struct {
	Hash  string
	Chain string
}

func (r BlockByHashRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("hash", r.Hash)
	q.Set("chain", r.Chain)
	return q
}

type BlockByHeightRequest struct {
	Height int64
	Chain  string
}

func (r BlockByHeightRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("height", itoa64(r.Height))
	q.Set("chain", r.Chain)
	return q
}

type Block struct {
	Hash         string   `json:"hash"`
	Number       int64    `json:"number"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

type TxByHashRequest struct {
	Hash  string
	Chain chain.ChainNetwork
}

func (r TxByHashRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("hash", r.Hash)
	q.Set("chain", string(r.Chain))
	return q
}

type EVMTransactionInfo struct {
	Hash                 string `json:"hash"`
	TransactionHash      string `json:"transactionHash"`
	BlockHash            string `json:"blockHash"`
	BlockNumber          int    `json:"blockNumber"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	Gas                  int    `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         int    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas int    `json:"maxPriorityFeePerGas"`
	Nonce                int    `json:"nonce"`
	TransactionIndex     int    `json:"transactionIndex"`
	Input                string `json:"input"`
	Type                 int    `json:"type"`
	AccessList           []any  `json:"accessList"`
	ChainId              int    `json:"chainId"`
	YParity              string `json:"yParity"`
	GasUsed              int    `json:"gasUsed"`
	EffectiveGasPrice    int    `json:"effectiveGasPrice"`
	CumulativeGasUsed    int    `json:"cumulativeGasUsed"`
	ContractAddress      string `json:"contractAddress"`
	Status               bool   `json:"status"`
	LogsBloom            string `json:"logsBloom"`
	Logs                 []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      int      `json:"blockNumber"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex int      `json:"transactionIndex"`
		BlockHash        string   `json:"blockHash"`
		LogIndex         int      `json:"logIndex"`
		Removed          bool     `json:"removed"`
		BlockTimestamp   string   `json:"blockTimestamp"`
	} `json:"logs"`
	Fee            string `json:"fee"`
	AdditionalProp any    `json:"additionalProp"`
}

type UTXOTransactionInfo struct {
	BlockNumber int    `json:"blockNumber"`
	Fee         int    `json:"fee"`
	Hash        string `json:"hash"`
	Hex         string `json:"hex"`
	Index       int    `json:"index"`
	Inputs      []struct {
		Prevout struct {
			Hash  string `json:"hash"`
			Index int    `json:"index"`
		} `json:"prevout"`
		Sequence int64  `json:"sequence"`
		Script   string `json:"script"`
		Coin     struct {
			Version  int    `json:"version"`
			Height   int    `json:"height"`
			Value    int    `json:"value"`
			Script   string `json:"script"`
			Address  string `json:"address"`
			Type     string `json:"type"`
			ReqSigs  int    `json:"reqSigs"`
			Coinbase bool   `json:"coinbase"`
		} `json:"coin"`
	} `json:"inputs"`
	Locktime int `json:"locktime"`
	Outputs  []struct {
		Value        int    `json:"value"`
		Script       string `json:"script"`
		Address      string `json:"address"`
		ScriptPubKey struct {
			Type    string `json:"type"`
			ReqSigs int    `json:"reqSigs"`
		} `json:"scriptPubKey"`
	} `json:"outputs"`
	Size        int    `json:"size"`
	Time        int    `json:"time"`
	Version     int    `json:"version"`
	Vsize       int    `json:"vsize"`
	Weight      int    `json:"weight"`
	WitnessHash string `json:"witnessHash"`
	Block       string `json:"block"`
}

type SolanaTransactionInfo struct {
	BlockTime int `json:"blockTime"`
	Meta      struct {
		ComputeUnitsConsumed int         `json:"computeUnitsConsumed"`
		CostUnits            int         `json:"costUnits"`
		Err                  interface{} `json:"err"`
		Fee                  int         `json:"fee"`
		InnerInstructions    []struct {
			Index        int `json:"index"`
			Instructions []struct {
				Accounts       []int  `json:"accounts"`
				Data           string `json:"data"`
				ProgramIdIndex int    `json:"programIdIndex"`
				StackHeight    int    `json:"stackHeight"`
			} `json:"instructions"`
		} `json:"innerInstructions"`
		LoadedAddresses struct {
			Readonly []interface{} `json:"readonly"`
			Writable []interface{} `json:"writable"`
		} `json:"loadedAddresses"`
		LogMessages       []string `json:"logMessages"`
		PostBalances      []int64  `json:"postBalances"`
		PostTokenBalances []struct {
			AccountIndex  int    `json:"accountIndex"`
			Mint          string `json:"mint"`
			Owner         string `json:"owner"`
			ProgramId     string `json:"programId"`
			UiTokenAmount struct {
				Amount         string  `json:"amount"`
				Decimals       int     `json:"decimals"`
				UiAmount       float64 `json:"uiAmount"`
				UiAmountString string  `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"postTokenBalances"`
		PreBalances      []int64 `json:"preBalances"`
		PreTokenBalances []struct {
			AccountIndex  int    `json:"accountIndex"`
			Mint          string `json:"mint"`
			Owner         string `json:"owner"`
			ProgramId     string `json:"programId"`
			UiTokenAmount struct {
				Amount         string  `json:"amount"`
				Decimals       int     `json:"decimals"`
				UiAmount       float64 `json:"uiAmount"`
				UiAmountString string  `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"preTokenBalances"`
		Rewards []interface{} `json:"rewards"`
		Status  struct {
			Ok interface{} `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Slot        int `json:"slot"`
	Transaction struct {
		Message struct {
			Header struct {
				NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
				NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
				NumRequiredSignatures       int `json:"numRequiredSignatures"`
			} `json:"header"`
			StaticAccountKeys    []string `json:"staticAccountKeys"`
			RecentBlockhash      string   `json:"recentBlockhash"`
			CompiledInstructions []struct {
				ProgramIdIndex    int   `json:"programIdIndex"`
				AccountKeyIndexes []int `json:"accountKeyIndexes"`
				Data              struct {
					Type string `json:"type"`
					Data []int  `json:"data"`
				} `json:"data"`
			} `json:"compiledInstructions"`
			AddressTableLookups []interface{} `json:"addressTableLookups"`
		} `json:"message"`
		Signatures []string `json:"signatures"`
	} `json:"transaction"`
	Version any `json:"version"`
}

type TronTransactionInfo struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature        []string `json:"signature"`
	BlockNumber      int      `json:"blockNumber"`
	TxID             string   `json:"txID"`
	NetFee           int      `json:"netFee"`
	Fee              int      `json:"fee"`
	EnergyFee        int      `json:"energyFee"`
	EnergyUsage      int      `json:"energyUsage"`
	EnergyUsageTotal int      `json:"energyUsageTotal"`
	RawData          struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Data                  string `json:"data"`
					OwnerAddress          string `json:"owner_address"`
					ContractAddress       string `json:"contract_address"`
					OwnerAddressBase58    string `json:"ownerAddressBase58"`
					ContractAddressBase58 string `json:"contractAddressBase58"`
				} `json:"value"`
				TypeUrl string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		FeeLimit      int    `json:"fee_limit"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"rawData"`
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	} `json:"log"`
}

type TransactionInfo struct {
	Hash        string
	BlockNumber int64

	EVM    *EVMTransactionInfo
	UTXO   *UTXOTransactionInfo
	Solana *SolanaTransactionInfo
	Tron   *TronTransactionInfo
}

func (tx *TransactionInfo) Normalize() {
	if tx.UTXO != nil {
		tx.Hash = tx.UTXO.Hash
		tx.BlockNumber = int64(tx.UTXO.BlockNumber)
	} else if tx.EVM != nil {
		tx.Hash = tx.EVM.Hash
		tx.BlockNumber = int64(tx.EVM.BlockNumber)
	} else if tx.Solana != nil {
		tx.Hash = tx.Solana.Transaction.Signatures[0]
		tx.BlockNumber = int64(tx.Solana.Slot)
	} else if tx.Tron != nil {
		tx.Hash = tx.Tron.TxID
		tx.BlockNumber = int64(tx.Tron.BlockNumber)
	}
}

type NativeBalanceRequest struct {
	Address string
	Chain   string
}

func (r NativeBalanceRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("address", r.Address)
	q.Set("chain", r.Chain)
	return q
}

type NativeBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type BatchBalanceRequest struct {
	Addresses []string
	Chain     string
}

func (r BatchBalanceRequest) toQuery() url.Values {
	q := url.Values{}
	for _, a := range r.Addresses {
		q.Add("addresses", a)
	}
	q.Set("chain", r.Chain)
	return q
}

type MempoolRequest struct {
	Chain string
}

func (r MempoolRequest) toQuery() url.Values {
	q := url.Values{}
	q.Set("chain", r.Chain)
	return q
}

type MempoolResponse struct {
	Transactions []string `json:"transactions"`
}
