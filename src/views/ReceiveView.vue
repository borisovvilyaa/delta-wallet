<template>
  <main>
    <div class="container">
      <div class="main-container">
        <MainHeader />
        <ReceiveContent :address="address" />
        <QrCodeReceive :address="address" />
      </div>
    </div>
  </main>
</template>

<script>
import MainHeader from '@/components/MainHeader.vue';
import ReceiveContent from '@/components/Receive//ReceiveContent.vue';
import QrCodeReceive from '@/components/Receive//QrCodeReceive.vue';
import WalletData from '../../keys/wallet.json';

export default {
  name: 'ReceiveView',
  components: {
    MainHeader,
    ReceiveContent,
    QrCodeReceive,
  },
  data() {
    return {
      address: WalletData.wallet.address,
    };
  },
}
</script>
