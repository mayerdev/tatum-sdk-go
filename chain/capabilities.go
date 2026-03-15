package chain

type DeriveType int

const (
	DeriveHD DeriveType = iota
	DeriveAlt
)

type CapabilitiesInfo struct {
	Derivable  DeriveType
	DerivePath string
}

var emvCap = CapabilitiesInfo{
	Derivable:  DeriveHD,
	DerivePath: "m/44'/60'/0'/0",
}

var capabilities = map[Chain]CapabilitiesInfo{
	Bitcoin: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/0'/0'/0",
	}, // legacy
	Litecoin: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/2'/0'/0",
	},
	BitcoinCash: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/145'/0'/0",
	},
	Dogecoin: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/3'/0'/0",
	},
	Dash: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/5'/0'/0",
	},
	ZCash: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/133'/0'/0",
	},
	Stacks: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/5757'/0'/0",
	},
	Horizen: {
		Derivable:  DeriveHD,
		DerivePath: "m/44'/121'/0'/0",
	},

	Ethereum:        emvCap,
	Polygon:         emvCap,
	BNBSmartChain:   emvCap,
	Avalanche:       emvCap,
	Optimism:        emvCap,
	Arbitrum:        emvCap,
	Cronos:          emvCap,
	Klaytn:          emvCap,
	Oasis:           emvCap,
	KuCoin:          emvCap,
	Aurora:          emvCap,
	Celo:            emvCap,
	Heco:            emvCap,
	Palm:            emvCap,
	Gnosis:          emvCap,
	Fantom:          emvCap,
	Flare:           emvCap,
	EthereumClassic: emvCap,
	Base:            emvCap,
	ZkSync:          emvCap,
}
