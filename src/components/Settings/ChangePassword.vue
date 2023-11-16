<template>
  <div>
    <MainHeader :returnUrl="settingsUrl" />
    <div class="container">
      <div class="row">
        <div>
          <h2>Change Password</h2>
          <form @submit.prevent="changePassword">
            <div class="form-group">
              <label for="currentPassword">Current Password:</label>
              <input type="password" class="form-control" id="currentPassword" v-model="currentPassword" required />
            </div>
            
            <div class="form-group mt-2">
              <label for="newPassword">New Password:</label>
              <input type="password" class="form-control" id="newPassword" v-model="newPassword" required />
            </div>
            
            <div class="form-group mt-2">
              <label for="confirmPassword">Confirm New Password:</label>
              <input type="password" class="form-control" id="confirmPassword" v-model="confirmPassword" required />
            </div>
            
            <button type="submit" class="btn btn-gradient mt-2 w-100">Change Password</button>
          </form>
          <div class="text-danger" v-if="errorMessage">{{ errorMessage }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import MainHeader from '../MainHeader.vue';

export default {
  name: 'ChangePassword',
  components: {
    MainHeader,
  },
  data() {
    return {
      currentPassword: '',
      newPassword: '',
      confirmPassword: '',
      errorMessage: '', // Error message if password change fails
    };
  },
  computed: {
    settingsUrl() {
      // Define the URL to return to settings
      return '/settings'; // Replace with the actual URL for settings
    },
  },
  methods: {
    changePassword() {
      // Here, you can add logic to change the password
      // Check if the current password matches the user's password
      if (this.currentPassword !== 'current_password') {
        this.errorMessage = 'Current password is incorrect';
        return;
      }
      
      // Check if the new password and confirmation match
      if (this.newPassword !== this.confirmPassword) {
        this.errorMessage = 'New password and confirmation do not match';
        return;
      }
      
      // Perform password change here, e.g., send a request to the server
      // If the password change is successful, redirect the user to the settings page
      // If there's an error, set this.errorMessage with the error message
    },
  },
};
</script>
