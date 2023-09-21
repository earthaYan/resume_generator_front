<script setup lang="ts">
import {
  NForm,
  NFormItem,
  type FormInst,
  NInput,
  NDynamicInput,
  NTabs,
  NTabPane,
  NInputNumber
} from 'naive-ui'
import { ref } from 'vue'
import { useFormStore } from '@/stores/form'
const formRef = ref<FormInst | null>(null)

const onCreate = (type: string) => {
  if (type === 'education') {
    return {
      time: '',
      school: '',
      degree: ''
    }
  } else if (type === 'work') {
    return {
      time: '',
      position: '',
      company: ''
    }
  } else {
    return {
      time: '',
      projectInfo: ''
    }
  }
}
const { formValue } = useFormStore()
</script>
<template>
  <n-form
    ref="formRef"
    :model="formValue"
    label-placement="left"
    label-align="right"
    label-width="100"
  >
    <n-form-item label="简历标题" path="title" label-align="left">
      <n-input v-model:value="formValue.title" />
    </n-form-item>
    <n-tabs>
      <n-tab-pane name="基本信息">
        <div class="section baseInfo">
          <h2 class="title">基础信息</h2>
          <n-form-item label="姓名" path="base_info.name" :required="true">
            <n-input v-model:value="formValue.base_info.name" placeholder="输入姓名" />
          </n-form-item>
          <n-form-item label="年龄" path="base_info.age" :required="true">
            <n-input-number
              v-model:value="formValue.base_info.age"
              placeholder="输入年龄"
              max="60"
              min="16"
            />
          </n-form-item>
          <n-form-item label="工作年限" path="base_info.experience_years" :required="true">
            <n-input-number
              v-model:value="formValue.base_info.experience_years"
              placeholder="输入工作年限"
            />
          </n-form-item>
          <n-form-item label="联系电话" path="base_info.phone" :required="true">
            <n-input v-model:value="formValue.base_info.phone" placeholder="输入联系电话" />
          </n-form-item>
          <n-form-item label="电子邮箱" path="base_info.email">
            <n-input v-model:value="formValue.base_info.email" placeholder="输入电子邮箱" />
          </n-form-item>
          <n-form-item label="简介" path="base_info.brief_introduction">
            <n-input
              v-model:value="formValue.base_info.brief_introduction"
              placeholder="一句话介绍自己"
            />
          </n-form-item>
        </div>
      </n-tab-pane>
      <n-tab-pane name="求职意向">
        <div class="section careerObjective">
          <h2 class="title">求职意向</h2>
          <n-form-item label="意向岗位" path="career_target.position">
            <n-input
              v-model:value="formValue.career_target.position"
              placeholder="请输入意向岗位"
            />
          </n-form-item>
          <n-form-item label="意向城市" path="career_target.city">
            <n-input v-model:value="formValue.career_target.city" placeholder="请输入意向城市" />
          </n-form-item>
          <n-form-item label="薪资要求" path="career_target.salary">
            <n-input v-model:value="formValue.career_target.salary" placeholder="请输入薪资要求" />
          </n-form-item>
          <n-form-item label="求职状态" path="career_target.status">
            <n-input v-model:value="formValue.career_target.status" placeholder="请输入求职状态" />
          </n-form-item>
        </div>
      </n-tab-pane>
      <n-tab-pane name="专业技能">
        <div class="section skills">
          <h2 class="title">专业技能</h2>
          <n-form-item>
            <n-dynamic-input
              v-model:value="formValue.skills"
              #="{ index }"
              :on-create="() => onCreate('skills')"
            >
              <div style="display: flex; justify-content: space-between">
                <n-form-item path="formValue.skills[index].name">
                  <n-input
                    v-model:value="formValue.skills[index].name"
                    placeholder="请输入技能名称"
                  />
                </n-form-item>
                <n-form-item path="formValue.skills[index].percent">
                  <n-input
                    v-model:value="formValue.skills[index].percent"
                    placeholder="请输入技能程度"
                  />
                </n-form-item>
              </div>
            </n-dynamic-input>
          </n-form-item>
        </div>
      </n-tab-pane>
      <n-tab-pane name="教育背景">
        <div class="section educationInfo">
          <h2 class="title">教育背景</h2>
          <n-form-item>
            <n-dynamic-input
              v-model:value="formValue.education"
              #="{ index }"
              :on-create="() => onCreate('education')"
            >
              <div style="display: flex; justify-content: space-between">
                <n-form-item path="formValue.education[index].time">
                  <n-input
                    v-model:value="formValue.education[index].time"
                    placeholder="请输入年份"
                  />
                </n-form-item>
                <n-form-item path="formValue.education[index].school">
                  <n-input
                    v-model:value="formValue.education[index].school"
                    placeholder="请输入学校信息"
                  />
                </n-form-item>
                <n-form-item path="formValue.education[index].degree">
                  <n-input
                    v-model:value="formValue.education[index].degree"
                    placeholder="请输入学位信息"
                  />
                </n-form-item>
              </div>
            </n-dynamic-input>
          </n-form-item>
        </div>
      </n-tab-pane>
      <n-tab-pane name="工作经历">
        <div class="section workExperience">
          <h2 class="title">工作经历</h2>
          <n-form-item>
            <n-dynamic-input
              v-model:value="formValue.work_experience"
              #="{ index }"
              :on-create="() => onCreate('work')"
            >
              <div style="display: flex; justify-content: space-between">
                <n-form-item path="formValue.work_experience[index].time">
                  <n-input
                    v-model:value="formValue.work_experience[index].time"
                    placeholder="请输入年份"
                  />
                </n-form-item>
                <n-form-item path="formValue.work_experience[index].company">
                  <n-input
                    v-model:value="formValue.work_experience[index].company"
                    placeholder="请输入公司信息"
                  />
                </n-form-item>
                <n-form-item path="formValue.work_experience[index].position">
                  <n-input
                    v-model:value="formValue.work_experience[index].position"
                    placeholder="请输入职位信息"
                  />
                </n-form-item>
              </div>
            </n-dynamic-input>
          </n-form-item>
        </div>
      </n-tab-pane>
      <n-tab-pane name="项目经验">
        <div class="section projectInfo">
          <h2 class="title">项目经验</h2>
          <n-form-item>
            <n-dynamic-input
              v-model:value="formValue.project_experience"
              #="{ index }"
              :on-create="() => onCreate('project')"
            >
              <div style="display: flex; justify-content: space-between">
                <n-form-item path="formValue.project_experience[index].time">
                  <n-input
                    v-model:value="formValue.project_experience[index].time"
                    placeholder="请输入年份"
                  />
                </n-form-item>
                <n-form-item path="formValue.project_experience[index].project_info">
                  <n-input
                    v-model:value="formValue.project_experience[index].project_info"
                    placeholder="请输入项目信息"
                  />
                </n-form-item>
              </div>
            </n-dynamic-input>
          </n-form-item>
        </div>
      </n-tab-pane>
    </n-tabs>
  </n-form>
</template>
<style scoped lang="less">
.section {
  box-shadow: 10px 7px 5px -5px gray;
  border: 1px solid gray;
  margin: 5px;
  padding: 10px;

  .title {
    margin-bottom: 20px;
  }
}
</style>
