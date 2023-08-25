package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func GenerateBitcoinWallet() (*BitcoinWallet, string, error) {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	masterKey, err := bip32.NewMasterKey(entropy)
	if err != nil {
		return nil, "", err // Return empty string for mnemonic as we encountered an error
	}

	privateKeyBytes := masterKey.Key
	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	NetParams := chaincfg.TestNet3Params
	pubKeyHash := btcutil.Hash160(publicKey.SerializeCompressed())
	segWitAddress, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &NetParams)
	if err != nil {
		return nil, "", err // Return empty string for mnemonic as we encountered an error
	}

	return &BitcoinWallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    segWitAddress.EncodeAddress(),
	}, mnemonic, nil // Return mnemonic on success
}

const keysFolder = "keys"

func ensureKeysFolderExists() error {
	// Создаем папку "keys" если она не существует
	_, err := os.Stat(keysFolder)
	if os.IsNotExist(err) {
		err := os.Mkdir(keysFolder, 0700)
		if err != nil {
			return err
		}
	}
	return nil
}
func GenerateBitcoinWalletFromSeed(seed []byte) (*BitcoinWallet, error) {
	masterKey, err := bip32.NewMasterKey(seed)
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

func SavePrivateKeyToFile(filename string, privateKey *btcec.PrivateKey) error {
	privateKeyBytes := privateKey.Serialize()
	filePath := filepath.Join(keysFolder, filename)
	return ioutil.WriteFile(filePath, privateKeyBytes, 0600)
}

func SavePassphraseToFile(filename string, mnemonic string, passphrase string) error {
	passphraseData := struct {
		Mnemonic   string
		Passphrase string
	}{
		Mnemonic:   mnemonic,
		Passphrase: passphrase,
	}

	passphraseJSON, err := json.MarshalIndent(passphraseData, "", "  ")
	if err != nil {
		return err
	}

	filePath := filepath.Join(keysFolder, filename)
	return ioutil.WriteFile(filePath, passphraseJSON, 0644)
}

func ReadPrivateKeyFromFile(filename string) (*btcec.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(filepath.Join(keysFolder, filename))
	if err != nil {
		return nil, err
	}

	// Восстановление приватного ключа из байтов
	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	if privateKey == nil {
		return nil, fmt.Errorf("failed to decode private key")
	}
	publicKey.ToECDSA()
	return privateKey, nil
}

func ReadPassphraseFromFile(filename string) (string, string, error) {
	passphraseData := struct {
		Mnemonic   string
		Passphrase string
	}{}

	passphraseJSON, err := ioutil.ReadFile(filepath.Join(keysFolder, filename))
	if err != nil {
		return "", "", err
	}

	if err := json.Unmarshal(passphraseJSON, &passphraseData); err != nil {
		return "", "", err
	}

	return passphraseData.Mnemonic, passphraseData.Passphrase, nil
}

// func main() {
// 	mnemonic := "afford worth interest casual level service virus infant penalty crystal cave black"
// 	seed := bip39.NewSeed(mnemonic, "delta")

// 	wallet, err := GenerateBitcoinWalletFromSeed(seed)
// 	if err != nil {
// 		log.Fatal("Ошибка при генерации кошелька:", err)
// 	}

// 	fmt.Println("Mnemonic:", mnemonic)
// 	fmt.Println("Приватный ключ:", wallet.PrivateKey.ToECDSA().D.Text(16))
// 	fmt.Println("Публичный ключ:", wallet.PublicKey)
// 	fmt.Println("Адрес:", wallet.Address)

// 	if err := SavePrivateKeyToFile("seed.seco", wallet.PrivateKey); err != nil {
// 		log.Fatal("Ошибка при сохранении приватного ключа:", err)
// 	}

// 	if err := SavePassphraseToFile("passphrase.json", mnemonic, "delta"); err != nil {
// 		log.Fatal("Ошибка при сохранении JSON в файл:", err)
// 	}

// 	readPrivateKey, err := ReadPrivateKeyFromFile("seed.seco")
// 	if err != nil {
// 		log.Fatal("Ошибка при чтении приватного ключа из файла:", err)
// 	}
// 	fmt.Println("\nприватный ключ:", readPrivateKey.ToECDSA().D.Text(16))

// 	readMnemonic, readPassphrase, err := ReadPassphraseFromFile("passphrase.json")
// 	if err != nil {
// 		log.Fatal("Ошибка при чтении JSON из файла:", err)
// 	}
// 	fmt.Println("Прочитанный mnemonic:", readMnemonic)
// 	fmt.Println("Прочитанный passphrase:", readPassphrase)
// }
