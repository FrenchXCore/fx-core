package helpers

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	hd2 "github.com/evmos/ethermint/crypto/hd"
	tronaddress "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func NewMnemonic() string {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		panic(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		panic(err)
	}
	return mnemonic
}

func PrivKeyFromMnemonic(mnemonic string, algo hd.PubKeyType, account, index uint32) (cryptotypes.PrivKey, error) {
	var hdPath *hd.BIP44Params
	if algo == hd.Secp256k1Type {
		hdPath = hd.CreateHDPath(118, account, index)
	} else if algo == hd2.EthSecp256k1Type {
		hdPath = hd.CreateHDPath(60, account, index)
	} else {
		return nil, fmt.Errorf("invalid algo")
	}
	signAlgo, err := keyring.NewSigningAlgoFromString(string(algo), hd2.SupportedAlgorithms)
	if err != nil {
		return nil, err
	}
	// create master key and derive first key for keyring
	derivedPriv, err := signAlgo.Derive()(mnemonic, "", hdPath.String())
	if err != nil {
		return nil, err
	}
	privKey := signAlgo.Generate()(derivedPriv)
	return privKey, nil
}

func CreateMultiECDSA(count int) []*ecdsa.PrivateKey {
	var ethKeys []*ecdsa.PrivateKey
	for i := 0; i < count; i++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			panic(err)
		}
		ethKeys = append(ethKeys, key)
	}
	return ethKeys
}

// NewPriKey generates cosmos-sdk accAddress private key.
func NewPriKey() cryptotypes.PrivKey {
	return secp256k1.GenPrivKey()
}

// NewEthPrivKey generates an Ethereum address and its corresponding private key.
func NewEthPrivKey() cryptotypes.PrivKey {
	privkey, err := ethsecp256k1.GenerateKey()
	if err != nil {
		panic(err)
	}
	key, err := privkey.ToECDSA()
	if err != nil {
		panic(err)
	}
	addr1 := crypto.PubkeyToAddress(key.PublicKey)
	addr2 := common.BytesToAddress(privkey.PubKey().Address())
	if addr1 != addr2 {
		panic("invalid private key")
	}
	return privkey
}

// GenerateAddress generates an Ethereum address.
func GenerateAddress() common.Address {
	return common.BytesToAddress(NewEthPrivKey().PubKey().Address())
}

// GenerateAddressByModule generates an Ethereum or Tron address.
func GenerateAddressByModule(module string) string {
	addr := GenerateAddress()
	if module == "tron" {
		return tronaddress.Address(append([]byte{tronaddress.TronBytePrefix}, addr.Bytes()...)).String()
	}
	return addr.String()
}

// GenerateZeroAddressByModule generates an Ethereum or Tron zero address.
func GenerateZeroAddressByModule(module string) string {
	addr := common.HexToAddress(common.Address{}.Hex())
	if module == "tron" {
		return tronaddress.Address(append([]byte{tronaddress.TronBytePrefix}, addr.Bytes()...)).String()
	}
	return addr.String()
}

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	if len(pkBytes) != ed25519.PubKeySize {
		panic(errorsmod.Wrap(errortypes.ErrInvalidPubKey, "invalid pubkey size"))
	}
	return &ed25519.PubKey{Key: pkBytes}
}
