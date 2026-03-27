package walletgen

import (
	"gitlab.com/mayerdev/tatum-sdk-go/chain"
)

type WalletResponse struct {
	Mnemonic   string `json:"mnemonic"`
	XPub       string `json:"xpub"`
	Address    string `json:"address"`    // Algorand, Solana, Flow
	PrivateKey string `json:"privateKey"` // Solana
	Secret     string `json:"secret"`     // Algorand
}

type addressResponse struct {
	Address string `json:"address"`
}

type privKeyRequest struct {
	Mnemonic string `json:"mnemonic"`
	Index    uint32 `json:"index"`
}

type privKeyResponse struct {
	Key        string `json:"key"`
	PrivateKey string `json:"privateKey"`
}

func (r privKeyResponse) value() string {
	if r.PrivateKey != "" {
		return r.PrivateKey
	}
	return r.Key
}

var apiPaths = map[chain.Chain]string{
	chain.BitcoinCash: "bcash",
	chain.Dogecoin:    "dogecoin",
	chain.KuCoin:      "kcs",
	chain.Algorand:    "algorand",
	chain.Arbitrum:    "arb",
	chain.Avalanche:   "avalanche",
	chain.Fantom:      "fantom",
}

func chainPath(c chain.Chain) string {
	p, ok := apiPaths[c]
	if !ok {
		return string(c)
	}

	return p
}

var customAddressPath = map[chain.Chain]string{
	chain.Flow: "/v3/flow/pubkey/%s/%d",
}
