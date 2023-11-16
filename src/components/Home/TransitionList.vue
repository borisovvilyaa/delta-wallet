<template>
  <div>
    <div class="mt-3">
      <label for="transactionType">Select Transaction Type:</label>
      <select id="transactionType" v-model="transactionType" class="list-select">
        <option value="api" active>API Transaction</option>
        <option value="local">Local Transaction</option>
      </select>
    </div>

    <div v-if="transactionType === 'api'">
      <div v-if="isNoTransactions" class="mt-3">
        <h3 style="color: #ccc">No operations found <i class="bi bi-search-heart"></i></h3>
      </div>
      <div v-else>
        <div class="mt-3">
          <div class="d-flex justify-content-between">
            <label for="transactionsPerPage">Transactions per page: </label>
            <select id="transactionsPerPage" v-model="transactionsPerPage" @change="updatePage(1)" class="list-select">
              <option v-for="option in transactionOptions" :value="option" :key="option">{{ option }}</option>
            </select>
          </div>
          <hr>

          <div v-for="(transaction, index) in displayedTransactions" :key="index">
            <TransactionItem :transaction="transaction" />
          </div>

          <div class="pagination d-flex align-items-center justify-content-center gap-2">
            <button :disabled="currentPage === 1" @click="changePage(currentPage - 1)" class="btn btn-gradient">Previous</button>
            <span class="pagination-info">Page {{ currentPage }} of {{ totalPages }}</span>
            <button :disabled="currentPage === totalPages" @click="changePage(currentPage + 1)" class="btn btn-gradient">Next</button>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="transactionType === 'local'">
      <!-- Ваш код для локальных транзакций здесь -->
    </div>
  </div>
</template>

<script>
import WalletData from '../../../keys/wallet.json';
import TransactionItem from './TransactionItem.vue';

export default {
  name: 'TransitionList',
  components: {
    TransactionItem
  },
  data() {
    return {
      transactions: [],
      transactionsPerPage: 5,
      currentPage: 1,
      transactionType: 'api'
    };
  },
  computed: {
    transactionOptions() {
      return [2, 5, 10, 30, 100];
    },
    totalPages() {
      return Math.ceil(this.transactions.length / this.transactionsPerPage);
    },
    displayedTransactions() {
      const startIndex = (this.currentPage - 1) * this.transactionsPerPage;
      const endIndex = startIndex + this.transactionsPerPage;
      return this.transactions.slice(startIndex, endIndex);
    },
    isNoTransactions() {
      return this.transactions.length === 0;
    }
  },
  watch: {
    transactionsPerPage() {
      this.updatePage(1);
    }
  },
  methods: {
    async fetchTransactions() {
      const address = WalletData.wallet.address;
      try {
        if (this.transactionType === 'api') {
          const response = await fetch(`https://blockchain.info/rawaddr/${address}`);
          const data = await response.json();

          if (data.txs && data.txs.length > 0) {
            this.transactions = data.txs;
          } else {
            this.transactions = [];
          }
        } else if (this.transactionType === 'local') {
          // Ваш код для локальных транзакций
        }
      } catch (error) {
        console.error('Error fetching transactions:', error);
      }
    },
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
    updatePage(page) {
      this.currentPage = page;
    },
    changePage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
      }
    }
  },
  mounted() {
    this.fetchTransactions();
  }
};
</script>
