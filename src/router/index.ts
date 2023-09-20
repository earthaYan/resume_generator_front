import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterLoginView from '../views/RegisterLoginView.vue'
import { useAuthStore } from '../stores/auth'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'Login',
    component: RegisterLoginView,
    meta: { requiresAuth: false }
  },
  {
    path: '/resumes',
    name: '简历列表',
    component: () => import('../views/ListResumes.vue')
  },
  {
    path: '/add',
    name: '添加简历',
    component: () => import('../views/GeneratorView.vue')
  }
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (!authStore.isLogin && to.path !== '/login') {
    // 未登录且不是访问登录页，则重定向到登录页
    next({ path: '/login' })
  } else if (authStore.isLogin && to.path === '/login') {
    // 已登录且正在访问登录页，则重定向到列表页
    next({ path: '/resumes' })
  } else {
    // 登录状态和目标路由符合预期，则继续导航
    next()
  }
})
export default router
