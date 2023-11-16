<template>
  <main>
    <div class="container">
      <div class="main-container">
        <div class="container mt-5">
          <h1 class="mb-4">Login to Your Bitcoin Wallet</h1>
          <div class="row">
            <form @submit.prevent="loginWallet" class="mb-3">
              <label class="form-label">Enter your password:</label>
              <input v-model="mnemonicPassword" class="form-control" type="password" placeholder="Enter your Password" />

              <button class="btn btn-gradient mt-2 w-100" type="submit">Login</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import walletData from "../../keys/wallet.json";
import CryptoJS from "crypto-js";

export default {
  name: 'LoginWalletView',
  data() {
    return {
      mnemonicPassword: '',
    };
  },
  created() {
    // Проверяем, была ли страница уже обновлена в текущей сессии
    const hasReloaded = localStorage.getItem('hasReloaded');

    if (!hasReloaded) {
      // Автоматически обновляем страницу только при первой загрузке
      localStorage.setItem('hasReloaded', 'true');
      location.reload();
    } else {
      // Сбрасываем флаг после обновления страницы, чтобы он сработал снова при следующей загрузке
      localStorage.removeItem('hasReloaded');
    }
  },
  methods: {
    loginWallet() {
      const enteredPassword = this.mnemonicPassword;
      const hashedEnteredPassword = CryptoJS.SHA256(enteredPassword).toString(CryptoJS.enc.Hex);
      const storedPassword = walletData.wallet.credentials.password;

      if (hashedEnteredPassword === storedPassword) {
        // Пароль верный, выполняйте логин
        console.log('Logged in successfully');

        sessionStorage.setItem('walletData', JSON.stringify(walletData));

        // Устанавливаем время жизни сессии на 1 час (3600 секунд)
        sessionStorage.setItem('sessionTimeout', 3600);

        // Force a page refresh by navigating to the current route
        this.$router.push({ path: '/' }).then(() => {
          this.$router.go(0);
        });
      } else {
        // Пароль неверный, выводите сообщение об ошибке
        console.error('Incorrect password. Please try again.');
      }
    },
  },
};
</script>



