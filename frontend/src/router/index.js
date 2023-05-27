import { createRouter, createWebHistory } from 'vue-router';
import HelpQueriesPage from '@/views/HelpQueriesPage.vue';
import AuthPage from '@/views/AuthPage.vue';
import RegisterPage from '@/views/RegisterPage.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      meta: { layout: 'main' },
      component: HelpQueriesPage,
    },
    {
      path: '/projects',
      name: 'projects',
      meta: { layout: 'main' },
      component: () => import('@/views/ProjectsPage.vue'),
    },
    {
      path: '/sign-in',
      name: 'signIn',
      meta: { layout: 'empty' },
      component: AuthPage,
    },
    {
      path: '/sign-up',
      name: 'signUp',
      meta: { layout: 'empty' },
      component: RegisterPage,
    },
  ],
});

export default router;
