import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const isLoggedIn = ref(false)
  const registeredCredentials = ref([])
  const checkUserLoggedIn = () => {
    // 检查浏览器是否支持 WebAuthn
    if (!window.PublicKeyCredential) {
      return false
    }
    // 检查用户是否已经注册过公钥凭证
    const storedCredentials = JSON.parse(localStorage.getItem('registeredCredentials') || '[]')
    registeredCredentials.value = storedCredentials
    // 检查当前浏览器中是否存在有效的公钥凭证
    const isValidKey = registeredCredentials.value.some((credential: { valid: boolean }) => {
      // 在这里根据你的判断逻辑检查凭证的有效性，比如判断过期时间等
      return credential.valid
    })
    isLoggedIn.value = isValidKey
  }
  return { isLoggedIn, checkUserLoggedIn }
})
