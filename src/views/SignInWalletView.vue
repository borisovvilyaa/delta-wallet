<template>
    <main>
      <div class="container">
        <div class="main-container">
          <div class="container mt-5">
            <h1 class="mb-4">Your Bitcoin Wallet Mnemonic</h1>
            <div class="row">
              <form @submit.prevent="generateWallet" class="mb-3">
                <label class="form-label">Select the number of words:</label>
                <select v-model="selectedInputCount" class="form-select">
                  <option value="12">12 words</option>
                  <option value="24">24 words</option>
                </select>
                
                <div class="d-flex flex-wrap mt-3">
                  <div v-for="input in inputCount" :key="input" class="mb-2 col-md-4">
                    <input v-model="mnemonicWords[input - 1]" class="form-control spaced-input" type="text" :placeholder="'Word ' + input">
                  </div>
                </div>
                
                <button class="btn btn-gradient mt-2 w-100" type="submit">Generate Wallet</button>
              </form>
              <div class="alert alert-info mt-3" role="alert">
                <p class="mb-0">Make sure to keep your mnemonic phrase safe and private. Do not share it with anyone!</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </template>
  
  <script>
  export default {
    name: 'SignInWalletView',
    data() {
      return {
        selectedInputCount: 12,
        mnemonicWords: new Array(12).fill(''),
      };
    },
    computed: {
      inputCount() {
        return parseInt(this.selectedInputCount);
      },
    },
    methods: {
      generateWallet() {
        // Handle the generated wallet logic
        console.log(this.mnemonicWords);
      },
    },
  }
  </script>
  
  <style>
  /* Add your styles here */
  .spaced-input {
    margin-right: 10px; /* Adjust as needed */
  }
  </style>
  