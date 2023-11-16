<template>
  <div class="pt-3 d-flex justify-content-between align-item-center">
    <router-link to="/receive" class="icon-menu">
      <div>
        <i class="bi bi-arrow-down-circle"></i>
      </div>
      <p>Receve</p>
    </router-link>
    <router-link to="/buy" class="icon-menu">
      <div>
        <i class="bi bi-bag"></i>
      </div>
      <p>Buy</p>
    </router-link>
    <router-link to="/withdraw" class="icon-menu">
      <div>
        <i class="bi bi-box-arrow-in-down"></i>
      </div>
      <p>Withdraw</p>
    </router-link>
    <router-link to="/send" class="icon-menu">
      <div>
        <i class="bi bi-arrow-up-circle"></i>
      </div>
      <p>Send</p>
    </router-link>
  </div>
</template>
<script>
export default {
  name: 'HomeMenu'
}
</script>
