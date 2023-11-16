
<template>
    <div class="activity-send">
      <p>
        <span class="c-blue" :data-copy="transaction.hash">{{ shortenedAddress(transaction.hash) }}</span>
        <a :href="'https://www.blockchain.com/btc/tx/' + transaction.hash" target="_blank">
          <i class="bi bi-arrow-up-right"></i>
        </a>
      </p>
      <div class="d-flex justify-content-between">
        <div>
          <p>
            <span class="font-monospace"> {{ isSent(transaction) ? 'Sent' : 'Received' }}</span>
            to
            <span class="c-blue" :data-copy="getAddress(transaction)">{{ shortenedAddress(getAddress(transaction)) }}</span>
          </p>
          <p class="date-transaction">{{ formatDate(transaction.time) }}</p>
        </div>
        <div>
          <p :class="isSent(transaction) ? 'c-red' : 'c-green'" class="amount-transaction">
            {{ (isSent(transaction) ? '-' : '+') }}
            <span :data-copy="Math.abs(transaction.out[0].value / 1e8)">
              {{ Math.abs(transaction.out[0].value / 1e8) }} BTC
            </span>
          </p>
        </div>
      </div>
      <hr>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      transaction: Object,
    },
    methods: {
      formatDate(timestamp) {
        const date = new Date(timestamp * 1000);
        return date.toLocaleDateString();
      },
      isSent(transaction) {
        return transaction.inputs[0].prev_out.addr === '165bF16Estyy6J4pHPgdRqtE9d1rKdzcaK';
      },
      getAddress(transaction) {
        return this.isSent(transaction) ? transaction.out[0].addr : transaction.inputs[0].prev_out.addr;
      },
      shortenedAddress(fullAddress) {
        const lengthToShow = 10;
  
        if (fullAddress.length <= lengthToShow) {
          return fullAddress;
        } else {
          const beginning = fullAddress.substring(0, lengthToShow / 2);
          const ending = fullAddress.substring(fullAddress.length - lengthToShow / 2);
          return `${beginning}...${ending}`;
        }
      },
    },
  };
  </script>

  