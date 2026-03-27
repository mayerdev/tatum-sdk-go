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
	Algorand        Chain = "algo"
	Arbitrum        Chain = "arbitrum"
	Aurora          Chain = "aurora"
	Avalanche       Chain = "avax"
	BNBSmartChain   Chain = "bsc"
	Base            Chain = "base"
	Bitcoin         Chain = "bitcoin"
	BitcoinCash     Chain = "bch"
	Cardano         Chain = "cardano"
	Celo            Chain = "celo"
	Chiliz          Chain = "chz"
	Cronos          Chain = "cro"
	Dash            Chain = "dash"
	Dogecoin        Chain = "doge"
	EGLD            Chain = "egld"
	Ethereum        Chain = "ethereum"
	EthereumClassic Chain = "etc"
	Fantom          Chain = "ftm"
	Flare           Chain = "flare"
	Flow            Chain = "flow"
	Gnosis          Chain = "gnosis"
	Harmony         Chain = "one"
	Heco            Chain = "heco"
	Hedera          Chain = "hbar"
	Horizen         Chain = "zen"
	IoTeX           Chain = "iotx"
	Klaytn          Chain = "klay"
	KuCoin          Chain = "kcc"
	Litecoin        Chain = "litecoin"
	Monero          Chain = "xmr"
	Near            Chain = "near"
	Oasis           Chain = "oasis"
	Optimism        Chain = "optimism"
	Palm            Chain = "palm"
	Polygon         Chain = "polygon"
	Ripple          Chain = "xrp"
	Solana          Chain = "solana"
	Stacks          Chain = "stacks"
	Stellar         Chain = "xlm"
	Tezos           Chain = "tezos"
	Theta           Chain = "theta"
	Tron            Chain = "tron"
	VeChain         Chain = "vet"
	XDC             Chain = "xdc"
	ZCash           Chain = "zec"
	ZkSync          Chain = "zksync"
)

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
)

func (chainID Chain) On(network Network) ChainNetwork {
	if (chainID == Ethereum || chainID == Base) && network == Testnet {
		return ChainNetwork(chainID + "-sepolia")
	} else if chainID == Solana && network == Testnet {
		return "solana-devnet"
	} else if chainID == Arbitrum {
		return "arb-one-mainnet"
	} else if chainID == Chiliz {
		return ChainNetwork("chiliz-" + network)
	} else if chainID == Fantom {
		return ChainNetwork("fantom-" + network)
	} else if chainID == Flare && network == Testnet {
		return "flare-coston2"
	} else if chainID == Klaytn {
		if network == Testnet {
			return "klaytn-baobab"
		}

		return "klaytn-cypress"
	} else if chainID == Litecoin {
		return ChainNetwork("litecoin-core-" + network)
	} else if chainID == Polygon && network == Testnet {
		return "polygon-amoy"
	} else if chainID == Ripple {
		return ChainNetwork("ripple-" + network)
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
		chainID == ZkSync
}

func (chainID Chain) IsUTXO() bool {
	return chainID == Bitcoin ||
		chainID == Litecoin ||
		chainID == BitcoinCash ||
		chainID == Dogecoin ||
		chainID == Dash ||
		chainID == ZCash ||
		chainID == Cardano
}

func (chainID Chain) GetCapabilities() *CapabilitiesInfo {
	info, ok := capabilities[chainID]
	if !ok {
		return nil
	}

	return &info
}

func RPCGatewayURL(cn ChainNetwork) string {
	return fmt.Sprintf("https://%s.gateway.tatum.io/", cn)
}
