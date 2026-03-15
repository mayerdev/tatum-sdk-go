package keygen

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	etherCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"

	"gitlab.com/mayerdev/tatum-sdk-go/chain"
)

var ErrUnsupportedChain = errors.New("keygen: unsupported chain")

var ErrSolanaPublicDerivation = errors.New(
	"keygen: Solana does not support public key derivation; " +
		"use DerivePrivateKey to get the private key at the desired index")

type Wallet struct {
	Mnemonic string
	XPub     string
}

const hardenedOffset = uint32(0x80000000)

const (
	coinBTC  = uint32(0)
	coinLTC  = uint32(2)
	coinDOGE = uint32(3)
	coinDASH = uint32(5)
	coinETH  = uint32(60)
	coinZEC  = uint32(133)
	coinXRP  = uint32(144)
	coinBCH  = uint32(145)
	coinTRON = uint32(195)
	coinSOL  = uint32(501)
)

type chainFamily int

const (
	familyEVM chainFamily = iota
	familyBTC
	familyLTC
	familyDOGE
	familyBCH
	familyDASH
	familyZEC
	familyTRON
	familySOL
	familyXRP
)

type chainInfo struct {
	coinType uint32
	family   chainFamily
}

var (
	ltcAddrParams  = &chaincfg.Params{PubKeyHashAddrID: 0x30, PrivateKeyID: 0xB0}
	dogeAddrParams = &chaincfg.Params{PubKeyHashAddrID: 0x1E, PrivateKeyID: 0x9E}
	bchAddrParams  = &chaincfg.Params{PubKeyHashAddrID: 0x00, PrivateKeyID: 0x80}
	dashAddrParams = &chaincfg.Params{PubKeyHashAddrID: 0x4C, PrivateKeyID: 0xCC}
)

const (
	btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	xrpAlphabet = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
)

func getChainInfo(c chain.Chain) (chainInfo, error) {
	switch c {
	case chain.Ethereum, chain.Polygon, chain.BNBSmartChain, chain.Avalanche,
		chain.Arbitrum, chain.Optimism, chain.Cronos, chain.Celo, chain.Gnosis,
		chain.Base, chain.ZkSync,
		chain.Fantom, chain.Harmony, chain.Aurora, chain.Heco, chain.KuCoin,
		chain.Klaytn, chain.Palm, chain.EthereumClassic, chain.VeChain,
		chain.Hedera, chain.Chiliz, chain.Theta:
		return chainInfo{coinType: coinETH, family: familyEVM}, nil
	case chain.Bitcoin:
		return chainInfo{coinType: coinBTC, family: familyBTC}, nil
	case chain.Litecoin:
		return chainInfo{coinType: coinLTC, family: familyLTC}, nil
	case chain.Dogecoin:
		return chainInfo{coinType: coinDOGE, family: familyDOGE}, nil
	case chain.BitcoinCash:
		return chainInfo{coinType: coinBCH, family: familyBCH}, nil
	case chain.Dash:
		return chainInfo{coinType: coinDASH, family: familyDASH}, nil
	case chain.ZCash:
		return chainInfo{coinType: coinZEC, family: familyZEC}, nil
	case chain.Tron:
		return chainInfo{coinType: coinTRON, family: familyTRON}, nil
	case chain.Solana:
		return chainInfo{coinType: coinSOL, family: familySOL}, nil
	case chain.Ripple:
		return chainInfo{coinType: coinXRP, family: familyXRP}, nil
	default:
		return chainInfo{}, ErrUnsupportedChain
	}
}

func GenerateWallet(c chain.Chain) (*Wallet, error) {
	info, err := getChainInfo(c)
	if err != nil {
		return nil, err
	}

	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	if info.family == familySOL {
		xpub, err := solanaXPub(seed)
		if err != nil {
			return nil, err
		}
		return &Wallet{Mnemonic: mnemonic, XPub: xpub}, nil
	}

	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	account, err := deriveAccount(master, info.coinType)
	if err != nil {
		return nil, err
	}

	pub, err := account.Neuter()
	if err != nil {
		return nil, err
	}

	return &Wallet{Mnemonic: mnemonic, XPub: pub.String()}, nil
}

func DeriveAddress(xpub string, index uint32, c chain.Chain) (string, error) {
	info, err := getChainInfo(c)
	if err != nil {
		return "", err
	}

	if info.family == familySOL {
		if index != 0 {
			return "", ErrSolanaPublicDerivation
		}
		return xpub, nil
	}

	key, err := hdkeychain.NewKeyFromString(xpub)
	if err != nil {
		return "", err
	}

	child, err := key.Derive(index)
	if err != nil {
		return "", err
	}

	ecPub, err := child.ECPubKey()
	if err != nil {
		return "", err
	}
	compressed := ecPub.SerializeCompressed()

	switch info.family {
	case familyEVM:
		pubKey, err := etherCrypto.DecompressPubkey(compressed)
		if err != nil {
			return "", err
		}
		return etherCrypto.PubkeyToAddress(*pubKey).Hex(), nil

	case familyTRON:
		pubKey, err := etherCrypto.DecompressPubkey(compressed)
		if err != nil {
			return "", err
		}
		addr := etherCrypto.PubkeyToAddress(*pubKey)
		return base58.CheckEncode(addr.Bytes(), 0x41), nil

	case familyBTC:
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		return a.EncodeAddress(), nil

	case familyLTC:
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, ltcAddrParams)
		if err != nil {
			return "", err
		}
		return a.EncodeAddress(), nil

	case familyDOGE:
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, dogeAddrParams)
		if err != nil {
			return "", err
		}
		return a.EncodeAddress(), nil

	case familyBCH:
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, bchAddrParams)
		if err != nil {
			return "", err
		}
		return a.EncodeAddress(), nil

	case familyDASH:
		hash := btcutil.Hash160(compressed)
		a, err := btcutil.NewAddressPubKeyHash(hash, dashAddrParams)
		if err != nil {
			return "", err
		}
		return a.EncodeAddress(), nil

	case familyZEC:
		return zecAddress(btcutil.Hash160(compressed)), nil

	case familyXRP:
		return xrpAddress(btcutil.Hash160(compressed)), nil
	}

	return "", ErrUnsupportedChain
}

