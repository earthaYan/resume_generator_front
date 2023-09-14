<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { NMessageProvider, NButton } from 'naive-ui'
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
      return route.name !== 'Login'
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
      <nav>
        <RouterLink
          class="router-link"
          v-for="route in filteredRoutes"
          :key="route.path"
          :to="route.path"
          >{{ route.name }}</RouterLink
        >
        <n-button @click="handleLogOut" v-if="isLogin"> 登出 </n-button>
      </nav>
    </header>
    <div class="wrapper">
      <RouterView />
    </div>
  </n-message-provider>
</template>

<style scoped lang="less">
header {
  height: 50px;
  nav {
    text-align: center;
    background-color: antiquewhite;
  }
  .router-link {
    font-size: 20px;
    font-weight: bold;
    margin-right: 20px;
  }
}
.wrapper {
  padding: 0 10px;
}
</style>
