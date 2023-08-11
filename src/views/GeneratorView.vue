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
</script>
<template>
  <n-card title="添加简历">
    <n-gradient-text
      class="resume-generator-header"
      gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)"
    >
      开始定制
    </n-gradient-text>
    <template #header-extra>
      <n-space>
        <n-button type="primary">预览</n-button>
        <n-button type="info">提交</n-button>
      </n-space>
    </template>
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
