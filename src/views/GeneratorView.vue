<script setup lang="ts">
import {
  NCard,
  NGradientText,
  NLayout,
  NLayoutSider,
  NLayoutContent,
  NButton,
  NSpace
} from 'naive-ui'
import { ResumeForm, ResumePreview } from '@/components/GeneratorViewComponents'
import { useFormStore } from '@/stores/form'
import { instance } from '@/utils/api'
import { useRouter } from 'vue-router'
const { formValue } = useFormStore()
const router = useRouter()
const handleSubmit = () => {
  // 提交
  instance.post('/resume/create', formValue).then((res) => {
    console.log(res)
  })
}
const id = router.currentRoute.value.params.resume_id
const handleUpdate = () => {
  instance
    .post('/resume/update', {
      resume_id: id,
      info: formValue
    })
    .then((res) => {
      console.log(res)
    })
}
</script>
<template>
  <n-card title="添加简历">
    <template #header-extra>
      <n-space>
        <n-button v-if="!id" type="info" @click="handleSubmit">提交</n-button>
        <n-button v-if="!!id" type="info" @click="handleUpdate">更新</n-button>
      </n-space>
    </template>
    <n-gradient-text
      class="resume-generator-header"
      gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)"
    >
      开始定制
    </n-gradient-text>
    <div class="resume-generator-wrapper">
      <n-layout has-sider sider-placement="right">
        <n-layout-content content-style="padding: 24px;">
          <ResumeForm />
        </n-layout-content>
        <n-layout-sider
          collapse-mode="width"
          :collapsed-width="0"
          width="50%"
          :show-collapsed-content="false"
          show-trigger="arrow-circle"
          content-style="padding: 24px;"
          bordered
        >
          <ResumePreview />
        </n-layout-sider>
      </n-layout>
    </div>
  </n-card>
</template>
<style lang="less" scoped>
.n-gradient-text {
  font-size: 24px;
  margin-bottom: 10px;
}
</style>
