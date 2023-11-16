<template>
  <div>
    <MainHeader :returnUrl="settingsUrl" />
    <div class="container">
      <div class="row justify-content-center">
        <div class="">
          <h2>Private Key</h2>
          <div class="alert alert-warning" role="alert">
            <strong>Warning:</strong> Displaying your private key is highly sensitive and should be done with caution.
          </div>
          <div class="form-group">
            <label for="privateKey">Your Private Key:</label>
            <textarea class="form-control" id="privateKey" v-model="privateKey" readonly></textarea>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import MainHeader from '../MainHeader.vue';

export default {
  name: 'ShowPrivateKey',
  components: {
    MainHeader,
  },
  data() {
    return {
      privateKey: 'Your_Private_Key_Goes_Here', // Replace with the actual private key
    };
  },
  computed: {
    settingsUrl() {
      // Define the URL to return to settings
      return '/settings'; // Replace with the actual URL for settings
    },
  },
  methods: {
    copyPrivateKey() {
      // You can add logic to copy the private key to the clipboard here
      // For example, you can use the Clipboard API or a third-party library
      // Be sure to handle any error messages or notifications as needed
    },
  },
};
</script>
