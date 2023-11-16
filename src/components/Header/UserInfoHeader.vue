<template>
    <div class="user-info d-flex align-items-center gap-3">
      <div class="user-details ml-2">
        <p class="user-name">Account 1</p>
        <div class="network-dropdown">
          <select class="network-select">
            <option value="mainnet">Bitcoin Mainnet</option>
            <option value="cash">Bitcoin Cash</option>
          </select>
        </div>
      </div>
      <AvatarDropdown />
    </div>
  </template>
  
  <script>
  import AvatarDropdown from './AvatarDropdown.vue';
  
  export default {
    name: 'UserInfoHeader',
    components: {
      AvatarDropdown,
    },
  }
  </script>
  