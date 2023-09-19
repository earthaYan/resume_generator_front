<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import type { AuthValueType } from '@/types/auth'
import {
  NForm,
  NFormItem,
  NInput,
  NButton,
  NTabs,
  NTabPane,
  type FormInst,
  useMessage
} from 'naive-ui'
import { ref } from 'vue'
const message = useMessage()
const registerForm = ref<FormInst | null>(null)
const loginForm = ref<FormInst | null>(null)
const registerValues = ref<AuthValueType>({
  user_name: '',
  password: ''
})
const loginValues = ref<AuthValueType>({
  user_name: '',
  password: ''
})
const { login, register } = useAuthStore()
const handleLoginClick = () => {
  loginForm.value?.validate((errors) => {
    if (!errors) {
      message.success('登录参数合法')
      login(loginValues.value)
    } else {
      message.error('登录参数不合法')
    }
  })
}
const handleRegisterClick = () => {
  registerForm.value?.validate((errors) => {
    if (!errors) {
      message.success('注册参数合法')
      register(registerValues.value)
    } else {
      console.log(errors)
      message.error('注册参数不合法')
    }
  })
}
</script>
<template>
  <div class="register-login-wrapper">
    <n-tabs type="segment">
      <n-tab-pane name="register" tab="注册">
        <n-form ref="registerForm" :model="registerValues">
          <n-form-item label="用户名" path="user_name">
            <n-input size="small" v-model:value="registerValues.user_name" round />
          </n-form-item>
          <n-form-item label="密码" path="password">
            <n-input size="small" v-model:value="registerValues.password" round />
          </n-form-item>
          <n-form-item>
            <n-button type="primary" block @click="handleRegisterClick"> 注册 </n-button>
          </n-form-item>
        </n-form>
      </n-tab-pane>
      <n-tab-pane name="login" tab="登录">
        <n-form ref="loginForm" :model="loginValues">
          <n-form-item label="用户名" path="user_name">
            <n-input size="small" round v-model:value="loginValues.user_name" />
          </n-form-item>
          <n-form-item label="密码" path="password">
            <n-input size="small" round v-model:value="loginValues.password" />
          </n-form-item>
          <n-form-item>
            <n-button type="primary" block @click="handleLoginClick"> 登入 </n-button>
          </n-form-item>
        </n-form>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>
<style scoped lang="less"></style>