func DerivePrivateKey(mnemonic string, index uint32, c chain.Chain) (string, error) {
	info, err := getChainInfo(c)
	if err != nil {
		return "", err
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}

	if info.family == familySOL {
		privKey, err := slip10Derive(seed, []uint32{
			44 + hardenedOffset,
			coinSOL + hardenedOffset,
			index + hardenedOffset,
			hardenedOffset,
		})
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(privKey), nil
	}

	master, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	key, err := deriveIndex(master, info.coinType, index)
	if err != nil {
		return "", err
	}

	ecKey, err := key.ECPrivKey()
	if err != nil {
		return "", err
	}

	switch info.family {
	case familyEVM, familyTRON, familyXRP:
		return "0x" + hex.EncodeToString(ecKey.Serialize()), nil

	case familyBTC, familyZEC:
		wif, err := btcutil.NewWIF(ecKey, &chaincfg.MainNetParams, true)
		if err != nil {
			return "", err
		}
		return wif.String(), nil

	case familyLTC:
		wif, err := btcutil.NewWIF(ecKey, ltcAddrParams, true)
		if err != nil {
			return "", err
		}
		return wif.String(), nil

	case familyDOGE:
		wif, err := btcutil.NewWIF(ecKey, dogeAddrParams, true)
		if err != nil {
			return "", err
		}
		return wif.String(), nil

	case familyBCH:
		wif, err := btcutil.NewWIF(ecKey, bchAddrParams, true)
		if err != nil {
			return "", err
		}
		return wif.String(), nil

	case familyDASH:
		wif, err := btcutil.NewWIF(ecKey, dashAddrParams, true)
		if err != nil {
			return "", err
		}
		return wif.String(), nil
	}

	return "", ErrUnsupportedChain
}

func deriveAccount(master *hdkeychain.ExtendedKey, coinType uint32) (*hdkeychain.ExtendedKey, error) {
	indices := []uint32{
		44 + hardenedOffset,
		coinType + hardenedOffset,
		hardenedOffset,
		0,
	}
	key := master
	for _, i := range indices {
		var err error
		key, err = key.Derive(i)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func deriveIndex(master *hdkeychain.ExtendedKey, coinType, index uint32) (*hdkeychain.ExtendedKey, error) {
	indices := []uint32{
		44 + hardenedOffset,
		coinType + hardenedOffset,
		hardenedOffset,
		0,
		index,
	}
	key := master
	for _, i := range indices {
		var err error
		key, err = key.Derive(i)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func slip10Derive(seed []byte, path []uint32) ([]byte, error) {
	mac := hmac.New(sha512.New, []byte("ed25519 seed"))
	mac.Write(seed)
	I := mac.Sum(nil)
	IL, IR := I[:32], I[32:]

	for _, index := range path {
		if index < hardenedOffset {
			return nil, errors.New("keygen: ed25519 SLIP-0010 requires hardened derivation")
		}
		mac = hmac.New(sha512.New, IR)
		mac.Write([]byte{0x00})
		mac.Write(IL)
		idxBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(idxBytes, index)
		mac.Write(idxBytes)
		I = mac.Sum(nil)
		IL, IR = I[:32], I[32:]
	}
	return IL, nil
}

func solanaXPub(seed []byte) (string, error) {
	privKey, err := slip10Derive(seed, []uint32{
		44 + hardenedOffset,
		coinSOL + hardenedOffset,
		hardenedOffset,
		hardenedOffset,
	})
	if err != nil {
		return "", err
	}
	pub := ed25519.NewKeyFromSeed(privKey).Public().(ed25519.PublicKey)
	return base58Encode(btcAlphabet, pub), nil
}

func SolanaAddressFromPrivKey(hexPrivKey string) (string, error) {
	b, err := hex.DecodeString(hexPrivKey)
	if err != nil {
		return "", err
	}
	pub := ed25519.NewKeyFromSeed(b).Public().(ed25519.PublicKey)
	return base58Encode(btcAlphabet, pub), nil
}

func zecAddress(hash160 []byte) string {
	return base58CheckEncode([]byte{0x1C, 0xB8}, hash160, btcAlphabet)
}

func xrpAddress(hash160 []byte) string {
	return base58CheckEncode([]byte{0x00}, hash160, xrpAlphabet)
}

func base58CheckEncode(version, payload []byte, alphabet string) string {
	data := make([]byte, 0, len(version)+len(payload))
	data = append(data, version...)
	data = append(data, payload...)
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	full := append(data, h2[:4]...) //nolint:gocritic
	return base58Encode(alphabet, full)
}

func base58Encode(alphabet string, data []byte) string {
	leadingZeros := 0
	for _, b := range data {
		if b != 0 {
			break
		}
		leadingZeros++
	}

	n := new(big.Int).SetBytes(data)
	result := []byte{}
	zero := big.NewInt(0)
	mod := new(big.Int)
	base := big.NewInt(58)

	for n.Cmp(zero) > 0 {
		n.DivMod(n, base, mod)
		result = append([]byte{alphabet[mod.Int64()]}, result...)
	}

	out := make([]byte, leadingZeros, leadingZeros+len(result))
	for i := range out {
		out[i] = alphabet[0]
	}
	return string(append(out, result...))
}
