package services

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type SerializedBitcoinWallet struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string
}
type BitcoinWallet struct {
	PrivateKey *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
	Address    string
}

func GenerateBitcoinWallet() (*BitcoinWallet, error) {
	entropy, _ := bip39.NewEntropy(128)
	masterKey, err := bip32.NewMasterKey(entropy)
	if err != nil {
		return nil, err
	}

	privateKeyBytes := masterKey.Key
	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	NetParams := chaincfg.TestNet3Params
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	segWitAddress, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &NetParams)
	if err != nil {
		return nil, err
	}

	return &BitcoinWallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    segWitAddress.EncodeAddress(),
	}, nil
}
