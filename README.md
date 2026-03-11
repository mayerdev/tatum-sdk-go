# tatum-sdk-go

Go SDK for the [Tatum API](https://tatum.io) — blockchain infrastructure for developers.

## Repositories

| Platform | URL |
|----------|-----|
| GitLab (primary) | https://gitlab.com/mayerdev/tatum-sdk-go |
| GitHub (mirror) | https://github.com/mayerdev/tatum-sdk-go |
| Codeberg (mirror) | https://codeberg.org/mayerdev/tatum-sdk-go |

> Go module path: `gitlab.com/mayerdev/tatum-sdk-go`

## Installation

```bash
go get gitlab.com/mayerdev/tatum-sdk-go
```

## Quick Start

```go
import (
    tatum "gitlab.com/mayerdev/tatum-sdk-go"
    "gitlab.com/mayerdev/tatum-sdk-go/chain"
)

client, err := tatum.NewClient("your-api-key")
if err != nil {
    log.Fatal(err)
}
```

## Services

### Keygen (offline)

HD wallet generation, address and private key derivation — fully offline, no API calls.

```go
import (
    "gitlab.com/mayerdev/tatum-sdk-go/chain"
    "gitlab.com/mayerdev/tatum-sdk-go/keygen"
)
```

**Generate wallet** — returns a 24-word mnemonic and xpub:

```go
w, err := keygen.GenerateWallet(chain.Ethereum)
// w.Mnemonic — "word1 word2 ... word24"
// w.XPub     — "xpub6Ev..."
```

**Derive address** from xpub and index:

```go
addr, err := keygen.DeriveAddress(w.XPub, 0, chain.Ethereum)
```

**Derive private key** from mnemonic and index (EVM → 0x-prefixed hex, UTXO → WIF):

```go
privKey, err := keygen.DerivePrivateKey(w.Mnemonic, 0, chain.Ethereum)
```

**Supported networks:**

| Network | Chain | Address format | Key format |
|---------|-------|----------------|------------|
| Ethereum | `chain.Ethereum` | `0x...` | `0x...` hex |
| Polygon | `chain.Polygon` | `0x...` | `0x...` hex |
| BNB Smart Chain | `chain.BNBSmartChain` | `0x...` | `0x...` hex |
| Avalanche | `chain.Avalanche` | `0x...` | `0x...` hex |
| Arbitrum | `chain.Arbitrum` | `0x...` | `0x...` hex |
| Optimism | `chain.Optimism` | `0x...` | `0x...` hex |
| Base | `chain.Base` | `0x...` | `0x...` hex |
| zkSync | `chain.ZkSync` | `0x...` | `0x...` hex |
| Linea | `chain.Linea` | `0x...` | `0x...` hex |
| Scroll | `chain.Scroll` | `0x...` | `0x...` hex |
| Blast | `chain.Blast` | `0x...` | `0x...` hex |
| Fantom | `chain.Fantom` | `0x...` | `0x...` hex |
| Cronos | `chain.Cronos` | `0x...` | `0x...` hex |
| Celo | `chain.Celo` | `0x...` | `0x...` hex |
| Gnosis | `chain.Gnosis` | `0x...` | `0x...` hex |
| Harmony | `chain.Harmony` | `0x...` | `0x...` hex |
| Aurora | `chain.Aurora` | `0x...` | `0x...` hex |
| Heco | `chain.Heco` | `0x...` | `0x...` hex |
| KuCoin | `chain.KuCoin` | `0x...` | `0x...` hex |
| Klaytn | `chain.Klaytn` | `0x...` | `0x...` hex |
| Palm | `chain.Palm` | `0x...` | `0x...` hex |
| Ethereum Classic | `chain.EthereumClassic` | `0x...` | `0x...` hex |
| VeChain | `chain.VeChain` | `0x...` | `0x...` hex |
| Hedera | `chain.Hedera` | `0x...` | `0x...` hex |
| Chiliz | `chain.Chiliz` | `0x...` | `0x...` hex |
| Theta | `chain.Theta` | `0x...` | `0x...` hex |
| Tron | `chain.Tron` | `T...` | `0x...` hex |
| Bitcoin | `chain.Bitcoin` | `1...` | WIF |
| Litecoin | `chain.Litecoin` | `L...` | WIF |
| Dogecoin | `chain.Dogecoin` | `D...` | WIF |
| Bitcoin Cash | `chain.BitcoinCash` | `1...` | WIF |
| Dash | `chain.Dash` | `X...` | WIF |
| ZCash | `chain.ZCash` | `t1...` | WIF |
| Solana | `chain.Solana` | base58 pubkey (see note) | hex |
| XRP | `chain.XRP` | `r...` | `0x...` hex |

```go
chains := []chain.Chain{
    chain.Ethereum, chain.Polygon, chain.BNBSmartChain, chain.Avalanche,
    chain.Arbitrum, chain.Optimism, chain.Base, chain.ZkSync,
    chain.Linea, chain.Scroll, chain.Blast, chain.Fantom,
    chain.Cronos, chain.Celo, chain.Gnosis, chain.Harmony,
    chain.Aurora, chain.Heco, chain.KuCoin, chain.Klaytn,
    chain.Palm, chain.EthereumClassic, chain.VeChain, chain.Hedera,
    chain.Chiliz, chain.Theta, chain.Tron,
    chain.Bitcoin, chain.Litecoin, chain.Dogecoin, chain.BitcoinCash,
    chain.Dash, chain.ZCash,
    chain.Solana, chain.XRP,
}

for _, c := range chains {
    w, err := keygen.GenerateWallet(c)
    if err != nil {
        log.Fatal(err)
    }
    addr, err := keygen.DeriveAddress(w.XPub, 0, c)
    if err != nil {
        log.Fatal(err)
    }
    privKey, err := keygen.DerivePrivateKey(w.Mnemonic, 0, c)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("chain=%s\n  mnemonic=%s\n  xpub=%s\n  address=%s\n  privkey=%s\n\n",
        c, w.Mnemonic, w.XPub, addr, privKey)
}
```

> **Solana multi-address derivation:** Because SLIP-0010 with ed25519 only supports hardened child
> derivation, public-key child derivation is cryptographically impossible for Solana. As a result,
> `DeriveAddress(xpub, index, chain.Solana)` only works for `index == 0` (it returns the xpub itself).
> For `index > 0` it returns `keygen.ErrSolanaPublicDerivation`. To derive multiple Solana addresses,
> use `DerivePrivateKey` with the desired index and then convert with `SolanaAddressFromPrivKey`:
>
> ```go
> privKey, err := keygen.DerivePrivateKey(w.Mnemonic, 5, chain.Solana)
> addr, err := keygen.SolanaAddressFromPrivKey(privKey)
> ```

Unsupported networks return `keygen.ErrUnsupportedChain`:

```go
_, err := keygen.GenerateWallet(chain.Cardano)
if errors.Is(err, keygen.ErrUnsupportedChain) {
    // handle unsupported chain
}
```

### WalletGen (API-based)

Generate wallets and derive addresses/keys via Tatum's REST API (server-side, 17 chains).

**Generate wallet:**

```go
w, err := client.WalletGen.GenerateWallet(ctx, chain.Ethereum)
// w.Mnemonic — "word1 word2 ... word24"
// w.XPub     — "xpub6Ev..."

// For Solana: w.Address, w.PrivateKey
// For Algorand: w.Address, w.Secret
```

**Derive address** from xpub and index:

```go
addr, err := client.WalletGen.DeriveAddress(ctx, chain.Bitcoin, xpub, 0)
```

**Derive private key** from mnemonic and index:

```go
key, err := client.WalletGen.DerivePrivateKey(ctx, chain.Ethereum, mnemonic, 0)
if errors.Is(err, walletgen.ErrNotSupported) {
    // chain does not support this operation
}
```

**Supported chains and operations:**

| Chain | `GenerateWallet` | `DeriveAddress` | `DerivePrivateKey` |
|-------|:----------------:|:---------------:|:-----------------:|
| `chain.Bitcoin` | ✓ | ✓ | ✓ |
| `chain.BitcoinCash` | ✓ | ✓ | ✓ |
| `chain.Dogecoin` | ✓ | ✓ | ✓ |
| `chain.Litecoin` | ✓ | ✓ | ✓ |
| `chain.Ethereum` | ✓ | ✓ | ✓ |
| `chain.BNBSmartChain` | ✓ | ✓ | ✓ |
| `chain.Polygon` | ✓ | ✓ | ✓ |
| `chain.Klaytn` | ✓ | ✓ | ✓ |
| `chain.KuCoin` | ✓ | ✓ | ✓ |
| `chain.VeChain` | ✓ | ✓ | ✓ |
| `chain.Harmony` | ✓ | ✓ | ✓ |
| `chain.Tron` | ✓ | ✓ | ✓ |
| `chain.XDC` | ✓ | ✓ | ✓ |
| `chain.Solana` | ✓ | — | — |
| `chain.Algorand` | ✓ | — | — |
| `chain.Flow` | ✓ | — | — |
| `chain.EGLD` | ✓ | — | ✓ |

### Wallet

```go
items, err := client.Wallet.GetPortfolio(ctx, wallet.PortfolioRequest{
    Addresses: []string{"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"},
    Chain:     "ethereum",
})

balance, err := client.Wallet.GetBalanceByTime(ctx, wallet.BalanceByTimeRequest{
    Address: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:   "ethereum",
    Time:    1700000000,
})

utxos, err := client.Wallet.GetUTXOs(ctx, wallet.UTXORequest{
    Address: "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
    Chain:   "bitcoin",
})
```

### Transactions

```go
txs, err := client.Transactions.GetHistory(ctx, transactions.HistoryRequest{
    Address:  "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:    "ethereum",
    PageSize: 10,
})

tx, err := client.Transactions.GetByHash(ctx, transactions.ByHashRequest{
    Hash:  "0xabc123...",
    Chain: "ethereum",
})
```

### Blockchain

```go
block, err := client.Blockchain.GetCurrentBlock(ctx, blockchain.CurrentBlockRequest{
    Chain: "ethereum",
})

balance, err := client.Blockchain.GetNativeBalance(ctx, blockchain.NativeBalanceRequest{
    Address: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:   "ethereum",
})
```

### NFT

```go
meta, err := client.NFT.GetMetadata(ctx, nft.MetadataRequest{
    TokenID:  "1",
    Contract: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
    Chain:    "ethereum",
})

tokens, err := client.NFT.GetCollectionTokens(ctx, nft.CollectionTokensRequest{
    Collection: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
    Chain:      "ethereum",
    PageSize:   20,
})
```

### Token

```go
transfers, err := client.Token.GetTransfers(ctx, token.TransfersRequest{
    Address: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:   "ethereum",
})

rate, err := client.Token.GetExchangeRates(ctx, token.ExchangeRatesRequest{
    Contract: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
    Chain:    "ethereum",
})
```

### Prices

```go
rate, err := client.Prices.GetRate(ctx, prices.RateRequest{
    Currency: "BTC",
    BasePair: "USD",
})

candles, err := client.Prices.GetOHLCV(ctx, prices.OHLCVRequest{
    Currency: "ETH",
    BasePair: "USD",
    From:     1700000000,
    To:       1700086400,
})
```

### Staking

```go
assets, err := client.Staking.GetStakedAssets(ctx, staking.StakedAssetsRequest{
    Address: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:   "ethereum",
})

validators, err := client.Staking.GetValidators(ctx, staking.ValidatorsRequest{
    Chain: "ethereum",
})
```

### DeFi

```go
events, err := client.DeFi.GetEvents(ctx, defi.EventsRequest{
    Contract: "0xUniswapV3Pool...",
    Chain:    "ethereum",
})

latest, err := client.DeFi.GetLatestBlock(ctx, defi.LatestBlockRequest{
    Chain: "ethereum",
})
```

### Fees

```go
gas, err := client.Fees.EstimateGas(ctx, fees.EstimateGasRequest{
    Chain: "ethereum",
    From:  "0xSender",
    To:    "0xReceiver",
    Value: "1000000000000000000",
})

fee, err := client.Fees.GetBlockchainFee(ctx, fees.BlockchainFeeRequest{
    Chain: "ethereum",
})
```

### Name Service

```go
resolved, err := client.NameService.Resolve(ctx, nameservice.ResolveRequest{
    Name:  "vitalik.eth",
    Chain: "ethereum",
})
```

### Security

```go
check, err := client.Security.CheckAddress(ctx, security.CheckAddressRequest{
    Address: "0xSomeAddress",
    Chain:   "ethereum",
})
if check.Malicious {
    log.Println("address flagged as malicious:", check.Tags)
}
```

### Notifications

```go
sub, err := client.Notifications.Create(ctx, notifications.CreateRequest{
    Type:    "ADDRESS_TRANSACTION",
    Address: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    Chain:   "ethereum",
    URL:     "https://your-webhook.example.com/hook",
})

list, err := client.Notifications.List(ctx, notifications.ListRequest{PageSize: 10})

err = client.Notifications.Cancel(ctx, notifications.CancelRequest{ID: sub.ID})
```

### RPC Gateway

```go
result, err := client.RPC.Call(ctx, chain.Ethereum, chain.Mainnet, "eth_blockNumber", nil)

result, err := client.RPC.Call(ctx, chain.Ethereum, chain.Mainnet, "eth_getBalance", []any{
    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    "latest",
})
```

Supported chains:

```go
chain.Ethereum, chain.Bitcoin, chain.Polygon, chain.BNBSmartChain,
chain.Solana, chain.Avalanche, chain.Tron, chain.Optimism,
chain.Arbitrum, chain.Base, chain.ZkSync, chain.Linea,
// ... and 35+ more
```

## Error Handling

```go
_, err := client.Wallet.GetPortfolio(ctx, req)
if err != nil {
    if tatum.IsNotFound(err) {
        // 404
    }
    if tatum.IsRateLimit(err) {
        // 429 — SDK automatically retries with exponential backoff
    }
    var apiErr *tatum.APIError
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode, apiErr.ErrorCode, apiErr.Message)
    }
}
```

The client automatically retries on 429, 502, 503, 504 responses using exponential backoff.

## Pagination

```go
txs, err := client.Transactions.GetHistory(ctx, transactions.HistoryRequest{
    Address:  "0x...",
    Chain:    "ethereum",
    PageSize: 50,
    Cursor:   "",
})
```

Pass the cursor from the previous response to fetch the next page.

## Client Options

```go
client, err := tatum.NewClient("your-api-key",
    tatum.WithBaseURL("https://api.tatum.io"),
    tatum.WithTimeout(15*time.Second),
    tatum.WithRetries(5),
    tatum.WithUserAgent("my-app/1.0"),
    tatum.WithRPCBaseURL("https://%s-%s.gateway.tatum.io/"),
)
```

## Requirements

- Go 1.25+
- Tatum API key: [https://tatum.io](https://tatum.io)

## License

MIT
