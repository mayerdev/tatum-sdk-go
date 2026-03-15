package chain

func (chainID Chain) GetNativeCurrency() string {
	switch chainID {
	case Bitcoin:
		return "BTC"
	case Litecoin:
		return "LTC"
	case BitcoinCash:
		return "BCH"
	case Dogecoin:
		return "DOGE"
	case Dash:
		return "DASH"
	case ZCash:
		return "ZEC"
	case Stacks:
		return "STX"
	case Horizen:
		return "ZEN"
	case Monero:
		return "XMR"
	case Ethereum:
		return "ETH"
	case Polygon:
		return "MATIC"
	case BNBSmartChain:
		return "BNB"
	case Avalanche:
		return "AVAX"
	case Optimism:
		return "ETH"
	case Arbitrum:
		return "ETH"
	case Cronos:
		return "CRO"
	case Klaytn:
		return "KLAY"
	case Oasis:
		return "ROSE"
	case KuCoin:
		return "KCS"
	case Aurora:
		return "ETH"
	case Celo:
		return "CELO"
	case Heco:
		return "HT"
	case Palm:
		return "PALM"
	case Gnosis:
		return "xDAI"
	case Fantom:
		return "FTM"
	case Flare:
		return "FLR"
	case EthereumClassic:
		return "ETC"
	case Base:
		return "ETH"
	case ZkSync:
		return "ETH"
	case XDC:
		return "XDC"
	case Chiliz:
		return "CHZ"
	case VeChain:
		return "VET"
	case IoTeX:
		return "IOTX"
	case Harmony:
		return "ONE"
	case Hedera:
		return "HBAR"
	case EGLD:
		return "EGLD"
	case Theta:
		return "THETA"
	case Flow:
		return "FLOW"
	case Solana:
		return "SOL"
	case Algorand:
		return "ALGO"
	case Cardano:
		return "ADA"
	case Ripple:
		return "XRP"
	case Stellar:
		return "XLM"
	case Tezos:
		return "XTZ"
	default:
		return ""
	}
}

func (chainID Chain) GetTatumNativeCurrency() string {
	switch chainID {
	case Optimism:
		return "ETH_OP"
	case Arbitrum:
		return "ETH_ARB"
	case Base:
		return "ETH_BASE"
	default:
		return chainID.GetNativeCurrency()
	}
}

func (chainID Chain) RemapCurrency(currency string) string {
	switch chainID {
	case Ethereum:
		if currency == "MATIC" || currency == "POL" {
			return currency + "_ETH"
		}
	case BNBSmartChain:
		if currency == "USDC" || currency == "USDT" || currency == "B2U" || currency == "GMC" || currency == "BUSD" {
			return currency + "_BSC"
		}
	case Arbitrum:
		if currency == "USDC" || currency == "USDT" {
			return currency + "_ARB"
		}
	case Optimism:
		if currency == "USDC" || currency == "USDT" {
			return currency + "_OP"
		}
	case Base:
		if currency == "USDC" || currency == "USDT" {
			return currency + "_BASE"
		}
	case Polygon:
		if currency == "USDC" || currency == "USDT" {
			return currency + "_MATIC"
		}
	}

	return currency
}
