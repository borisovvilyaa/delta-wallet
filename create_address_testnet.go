package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type BitcoinWallet struct {
	PrivateKey *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
	Address    string
}

func GenerateBitcoinWalletFromSeed(seed []byte) (*BitcoinWallet, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	// Вместо генерации нового приватного ключа, используем мастер-приватный ключ
	privateKeyBytes := masterKey.Key
	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	// Продолжаем как раньше
	testNetParams := chaincfg.TestNet3Params
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	segWitAddress, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &testNetParams)
	if err != nil {
		return nil, err
	}

	return &BitcoinWallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    segWitAddress.EncodeAddress(),
	}, nil
}

func main() {
	// Пример использования
	// entropy, _ := bip39.NewEntropy(128)
	mnemonic := "afford worth interest casual level service virus infant penalty crystal cave black"
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	wallet, err := GenerateBitcoinWalletFromSeed(seed)
	if err != nil {
		log.Fatal("Ошибка при генерации кошелька:", err)
	}

	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Приватный ключ:", wallet.PrivateKey)
	fmt.Println("Публичный ключ:", wallet.PublicKey)
	fmt.Println("Адрес:", wallet.Address)
}
