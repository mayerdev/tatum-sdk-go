package chain

func (chainID Chain) GetNativeCurrency() string {
	switch chainID {
	case Algorand:
		return "ALGO"
	case Arbitrum:
		return "ETH"
	case Aurora:
		return "ETH"
	case Avalanche:
		return "AVAX"
	case BNBSmartChain:
		return "BNB"
	case Base:
		return "ETH"
	case Bitcoin:
		return "BTC"
	case BitcoinCash:
		return "BCH"
	case Cardano:
		return "ADA"
	case Celo:
		return "CELO"
	case Chiliz:
		return "CHZ"
	case Cronos:
		return "CRO"
	case Dash:
		return "DASH"
	case Dogecoin:
		return "DOGE"
	case EGLD:
		return "EGLD"
	case Ethereum:
		return "ETH"
	case EthereumClassic:
		return "ETC"
	case Fantom:
		return "FTM"
	case Flare:
		return "FLR"
	case Flow:
		return "FLOW"
	case Gnosis:
		return "xDAI"
	case Harmony:
		return "ONE"
	case Heco:
		return "HT"
	case Hedera:
		return "HBAR"
	case Horizen:
		return "ZEN"
	case IoTeX:
		return "IOTX"
	case Klaytn:
		return "KLAY"
	case KuCoin:
		return "KCS"
	case Litecoin:
		return "LTC"
	case Monero:
		return "XMR"
	case Near:
		return "NEAR"
	case Oasis:
		return "ROSE"
	case Optimism:
		return "ETH"
	case Palm:
		return "PALM"
	case Polygon:
		return "MATIC"
	case Ripple:
		return "XRP"
	case Solana:
		return "SOL"
	case Stacks:
		return "STX"
	case Stellar:
		return "XLM"
	case Tezos:
		return "XTZ"
	case Theta:
		return "THETA"
	case Tron:
		return "TRX"
	case VeChain:
		return "VET"
	case XDC:
		return "XDC"
	case ZCash:
		return "ZEC"
	case ZkSync:
		return "ETH"
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
