import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterLoginView from '../views/RegisterLoginView.vue'
import { useAuthStore } from '../stores/auth'
type RouteItem = RouteRecordRaw & {
  key: string
  no_show_menu?: Boolean
}
export const routes: RouteItem[] = [
  {
    path: '/',
    name: 'Home',
    key: 'Home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'Login',
    key: 'Login',
    component: RegisterLoginView,
    meta: { requiresAuth: false }
  },
  {
    path: '/resumes',
    name: '简历列表',
    key: 'resume_list',
    component: () => import('../views/ListResumes.vue')
  },
  {
    path: '/add',
    name: '添加简历',
    key: 'resume_add',
    component: () => import('../views/GeneratorView.vue'),
    no_show_menu: true
  },
  {
    path: '/update/:resume_id',
    name: '修改简历',
    key: 'resume_update',
    component: () => import('../views/GeneratorView.vue'),
    no_show_menu: true
  }
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  authStore.checkLoginStatus()
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
