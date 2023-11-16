
<template>
  <div>
    <MainHeader :returnUrl="settingsUrl" />
    <div class="about-us">
            <h1>
                Delta it's Wallet
          </h1>
          <p>Version: <strong>0.01</strong></p>



      <p>
        Made with <span class="text-danger">❤️</span> by
        <a
          href="https://borusovillia.netlify.app/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Borusov Illia
        </a>
      </p>
    </div>
  </div>
</template>

<script>
import MainHeader from '../MainHeader.vue';

export default {
  name: 'AboutUS',
  components: {
    MainHeader,
  },
};
</script>
