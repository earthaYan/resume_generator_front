import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const isLogin = ref(false)
  function login() {
    isLogin.value = true
    router.push({
      path: '/list'
    })
  }
  function logout() {
    isLogin.value = false
    router.push({
      path: '/login'
    })
  }
  return {
    isLogin,
    login,
    logout
  }
})
