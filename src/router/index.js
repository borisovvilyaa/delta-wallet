import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ReceiveView from '../views/ReceiveView.vue'
import SendView from '../views/SendView.vue'
import SignInWalletView from '../views/SignInWalletView.vue'
import LoginWalletViewVue from '@/views/LoginWalletView.vue'

import SettingsView from '../views/SettingsView.vue'; 
import DeleteAddress from '../components/Settings/DeleteAddress.vue';
import ShowPrivateKey from '../components/Settings/ShowPrivateKey.vue';
import ChangePassword from '../components/Settings/ChangePassword.vue';
import AboutUsVue from '@/components/Settings/AboutUs.vue'


const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/receive',
    name: 'receive',
    component: ReceiveView
  },
  {
    path: '/send',
    name: 'send',
    component: SendView
  },
  {
    path: '/sign-in',
    name: 'signin',
    component: SignInWalletView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginWalletViewVue
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView,
    children: [
      { path: 'delete-address', name: 'delete-address', component: DeleteAddress },
      { path: 'show-private-key', name: 'show-private-key', component: ShowPrivateKey },
      { path: 'change-password', name: 'change-password', component: ChangePassword },
      { path: 'about-us', name: 'AboutUs', component: AboutUsVue },

    ]
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});
router.beforeEach((to, from, next) => {
  const hasSession = sessionStorage.getItem('walletData');
  const hasWalletData = localStorage.getItem('walletData');

  if (to.path === '/login') {
    console.log('Removing sessionStorage data');
    sessionStorage.removeItem('sessionTimeout');
    sessionStorage.removeItem('walletData');
    next();
  } else if (!hasSession) {
    // Redirect to /login if there is no session
    next('/login');
  } else if (!hasWalletData) {
    // Redirect to /sign-in if there is no walletData
    next('/sign-in');
  } else {
    // Allow access to other routes if session and walletData exist
    next();
  }
});


export default router;
