<template>
  <header>
    <div class="container">
      <div class="d-flex align-items-center justify-content-between pt-4">
        <LogoHeader />
        <UserInfoHeader v-if="hasSession" />
      </div>
    </div>
  </header>
</template>

<script>
import LogoHeader from './Header/LogoHeader.vue';
import UserInfoHeader from './Header/UserInfoHeader.vue';

export default {
  name: 'AppHeader',
  components: {
    LogoHeader,
    UserInfoHeader,
  },
  data() {
    return {
      hasSession: sessionStorage.getItem('walletData') !== null,
    };
  },
};
</script>
