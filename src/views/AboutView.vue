<script setup lang="ts">
import axios from 'axios';
import { NList, NListItem,NSpace,NButton } from 'naive-ui'
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
let list = ref<Array<{
  id: number,
  name: string
  update_time: string
}>>([])
const router = useRouter()
const handleJump=(id:number)=>{
  router.push({
    path:'/add',
    query:{
      resume_id:id
    }
  })
}
const getList=()=>{
  axios.get('http://localhost:3000/api/resume/list').then(res => {
    list.value = res.data
  })
}
onMounted(() => {
  getList()
})
const handleDelete=(id:number)=>{
axios.delete(`http://localhost:3000/api/resume/${id}`).then(res=>{
  getList()
  
})
}
</script>
<template>
  <n-list hoverable clickable>
    <n-list-item v-for="item in list" :key="item.id">
      <n-space>
        <a @click="handleJump(item.id)">{{ item.name }} - 最后更新于{{ item.update_time }}</a>

        <n-button @click="handleDelete(item.id)">删除</n-button>
      </n-space>
      
    </n-list-item>
  </n-list>
</template>
