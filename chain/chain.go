package chain

import (
	"fmt"
	"strings"
)

type Chain string
type Network string
type ChainNetwork string

func (cn ChainNetwork) Split() (Chain, Network) {
	parts := strings.SplitN(string(cn), "-", 2)
	return Chain(parts[0]), Network(parts[1])
}

const (
	Ethereum        Chain = "ethereum"
	Bitcoin         Chain = "bitcoin"
	Polygon         Chain = "polygon"
	BNBSmartChain   Chain = "bsc"
	Solana          Chain = "solana"
	Avalanche       Chain = "avax"
	Tron            Chain = "tron"
	Litecoin        Chain = "litecoin"
	BitcoinCash     Chain = "bch"
	Dogecoin        Chain = "doge"
	Cardano         Chain = "cardano"
	XRP             Chain = "xrp"
	Stellar         Chain = "xlm"
	Algorand        Chain = "algo"
	Tezos           Chain = "tezos"
	Monero          Chain = "xmr"
	Dash            Chain = "dash"
	ZCash           Chain = "zec"
	Optimism        Chain = "optimism"
	Arbitrum        Chain = "arbitrum"
	Cronos          Chain = "cro"
	Klaytn          Chain = "klay"
	Oasis           Chain = "oasis"
	KuCoin          Chain = "kcc"
	Aurora          Chain = "aurora"
	Celo            Chain = "celo"
	Heco            Chain = "heco"
	Near            Chain = "near"
	Palm            Chain = "palm"
	Gnosis          Chain = "gnosis"
	Theta           Chain = "theta"
	Stacks          Chain = "stacks"
	Horizen         Chain = "zen"
	Fantom          Chain = "ftm"
	Flare           Chain = "flare"
	Flow            Chain = "flow"
	EthereumClassic Chain = "etc"
	VeChain         Chain = "vet"
	IoTeX           Chain = "iotx"
	Harmony         Chain = "one"
	Hedera          Chain = "hbar"
	Chiliz          Chain = "chz"
	Base            Chain = "base"
	Scroll          Chain = "scroll"
	Linea           Chain = "linea"
	Blast           Chain = "blast"
	ZkSync          Chain = "zksync"
	EGLD            Chain = "egld"
	XDC             Chain = "xdc"
)

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

func (chainID Chain) On(network Network) ChainNetwork {
	if chainID == Ethereum && network == Testnet {
		return ChainNetwork("ethereum-sepolia")
	}

	return ChainNetwork(fmt.Sprintf("%s-%s", chainID, network))
}

func (chainID Chain) IsEVM() bool {
	return chainID == Ethereum ||
		chainID == Polygon ||
		chainID == BNBSmartChain ||
		chainID == Avalanche ||
		chainID == Optimism ||
		chainID == Arbitrum ||
		chainID == Cronos ||
		chainID == Klaytn ||
		chainID == Oasis ||
		chainID == KuCoin ||
		chainID == Aurora ||
		chainID == Celo ||
		chainID == Heco ||
		chainID == Palm ||
		chainID == Gnosis ||
		chainID == Fantom ||
		chainID == Flare ||
		chainID == EthereumClassic ||
		chainID == Base ||
		chainID == Scroll ||
		chainID == Linea ||
		chainID == Blast ||
		chainID == ZkSync
}

func RPCGatewayURL(cn ChainNetwork) string {
	return fmt.Sprintf("https://%s.gateway.tatum.io/", cn)
}
