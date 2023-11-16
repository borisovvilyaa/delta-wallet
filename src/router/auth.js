export default function authMiddleware(to, from, next) {
    // Проверка наличия сессии и wallet.js
    const hasSession = sessionStorage.getItem('walletData');
    const hasWalletData = localStorage.getItem('walletData');
  
    if (!hasSession) {
      next('/login'); // Перенаправление на /login
    } else if (!hasWalletData) {
      next('/sign-in'); // Перенаправление на /sign-in
    } else {
      next(); // Продолжить навигацию
    }
  }