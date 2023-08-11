<script setup lang="ts">
import { routes } from '@/router'
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <nav>
      <RouterLink class="router-link" v-for="route in routes" :key="route.path" :to="route.path">{{
        route.name
      }}</RouterLink>
    </nav>
  </header>
  <div class="wrapper">
    <RouterView />
  </div>
 
</template>

<style scoped lang="less">
header {
  height: 100px;
  nav {
    text-align: center;
    padding: 2rem;
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
