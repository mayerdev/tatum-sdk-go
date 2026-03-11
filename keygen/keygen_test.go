package keygen

import (
	"errors"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	etherCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
)

const abandonMnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

func xpubForMnemonic(t *testing.T, mnemonic string, coinType uint32) string {
	t.Helper()
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		t.Fatalf("seed: %v", err)
	}
	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		t.Fatalf("master: %v", err)
	}
	key, err := deriveAccount(master, coinType)
	if err != nil {
		t.Fatalf("deriveAccount: %v", err)
	}
	pub, err := key.Neuter()
	if err != nil {
		t.Fatalf("neuter: %v", err)
	}
	return pub.String()
}

func TestGenerateWallet(t *testing.T) {
	chains := []chain.Chain{chain.Ethereum, chain.Bitcoin, chain.Tron}
	for _, c := range chains {
		w, err := GenerateWallet(c)
		if err != nil {
			t.Fatalf("GenerateWallet(%s): %v", c, err)
		}
		if w.Mnemonic == "" {
			t.Errorf("GenerateWallet(%s): empty Mnemonic", c)
		}
		if w.XPub == "" {
			t.Errorf("GenerateWallet(%s): empty XPub", c)
		}
		words := strings.Fields(w.Mnemonic)
		if len(words) != 24 {
			t.Errorf("GenerateWallet(%s): got %d words, want 24", c, len(words))
		}
	}
}

func TestGenerateWallet_Solana(t *testing.T) {
	w, err := GenerateWallet(chain.Solana)
	if err != nil {
		t.Fatalf("GenerateWallet(Solana): %v", err)
	}
	if w.Mnemonic == "" || w.XPub == "" {
		t.Error("GenerateWallet(Solana): empty Mnemonic or XPub")
	}
}

func TestDeriveAddress_ETH_KnownVector(t *testing.T) {
	xpub := xpubForMnemonic(t, abandonMnemonic, coinETH)
	addr, err := DeriveAddress(xpub, 0, chain.Ethereum)
	if err != nil {
		t.Fatalf("DeriveAddress: %v", err)
	}
	const want = "0x9858EfFD232B4033E47d90003D41EC34EcaEda94"
	if addr != want {
		t.Errorf("got %s, want %s", addr, want)
	}
}

func TestDeriveAddress(t *testing.T) {
	tests := []struct {
		name     string
		c        chain.Chain
		coinType uint32
		prefix   string
	}{
		{"Bitcoin", chain.Bitcoin, coinBTC, "1"},
		{"Litecoin", chain.Litecoin, coinLTC, "L"},
		{"Dogecoin", chain.Dogecoin, coinDOGE, "D"},
		{"Dash", chain.Dash, coinDASH, "X"},
		{"ZCash", chain.ZCash, coinZEC, "t1"},
		{"Tron", chain.Tron, coinTRON, "T"},
		{"XRP", chain.XRP, coinXRP, "r"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xpub := xpubForMnemonic(t, abandonMnemonic, tt.coinType)
			addr, err := DeriveAddress(xpub, 0, tt.c)
			if err != nil {
				t.Fatalf("DeriveAddress: %v", err)
			}
			if addr == "" {
				t.Fatal("empty address")
			}
			if !strings.HasPrefix(addr, tt.prefix) {
				t.Errorf("address %q does not start with %q", addr, tt.prefix)
			}
		})
	}
}

func TestDerivePrivateKey(t *testing.T) {
	t.Run("Ethereum_roundtrip", func(t *testing.T) {
		w, err := GenerateWallet(chain.Ethereum)
		if err != nil {
			t.Fatal(err)
		}
		addr, err := DeriveAddress(w.XPub, 0, chain.Ethereum)
		if err != nil {
			t.Fatal(err)
		}
		privHex, err := DerivePrivateKey(w.Mnemonic, 0, chain.Ethereum)
		if err != nil {
			t.Fatal(err)
		}
		privBytes, err := etherCrypto.HexToECDSA(strings.TrimPrefix(privHex, "0x"))
		if err != nil {
			t.Fatalf("HexToECDSA: %v", err)
		}
		derived := etherCrypto.PubkeyToAddress(privBytes.PublicKey).Hex()
		if derived != addr {
			t.Errorf("roundtrip mismatch: DeriveAddress=%s, from privkey=%s", addr, derived)
		}
	})

	t.Run("Bitcoin_roundtrip", func(t *testing.T) {
		w, err := GenerateWallet(chain.Bitcoin)
		if err != nil {
			t.Fatal(err)
		}
		addr, err := DeriveAddress(w.XPub, 0, chain.Bitcoin)
		if err != nil {
			t.Fatal(err)
		}
		wifStr, err := DerivePrivateKey(w.Mnemonic, 0, chain.Bitcoin)
		if err != nil {
			t.Fatal(err)
		}
		wif, err := btcutil.DecodeWIF(wifStr)
		if err != nil {
			t.Fatalf("DecodeWIF: %v", err)
		}
		compressed := wif.PrivKey.PubKey().SerializeCompressed()
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, &chaincfg.MainNetParams)
		if err != nil {
			t.Fatal(err)
		}
		if a.EncodeAddress() != addr {
			t.Errorf("roundtrip mismatch: DeriveAddress=%s, from WIF=%s", addr, a.EncodeAddress())
		}
	})

	t.Run("Tron_deterministic", func(t *testing.T) {
		k1, err := DerivePrivateKey(abandonMnemonic, 0, chain.Tron)
		if err != nil {
			t.Fatal(err)
		}
		k2, err := DerivePrivateKey(abandonMnemonic, 0, chain.Tron)
		if err != nil {
			t.Fatal(err)
		}
		if k1 != k2 {
			t.Error("non-deterministic: same inputs produced different keys")
		}
		if !strings.HasPrefix(k1, "0x") {
			t.Errorf("expected 0x prefix, got %s", k1)
		}
	})
}

func TestUnsupportedChain(t *testing.T) {
	unsupported := []chain.Chain{chain.Cardano, chain.Monero}
	for _, c := range unsupported {
		_, err := GenerateWallet(c)
		if !errors.Is(err, ErrUnsupportedChain) {
			t.Errorf("GenerateWallet(%s): want ErrUnsupportedChain, got %v", c, err)
		}
		_, err = DeriveAddress("xpub", 0, c)
		if !errors.Is(err, ErrUnsupportedChain) {
			t.Errorf("DeriveAddress(%s): want ErrUnsupportedChain, got %v", c, err)
		}
		_, err = DerivePrivateKey(abandonMnemonic, 0, c)
		if !errors.Is(err, ErrUnsupportedChain) {
			t.Errorf("DerivePrivateKey(%s): want ErrUnsupportedChain, got %v", c, err)
		}
	}
}
