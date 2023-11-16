<template>
  <main>
    <div class="container">
      <div class="main-container">
        <MainHeader />
        <RecipientContent />
        <SendForm />
      </div>
    </div>
  </main>
</template>

<script>
import MainHeader from '@/components/MainHeader.vue';
import RecipientContent from '@/components/Send/RecipientContent.vue';
import SendForm from '@/components/Send/SendForm.vue';

export default {
  name: 'SendView',
  components: {
    MainHeader,
    RecipientContent,
    SendForm,
  },
}
</script>
