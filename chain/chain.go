package chain

import "fmt"

type Chain string
type Network string

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
)

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

func RPCGatewayURL(c Chain, n Network) string {
	return fmt.Sprintf("https://%s-%s.gateway.tatum.io/", c, n)
}
