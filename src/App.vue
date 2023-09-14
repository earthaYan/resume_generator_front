<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { NMessageProvider } from 'naive-ui'
import { routes } from './router'
</script>

<template>
  <n-message-provider>
    <header>
      <nav>
        <RouterLink
          class="router-link"
          v-for="route in routes"
          :key="route.path"
          :to="route.path"
          >{{ route.name }}</RouterLink
        >
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
