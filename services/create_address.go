package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

func GenerateBitcoinWalletFromSeed(seed []byte) (*BitcoinWallet, error) {
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	privateKey, _ := btcec.NewPrivateKey() // Create a new private key
	privateKeyBytes := masterKey.Key
	copy(privateKey.ToECDSA().D.Bytes(), privateKeyBytes) // Copy the master private key bytes to the new private key

	publicKey := privateKey.PubKey()

	address, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(publicKey.SerializeCompressed()), &chaincfg.TestNet3Params)
	if err != nil {
		return nil, err
	}

	return &BitcoinWallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address.EncodeAddress(),
	}, nil
}

func SerializeBitcoinWallet(wallet *BitcoinWallet) ([]byte, error) {
	serializedWallet := SerializedBitcoinWallet{
		PrivateKey: wallet.PrivateKey.Serialize(),
		PublicKey:  wallet.PublicKey.SerializeCompressed(),
		Address:    wallet.Address,
	}
	return json.Marshal(serializedWallet)
}

func SaveToFile(data []byte, filename string) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func ReadFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func DeserializeBitcoinWallet(data []byte) (*BitcoinWallet, error) {
	var serializedWallet SerializedBitcoinWallet
	err := json.Unmarshal(data, &serializedWallet)
	if err != nil {
		return nil, err
	}

	privateKey, publicKey := btcec.PrivKeyFromBytes(serializedWallet.PrivateKey)
	address, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(publicKey.SerializeCompressed()), &chaincfg.TestNet3Params)
	if err != nil {
		return nil, err
	}

	return &BitcoinWallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address.EncodeAddress(),
	}, nil
}

func main() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	wallet, err := GenerateBitcoinWalletFromSeed(seed)
	if err != nil {
		log.Fatal("Ошибка при генерации кошелька:", err)
	}

	// Сериализация данных в JSON
	serializedData, err := SerializeBitcoinWallet(wallet)
	if err != nil {
		log.Fatal("Ошибка при сериализации данных:", err)
	}

	// Сохранение сериализованных данных в файл
	err = SaveToFile(serializedData, "bitcoin_wallet.json")
	if err != nil {
		log.Fatal("Ошибка при сохранении данных в файл:", err)
	}

	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Приватный ключ:", wallet.PrivateKey)
	fmt.Println("Публичный ключ:", wallet.PublicKey)
	fmt.Println("Адрес:", wallet.Address)

	// Чтение данных из файла и десериализация
	readData, err := ReadFromFile("bitcoin_wallet.json")
	if err != nil {
		log.Fatal("Ошибка при чтении данных из файла:", err)
	}

	deserializedWallet, err := DeserializeBitcoinWallet(readData)
	if err != nil {
		log.Fatal("Ошибка при десериализации данных:", err)
	}

	fmt.Println("\nДесериализованный приватный ключ:", deserializedWallet.PrivateKey)
	fmt.Println("Десериализованный публичный ключ:", deserializedWallet.PublicKey)
	fmt.Println("Десериализованный адрес:", deserializedWallet.Address)

	fmt.Println(deserializedWallet.Address == wallet.Address)

}
