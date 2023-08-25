package services

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Valide")

	// Process the form data
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		var inputs []string

		// Loop through the form values and collect them
		for i := 0; i <= 11; i++ {
			inputName := fmt.Sprintf("word%d", i)
			inputValue := r.FormValue(inputName)
			inputs = append(inputs, fmt.Sprintf("%s ", inputValue))
		}

		// Join the collected inputs and print them in a single line
		// Join the collected inputs and print them in a single line
		mnemonic := strings.Join(inputs, "")

		// Validate the mnemonic
		if !bip39.IsMnemonicValid(mnemonic) {
			http.Error(w, "Invalid mnemonic", http.StatusBadRequest)
			return
		}
		// 	mnemonic := "afford worth interest casual level service virus infant penalty crystal cave black"
		seed := bip39.NewSeed(mnemonic, "delta")

		wallet, err := GenerateBitcoinWalletFromSeed(seed)
		if err != nil {
			log.Fatal("Ошибка при генерации кошелька:", err)
		}

		fmt.Println("Mnemonic:", mnemonic)
		fmt.Println("Приватный ключ:", wallet.PrivateKey.ToECDSA().D.Text(16))
		fmt.Println("Публичный ключ:", wallet.PublicKey)
		fmt.Println("Адрес:", wallet.Address)

		if err := SavePrivateKeyToFile("seed.seco", wallet.PrivateKey); err != nil {
			log.Fatal("Ошибка при сохранении приватного ключа:", err)
		}

		if err := SavePassphraseToFile("passphrase.json", mnemonic, "delta"); err != nil {
			log.Fatal("Ошибка при сохранении JSON в файл:", err)
		}

		readPrivateKey, err := ReadPrivateKeyFromFile("seed.seco")
		if err != nil {
			log.Fatal("Ошибка при чтении приватного ключа из файла:", err)
		}
		fmt.Println("\nприватный ключ:", readPrivateKey.ToECDSA().D.Text(16))

		readMnemonic, readPassphrase, err := ReadPassphraseFromFile("passphrase.json")
		if err != nil {
			log.Fatal("Ошибка при чтении JSON из файла:", err)
		}
		fmt.Println("Прочитанный mnemonic:", readMnemonic)
		fmt.Println("Прочитанный passphrase:", readPassphrase)
		// Save session data
		err = CreateCryptoSession(w, r, "address", wallet.Address)
		if err != nil {
			log.Fatal("Error creating session:", err)
		}

		// Redirect to /home
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}
}
