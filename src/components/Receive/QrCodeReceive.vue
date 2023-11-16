<template>
    <div class="qr-code-receive">
      <p>QRCode</p>
      <qrcode-vue :value="address"></qrcode-vue>
    </div>
  </template>
  
  <script>
  import QrcodeVue from 'vue-qrcode';
  
  export default {
    props: {
      address: String,
    },
    components: {
      QrcodeVue,
    },
  }
  </script>
  