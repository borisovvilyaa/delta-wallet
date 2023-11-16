<template>
  <main>
    <div class="container">
      <div class="main-container">
        <!-- General Settings -->
        <div v-if="$route.path === '/settings'" class="row align-items-stretch">
          <div class="col-md-6 mb-3">
            <div class="p-3 border rounded shadow-sm d-flex flex-column align-items-center h-100">
              <i class="bi bi-gear fs-3 mb-2"></i> <!-- Увеличение размера иконки -->
              <div>
                <router-link to="/sign-in" class="btn btn-link w-100">Import Wallet Address</router-link>
                <p class="text-muted small">Import your wallet address for easy access.</p>
              </div>
            </div>
          </div>

          <div class="col-md-6 mb-3">
            <div class="p-3 border rounded shadow-sm d-flex flex-column align-items-center h-100">
              <i class="bi bi-box-arrow-left fs-3 mb-2"></i> <!-- Увеличение размера иконки -->
              <div>
                <router-link to="/settings/delete-address" class="btn btn-link w-100">Exit Wallet Address</router-link>
                <p class="text-muted small">Exit from your wallet address for security reasons.</p>
              </div>
            </div>
          </div>

          <div class="col-md-6 mb-3">
            <div class="p-3 border rounded shadow-sm d-flex flex-column align-items-center h-100">
              <i class="bi bi-key fs-3 mb-2"></i> <!-- Увеличение размера иконки -->
              <div>
                <router-link to="/settings/show-private-key" class="btn btn-link w-100">Show PK</router-link>
                <p class="text-muted small">Display your private key securely.</p>
              </div>
            </div>
          </div>

          <div class="col-md-6 mb-3">
            <div class="p-3 border rounded shadow-sm d-flex flex-column align-items-center h-100">
              <i class="bi bi-lock fs-3 mb-2"></i> <!-- Увеличение размера иконки -->
              <div>
                <router-link to="/settings/change-password" class="btn btn-link w-100">Change Password</router-link>
                <p class="text-muted small">Change your account password.</p>
              </div>
            </div>
          </div>

          <!-- Новый пункт "About Us" -->
          <div class="col-md-6 mb-3">
            <div class="p-3 border rounded shadow-sm d-flex flex-column align-items-center h-100">
              <i class="bi bi-info-circle fs-3 mb-2"></i> <!-- Иконка "About Us" -->
              <div>
                <router-link to="/settings/about-us" class="btn btn-link w-100">About Us</router-link>
                <p class="text-muted small">Learn more about our organization.</p>
              </div>
            </div>
          </div>
          <!-- Конец нового пункта "About Us" -->

        </div>
        <!-- End of General Settings -->
        <router-view></router-view>
      </div>
    </div>
  </main>
</template>




<script>
export default {
  name: 'SettingsView',
};
</script>
