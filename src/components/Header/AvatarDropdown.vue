<template>
    <div class="avatar dropdown" @click="toggleDropdown">
      <div class="dropdown-toggle" role="button" :class="{ active: isDropdownOpen }" aria-haspopup="true" aria-expanded="false">
        <img src="@/assets/img/avatar/avatar_small.jpg" alt="avatar" class="user-avatar shadow">
      </div>
      <div class="dropdown-menu" :class="{ show: isDropdownOpen }" aria-labelledby="avatarDropdown">
        <router-link to="/login" class="dropdown-item">
          <i class="bi bi-lock"></i> Login
        </router-link>
        <router-link to="/sign-in" class="dropdown-item">
          <i class="bi bi-box-arrow-in-down"></i> Import Account

        </router-link>
        <router-link to="/settings" class="dropdown-item">
          <i class="bi bi-gear"></i> Settings
        </router-link>
      </div>
    </div>

  </template>
  <style scoped>
  .avatar {
    position: relative;
    display: inline-block;
  }
  
  .dropdown-toggle {
    cursor: pointer;
  }
  
  
  .dropdown-menu {
    position: absolute;
    top: 100%;
    right: 0;
    display: none;
    background-color: #ffffff;
    box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
    border-radius: 4px;
    padding: 8px;
  }
  
  .dropdown-menu.show {
    display: block;
  }
  .lock-screen {
    
  position: fixed;
  top: 0;
  left: 0;
  z-index: 999;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.lock-box {
  background-color: white;
  padding: 20px;
  border-radius: 4px;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  align-items: center;
}

.lock-box input {
  margin-bottom: 10px;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.lock-box button {
  padding: 5px 10px;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.lock-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(8px); /* Добавляем блюр на фон */
  z-index: 1;
}

  </style>
<script>
export default {
  name: 'AvatarDropdown',
  data() {
    return {
      isDropdownOpen: false,
      isLocked: false,
      password: '1'
    };
  },
  methods: {
    toggleDropdown() {
      this.isDropdownOpen = !this.isDropdownOpen;
    },
    reloadLoginPage() {
      window.location.reload();
    }
  }
};
</script>
