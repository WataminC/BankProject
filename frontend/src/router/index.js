import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/LoginView.vue';
import RegisterView from '@/views/RegisterView.vue';
import UserDashboard from '@/views/DashboardView.vue';
import TransactionView from '@/views/TransactionView.vue';
import TransferView from '@/views/TransferView.vue';
import LoanView from '@/views/LoanView.vue';
import AdminView from '@/views/AdminView.vue';

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
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: UserDashboard,
    meta: { requiresAuth: true },
  },
  {
    path: '/transaction',
    name: 'Transaction',
    component: TransactionView,
    meta: { requiresAuth: true },
  },
  {
    path: '/transfer',
    name: 'Transfer',
    component: TransferView,
    meta: { requiresAuth: true },
  },
  {
    path: '/loan',
    name: 'Loan',
    component: LoanView,
    meta: { requiresAuth: true },
  },
  {
    path: '/admin',
    name: 'Admin',
    component: AdminView,
    meta: { requiresAuth: true },
  },
  // 其他路由配置
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const token = localStorage.getItem('token');

  if (requiresAuth && !token) {
    next('/login');
    alert('You must be logged in to access this page.');
  } else {
    next();
  }
});

export default router;