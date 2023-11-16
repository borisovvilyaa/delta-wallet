<template>
  <div class="main-header mb-2">
    <div class="d-flex justify-content-between">
      <router-link :to="returnUrl" ><i class="bi bi-arrow-left"></i> back</router-link>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    returnUrl: {
      type: String,
      default: '/', // Значение по умолчанию, если параметр returnUrl не передан
    },
  },
};
</script>
