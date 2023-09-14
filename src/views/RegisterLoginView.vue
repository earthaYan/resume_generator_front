<script setup lang="ts">
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
const registerValues = ref({
  email: '',
  displayName: ''
})
const loginValues = ref({
  email: ''
})
const handleLoginClick = () => {
  loginForm.value?.validate((errors) => {
    if (!errors) {
      message.success('登录参数合法')
    } else {
      console.log(errors)
      message.error('登录参数不合法')
    }
  }).then(()=>{
  })
}
const handleRegisterClick = () => {
  registerForm.value?.validate((errors) => {
    if (!errors) {
      message.success('注册参数合法')
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
          <n-form-item label="邮箱" path="email">
            <n-input size="small" v-model:value="registerValues.email" round placeholder="邮箱" />
          </n-form-item>
          <n-form-item label="昵称" path="displayName">
            <n-input
              size="small"
              v-model:value="registerValues.displayName"
              round
              placeholder="昵称"
            />
          </n-form-item>
          <n-form-item>
            <n-button type="primary" block @click="handleRegisterClick"> 注册 </n-button>
          </n-form-item>
        </n-form>
      </n-tab-pane>
      <n-tab-pane name="login" tab="登录">
        <n-form ref="loginForm" :model="loginValues">
          <n-form-item label="邮箱" path="email">
            <n-input size="small" round placeholder="邮箱" v-model:value="loginValues.email" />
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
