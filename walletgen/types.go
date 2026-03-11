package walletgen

import (
	"fmt"

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
	Key string `json:"key"`
}

var apiPaths = map[chain.Chain]string{
	chain.Bitcoin:       "bitcoin",
	chain.BitcoinCash:   "bcash",
	chain.Dogecoin:      "dogecoin",
	chain.Litecoin:      "litecoin",
	chain.Ethereum:      "ethereum",
	chain.BNBSmartChain: "bsc",
	chain.Polygon:       "polygon",
	chain.Klaytn:        "klaytn",
	chain.KuCoin:        "kcs",
	chain.VeChain:       "vet",
	chain.Harmony:       "one",
	chain.Tron:          "tron",
	chain.Algorand:      "algorand",
	chain.Solana:        "solana",
	chain.Flow:          "flow",
	chain.EGLD:          "egld",
	chain.XDC:           "xdc",
}

func chainPath(c chain.Chain) (string, error) {
	p, ok := apiPaths[c]
	if !ok {
		return "", fmt.Errorf("walletgen: unsupported chain %q", c)
	}
	return p, nil
}

var noAddressDerivation = map[chain.Chain]bool{
	chain.Solana:   true,
	chain.Algorand: true,
	chain.Flow:     true,
	chain.EGLD:     true,
}

var noPrivKeyDerivation = map[chain.Chain]bool{
	chain.Solana:   true,
	chain.Algorand: true,
	chain.Flow:     true,
}

func supportsAddressDerivation(c chain.Chain) bool { return !noAddressDerivation[c] }
func supportsPrivKeyDerivation(c chain.Chain) bool { return !noPrivKeyDerivation[c] }
