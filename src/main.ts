import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
// 通过 createApp(根组件) 创建一个新的应用实例
const app = createApp(App)
const pinia = createPinia()
app.use(router)
app.use(pinia)
// 挂载应用
app.mount('#app')
