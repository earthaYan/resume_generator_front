import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { instance } from '@/utils/api'
import type { AuthValueType } from '@/types/auth'
export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const isLogin = ref(false)
  async function login(loginInfo: AuthValueType) {
    try {
      const res = await instance.post('/user/login', {
        ...loginInfo
      })
      localStorage.setItem('token', res.data.data.access_token)
      isLogin.value = true
      router.push({
        path: '/resumes'
      })
    } catch (err) {
      console.log(err)
    }
  }
  function logout() {
    isLogin.value = false
    localStorage.removeItem('token')
    router.push({
      path: '/login'
    })
  }
  async function register(registerInfo: AuthValueType) {
    try {
      await instance.post('/user/register', {
        ...registerInfo
      })
      router.push({
        path: '/login'
      })
    } catch (err) {
      console.log(err)
    }
  }
  function checkLoginStatus() {
    const token = localStorage.getItem('token')
    if (token) {
      console.log('login')

      isLogin.value = true
    }
  }
  return {
    isLogin,
    login,
    logout,
    register,
    checkLoginStatus
  }
})
