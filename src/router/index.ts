import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterLoginView from '../views/RegisterLoginView.vue'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    meta: { requiresAuth: true } 
  },
  {
    path: '/login',
    name: 'Login',
    component: RegisterLoginView
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterLoginView
  },
  {
    path: '/list',
    name: '简历列表',
    component: () => import('../views/AboutView.vue'),
  },
  {
    path: '/add',
    name: '添加简历',
    component: () => import('../views/GeneratorView.vue'),
  }
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})
export default router
