<script setup lang="ts">
import LoginForm from '@/components/RegisterLogin/LoginForm.vue'
import RegisterForm from '@/components/RegisterLogin/RegisterForm.vue'
</script>
<template>
  <div class="register-login-wrapper">
    <LoginForm />
    <RegisterForm />
    <div class="switch">
      <div class="switch-circle"></div>
      <div class="switch_circle switch_circle-t"></div>
      <!-- 注册左侧 -->
      <div class="switch_container" id="switch-c1">
        <h2 class="switch_title title">欢迎回来！</h2>
        <p class="switch_description description">
          已经有账号了嘛，去登入账号来进入奇妙世界吧！！！
        </p>
        <button class="switch_button button switch-btn">SIGN IN</button>
      </div>
      <!-- 登录右侧 -->
      <div class="switch_container is-hidden" id="switch-c2">
        <h2 class="switch_title title" style="letter-spacing: 0">Hello Friend！</h2>
        <p class="switch_description description">
          去注册一个账号，成为尊贵的粉丝会员，让我们踏入奇妙的旅途！
        </p>
        <button class="switch_button button switch-btn">SIGN UP</button>
      </div>
    </div>
  </div>
</template>
<style scoped lang="less"></style>
