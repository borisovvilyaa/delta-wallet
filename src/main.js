import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/css/main.css'; // Подключение глобальных стилей
import './assets/js/main.js'; 

createApp(App).use(store).use(router).mount('#app')
