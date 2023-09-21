<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { NMessageProvider, NButton, NSpace } from 'naive-ui'
import { routes } from './router'
import { computed } from 'vue'
import { useAuthStore } from './stores/auth'
import { storeToRefs } from 'pinia'
const authStore = useAuthStore()
const { isLogin } = storeToRefs(authStore)
const { logout } = authStore
const filteredRoutes = computed(() => {
  if (isLogin.value) {
    return routes.filter((route) => {
      return route.name !== 'Login' && !route.no_show_menu
    })
  } else {
    return []
  }
})
const handleLogOut = () => {
  logout()
}
</script>

<template>
  <n-message-provider>
    <header>
      <n-space justify="space-between">
        <nav>
          <RouterLink
            class="router-link"
            v-for="route in filteredRoutes"
            :key="route.path"
            :to="route.path"
            >{{ route.name }}</RouterLink
          >
        </nav>
        <n-button @click="handleLogOut" v-if="isLogin"> 登出 </n-button>
      </n-space>
    </header>
    <div class="wrapper">
      <RouterView />
    </div>
  </n-message-provider>
</template>

<style scoped lang="less">
header {
  height: 50px;
  padding: 4px 20px;
  background-color: lightblue;
  margin-bottom: 30px;
  .router-link {
    font-size: 20px;
    font-weight: bold;
    margin-right: 40px;
  }
}
.wrapper {
  padding: 0 10px;
}
</style>
