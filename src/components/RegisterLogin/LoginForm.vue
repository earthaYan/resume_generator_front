<script setup lang="ts">
import { ref } from 'vue'
import { NForm, NFormItem, NInput, NButton, type FormInst } from 'naive-ui'
const loginFormRef = ref<FormInst | null>(null)
const formValue = ref({
  name: '',
  email: '',
  password: ''
})
</script>
<template>
  <n-form ref="loginFormRef" :model="formValue">
    <h2>登入账号</h2>
    <n-form-item path="email">
      <n-input placeholder="邮箱" />
    </n-form-item>
    <n-form-item path="password">
      <n-input placeholder="密码" />
    </n-form-item>
    <n-form-item>
      <n-button type="info" round size="small"> 登入 </n-button>
    </n-form-item>
  </n-form>
</template>
