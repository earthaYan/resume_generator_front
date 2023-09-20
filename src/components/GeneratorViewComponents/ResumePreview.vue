<script setup lang="ts">
import { NAvatar, NDataTable, NDescriptions, NDescriptionsItem, NDivider } from 'naive-ui'
import { generateEductionColumns, generateJobColumns } from './index.data'
import { useFormStore } from '@/stores/form'
import { onMounted } from 'vue'
import { instance } from '@/utils/api'
import { useRouter } from 'vue-router'
const { formValue, setDefaultInfo } = useFormStore()
const router = useRouter()
onMounted(() => {
  const { resume_id } = router.currentRoute.value.query
  if (resume_id) {
    instance.get(`/api/resume/${resume_id}`).then((res) => {
      setDefaultInfo(res.data)
    })
  }
})
</script>
<template>
  <div class="resume">
    <div class="left-side">
      <div class="profile">
        <n-avatar :size="120" round src="/avatar.jpg" />
      </div>
      <div class="baseInfo block">
        <h2>基础信息</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item label="年龄"> {{ formValue.base_info.age }}</n-descriptions-item>
            <n-descriptions-item label="工作年限">
              {{ formValue.base_info.experience_years }}年
            </n-descriptions-item>
            <n-descriptions-item label="联系电话">
              {{ formValue.base_info.phone }}
            </n-descriptions-item>
            <n-descriptions-item label="邮箱">
              {{ formValue.base_info.email }}
            </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
      <n-divider />
      <div class="jobRequire block">
        <h2>求职意向</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item label="意向岗位">
              {{ formValue.career_target.position }}
            </n-descriptions-item>
            <n-descriptions-item label="意向城市">
              {{ formValue.career_target.city }}
            </n-descriptions-item>
            <n-descriptions-item label="薪资要求">
              {{ formValue.career_target.salary }}
            </n-descriptions-item>
            <n-descriptions-item label="求职状态">
              {{ formValue.career_target.status }}
            </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
      <n-divider />
      <div class="skills block">
        <h2>专业技能</h2>
        <div class="item">
          <n-descriptions label-placement="left" :column="1">
            <n-descriptions-item
              v-for="skill in formValue.skills"
              :key="skill.name"
              :label="skill.name"
              >{{ skill.percent }}
            </n-descriptions-item>
          </n-descriptions>
        </div>
      </div>
    </div>
    <div class="split"></div>
    <div class="right-side">
      <h1 class="name">{{ formValue.base_info.name }}</h1>
      <div class="block breief">
        <p class="para">{{ formValue.base_info.brief_introduction }}</p>
      </div>
      <div class="block education">
        <h2>教育背景</h2>
        <n-data-table :columns="generateEductionColumns()" :data="formValue.education" />
      </div>
      <div class="block job">
        <h2>工作经历</h2>
        <n-data-table :columns="generateJobColumns()" :data="formValue.work_experience" />
      </div>
      <div class="block project">
        <h2>项目经验</h2>
        <n-descriptions label-placement="top" :column="1">
          <n-descriptions-item
            v-for="project in formValue.project_experience"
            :key="project.time"
            label="2020~2021"
          >
            {{ project.project_info }}
          </n-descriptions-item>
        </n-descriptions>
      </div>
    </div>
  </div>
</template>
<style lang="less" scoped>
h1 {
  font-weight: bold;
  font-size: 2em;
}

.resume {
  box-sizing: border-box;
  display: flex;
  background-color: #b0c4de;
  border: 1px solid #000;
  padding: 40px;

  .left-side {
    flex-grow: 0.3;
  }

  .right-side {
    flex-grow: 1;
    margin-top: 10px;
  }

  .split {
    width: 1px;
    background-color: #fff;
    margin: -40px 20px;
  }

  .block {
    margin-bottom: 20px;

    > h2 {
      margin-bottom: 10px;
    }
  }
}
</style>
