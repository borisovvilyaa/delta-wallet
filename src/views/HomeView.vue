<template>
  <main>
    <div class="container">
      <div class="main-container">
        <div class="name-wallet" v-if="!isEditingName" @click="toggleNameEditing">
          {{ walletName }}
        </div>
        <div v-else class="d-flex gap-2">
          <input type="text" v-model="walletName" class="form-control">
          <button @click="saveWalletName" class="btn btn-primary">Save</button>
        </div>
        <hr class="mb-3">

        <div class="d-flex justify-content-between">
          <div>
            <p>Total Equity Value</p>
            <p class="amount mt-3">≈ <span :data-copy="bitcoinBalance" id="copy">{{ bitcoinBalance }}</span> BTC</p>
            <p :data-copy="bitcoinPrice * bitcoinBalance" id="copy">{{ bitcoinPrice * bitcoinBalance}} USD</p> 
          </div>
          <div>
            <p id="copy" class="address" :data-copy="walletAddress">{{ shortenedAddress }}</p>
          </div>
        </div>
        <HomeMenu/>
        <div class="activity pt-5">
          <h3>Transaction history</h3>
          <TransactionList/>
        </div>
      </div>
    </div>
  </main>
</template>
<script>
import TransactionList from '../components/Home/TransitionList.vue';
import HomeMenu from '@/components/Home/HomeMenu.vue';

export default {
  name: 'HomeView',
  components: {
    TransactionList,
    HomeMenu
  },
  data() {
    const savedWalletData = sessionStorage.getItem('walletData'); // Используем sessionStorage вместо localStorage
    const initialWalletData = savedWalletData ? JSON.parse(savedWalletData) : null;

    console.log('Initial wallet data:', initialWalletData);

    return {
      isEditingName: false,
      walletName: initialWalletData ? initialWalletData.wallet.name : '',
      editedWalletName: initialWalletData ? initialWalletData.wallet.name : '',
      walletAddress: initialWalletData ? initialWalletData.wallet.address : '',
      bitcoinBalance: 0,
      bitcoinPrice: 0,
    };
  },

  created() {
    if (this.walletAddress) {
      this.getBitcoinBalance();
      this.getBitcoinPrice();
    }
  },
  computed: {
    shortenedAddress() {
      const fullAddress = this.walletAddress;
      const lengthToShow = 10;

      if (fullAddress.length <= lengthToShow) {
        return fullAddress;
      } else {
        const beginning = fullAddress.substring(0, lengthToShow / 2);
        const ending = fullAddress.substring(fullAddress.length - lengthToShow / 2);
        return `${beginning}...${ending}`;
      }
    }
  },
  methods: {
    async getBitcoinPrice() {
      try {
        const response = await fetch("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd");
        const data = await response.json();
        this.bitcoinPrice = data.bitcoin.usd;
      } catch (error) {
        console.error("Error fetching Bitcoin price:", error);
      }
    },
    async getBitcoinBalance() {
      try {
        const response = await fetch(`https://blockchain.info/q/addressbalance/${this.walletAddress}`);
        const balanceInSatoshi = await response.text();
        this.bitcoinBalance = balanceInSatoshi / 10 ** 8;

      } catch (error) {
        console.error("Error fetching Bitcoin balance:", error);
      }
    },
    toggleNameEditing() {
      this.isEditingName = !this.isEditingName;
    },
    saveWalletName() {
      this.walletName = this.editedWalletName;
      this.isEditingName = false;

      const walletData = {
        wallet: {
          name: this.walletName,
          address: this.walletAddress,
        }
      };

      sessionStorage.setItem('walletData', JSON.stringify(walletData));
      console.log('Wallet data saved successfully.');
    }
  }
};
</script>



<style>
/* ... */
</style>
