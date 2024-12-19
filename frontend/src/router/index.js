import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/LoginView.vue';
import RegisterView from '@/views/RegisterView.vue';

const routes = [
  // redirect to login view
  { path: '/', redirect: '/login' },

  {
    path: '/login',
    name: 'Login',
    component: LoginView,
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterView,
  },
  // 其他路由配置
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;