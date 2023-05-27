import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/views/HomeView.vue';
import AuthPage from '@/views/AuthPage.vue';
import RegisterPage from '@/views/RegisterPage.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/sign-in',
      name: 'signIn',
      component: AuthPage,
    },
    {
      path: '/sign-up',
      name: 'signUp',
      component: RegisterPage,
    },
  ],
});

export default router;
