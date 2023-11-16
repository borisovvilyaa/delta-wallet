<template>
    <div class="receive-content">
      <h2>Address</h2>
      <p>Your current account address. Share it to receive funds.</p>
      <span id="copy" :data-copy="address">{{ address }}</span>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      address: String,
    },
  }
  </script>
  