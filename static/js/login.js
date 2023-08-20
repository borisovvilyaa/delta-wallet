document.getElementById('generateWallet').addEventListener('click', function() {
    const words = [];
    for (let i = 1; i <= 12; i++) {
      const wordInput = document.getElementById('word' + i);
      words.push(wordInput.value);
    }
    const mnemonic = words.join(' ');
  
    // Display the generated mnemonic phrase
    document.getElementById('mnemonicPhrase').textContent = mnemonic;
    document.getElementById('walletInfo').classList.remove('d-none');
  });