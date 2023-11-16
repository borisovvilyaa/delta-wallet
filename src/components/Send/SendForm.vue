<template>
    <form class="send-form" @submit.prevent="sendFunds">
      <input v-model="recipientAddress" type="text" placeholder="e.g. 1FfmbHfnpa...455paPh">
      <div class="receive-content mt-3">
        <h2>Amount</h2>
      </div>
      <input v-model="amount" type="text" placeholder="e.g. 123.45">
      <div v-if="showError" class="error-message">Invalid input. Please check your data.</div>
      <button type="submit">Send</button>
    </form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipientAddress: '',
        amount: '',
        showError: false,
      };
    },
    methods: {
      sendFunds() {
        // Пример простой валидации
        if (this.recipientAddress && this.amount && !isNaN(this.amount)) {
          this.showError = false;
          // Ваш код для отправки средств
        } else {
          this.showError = true;
        }
      },
    },
  }
  </script>
  
  <style>
  .error-message {
    color: red;
    margin-top: 5px;
  }
  </style>
  