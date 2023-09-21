<script setup lang="ts">
import { instance } from '@/utils/api'
import { NButton, NDataTable, type DataTableColumns } from 'naive-ui'
import { h, onMounted, ref } from 'vue'
import type { ResumeItem } from '@/types/resume'
import dayjs from 'dayjs'
import router from '@/router'
let list = ref<
  Array<{
    id: number
    name: string
    update_time: string
  }>
>([])
const getList = () => {
  instance.get('/resumes').then((res) => {
    list.value = res.data.data.item
  })
}
onMounted(() => {
  getList()
})
const handleDelete = (row: ResumeItem) => {
  const resume_id = row.id
  instance
    .post('/resume_delete', {
      resume_id
    })
    .then((res) => {
      console.log(res)
      getList()
    })
}
const jumpToDetail = (row: ResumeItem) => {
  const resume_id = row.id
  router.push(`/update/${resume_id}`)
}
const createColumns = (): DataTableColumns<ResumeItem> => {
  return [
    {
      title: '简历名称',
      key: 'title'
    },
    {
      title: '作者',
      key: 'author'
    },
    {
      title: '创建时间',
      key: 'created_at',
      render(rowData) {
        return dayjs(rowData.created_at * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    },
    {
      title: '最后更新',
      key: 'updated_at',
      render(rowData) {
        return dayjs(rowData.updated_at * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    },
    {
      title: '操作',
      key: 'actions',
      render(row) {
        return [
          h(
            NButton,
            {
              strong: true,
              tertiary: true,
              size: 'small',
              onClick: () => handleDelete(row)
            },
            { default: () => '删除' }
          ),
          h(
            NButton,
            {
              strong: true,
              tertiary: true,
              size: 'small',
              onClick: () => jumpToDetail(row)
            },
            { default: () => '详情' }
          )
        ]
      }
    }
  ]
}
const columns = createColumns()
</script>
<template>
  <n-data-table :columns="columns" :data="list" :pagination="false" :bordered="false" />
</template>
