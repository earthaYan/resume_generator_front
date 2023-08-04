<script setup lang="ts">
import { ref } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
const registerFormRef = ref(null)
const formValue = ref({
  name: '',
  email: '',
  password: ''
})
</script>
<template>
  <h2>创建账号</h2>
  <n-form ref="registerFormRef" :model="formValue">
    <n-form-item path="name">
      <n-input placeholder="用户名" />
    </n-form-item>
    <n-form-item path="email">
      <n-input placeholder="邮箱" />
    </n-form-item>
    <n-form-item path="password">
      <n-input placeholder="密码" />
    </n-form-item>
    <n-form-item>
      <n-button type="info" round size="small"> 注册 </n-button>
    </n-form-item>
  </n-form>
</template>
