import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { instance } from '@/utils/api'
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
  async function handleWebauthnLogin(loginValues: any) {
    try {
      const publicKey = await instance.get('/api/login') // 发起后端API请求获取公钥

      const credential = await navigator.credentials.get({
        publicKey: publicKey.data // 使用公钥调用WebAuthn API
      })
      // 将生成的凭证发送到服务器进行验证
      const response = await instance.post('/api/login', {
        credential: credential // 发送生成的凭证到服务器进行验证
      })
      if (response.data.success) {
        // 登录成功后的逻辑
        console.log('登录成功！')
        login()
      } else {
        // 登录失败的处理
        console.error('登录失败：' + response.data.message)
      }
    } catch (error) {
      console.error('登录失败', error)
    }
  }
  async function handleWebauthnRegister(registerValues: any) {
    try {
      const publicKey = await instance.get('/api/register') // 发起后端API请求获取公钥
      const credential = await navigator.credentials.create({
        publicKey: publicKey.data // 使用公钥调用WebAuthn API
      })
      // 处理登录成功后的逻辑，发送凭证到服务器进行验证等操作
      const response = await instance.post('/api/register', {
        credential: credential // 发送生成的凭证到服务器进行验证
      })
      if (response.data.success) {
        // 注册成功后的逻辑
        console.log('注册成功！')
      } else {
        // 注册失败的处理
        console.error('注册失败：' + response.data.message)
      }
    } catch (error) {
      console.error('注册失败', error)
    }
  }
  return {
    isLogin,
    login,
    logout,
    handleWebauthnLogin,
    handleWebauthnRegister
  }
})
