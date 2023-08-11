<script setup lang="ts">
import { type formFieldsType } from './index.d'
import { NForm, NFormItem, type FormInst, NInput, NDynamicInput } from 'naive-ui'
import { ref } from 'vue'
const formRef = ref<FormInst | null>(null)
const formValue = ref<formFieldsType>({
  baseInfo: {
    name: '',
    age: '',
    experienceYears: '',
    email: '',
    phone: '',
    briefIntroduction: ''
  },
  careerObjective: {
    position: '',
    city: '',
    salary: '',
    status: ''
  },
  workExperience: [
    {
      time: '',
      workInfo: ''
    }
  ],
  educationInfo: [
    {
      time: '',
      school: '',
      master: ''
    }
  ],
  projectExperience: [
    {
      time: '',
      projectInfo: ''
    }
  ],
  skills: ''
})
const onCreate = (type: string) => {
  if (type === 'education') {
    return {
      time: '',
      school: '',
      master: ''
    }
  } else if (type === 'work') {
    return {
      time: '',
      workInfo: ''
    }
  } else {
    return {
      time: '',
      projectInfo: ''
    }
  }
}
</script>
<template>
  <n-form
    ref="formRef"
    v-model="formValue"
    label-placement="left"
    label-align="right"
    label-width="auto"
  >
    <div class="section baseInfo">
      <h2 class="title">基础信息</h2>
      <n-form-item label="姓名" path="baseInfo.name">
        <n-input v-model:value="formValue.baseInfo.name" placeholder="输入姓名" />
      </n-form-item>
      <n-form-item label="年龄" path="baseInfo.age">
        <n-input v-model:value="formValue.baseInfo.age" placeholder="输入年龄" />
      </n-form-item>
      <n-form-item label="工作年限" path="baseInfo.experienceYears">
        <n-input v-model:value="formValue.baseInfo.experienceYears" placeholder="输入工作年限" />
      </n-form-item>
      <n-form-item label="联系电话" path="baseInfo.phone">
        <n-input v-model:value="formValue.baseInfo.phone" placeholder="输入联系电话" />
      </n-form-item>
      <n-form-item label="电子邮箱" path="baseInfo.email">
        <n-input v-model:value="formValue.baseInfo.email" placeholder="输入电子邮箱" />
      </n-form-item>
      <n-form-item label="简介" path="baseInfo.briefIntroduction">
        <n-input
          v-model:value="formValue.baseInfo.briefIntroduction"
          placeholder="一句话介绍自己"
        />
      </n-form-item>
    </div>
    <div class="section skills">
      <h2 class="title">专业技能</h2>
      <n-form-item label="专业技能" path="skills">
        <n-input v-model:value="formValue.skills" placeholder="输入技能" />
      </n-form-item>
    </div>
    <div class="section careerObjective">
      <h2 class="title">求职意向</h2>
      <n-form-item label="意向岗位" path="careerObjective.position">
        <n-input v-model:value="formValue.careerObjective.position" placeholder="请输入意向岗位" />
      </n-form-item>
      <n-form-item label="意向城市" path="careerObjective.city">
        <n-input v-model:value="formValue.careerObjective.city" placeholder="请输入意向城市" />
      </n-form-item>
      <n-form-item label="薪资要求" path="careerObjective.salary">
        <n-input v-model:value="formValue.careerObjective.salary" placeholder="请输入薪资要求" />
      </n-form-item>
      <n-form-item label="求职状态" path="careerObjective.status">
        <n-input v-model:value="formValue.careerObjective.status" placeholder="请输入求职状态" />
      </n-form-item>
    </div>
    <div class="section educationInfo">
      <h2 class="title">教育背景</h2>
      <n-form-item>
        <n-dynamic-input
          v-model:value="formValue.educationInfo"
          #="{ index }"
          :on-create="() => onCreate('education')"
        >
          <div style="display: flex; justify-content: space-between">
            <n-form-item path="formValue.educationInfo[index].time">
              <n-input
                v-model:value="formValue.educationInfo[index].time"
                placeholder="请输入年份"
              />
            </n-form-item>
            <n-form-item path="formValue.educationInfo[index].school">
              <n-input
                v-model:value="formValue.educationInfo[index].school"
                placeholder="请输入学校信息"
              />
            </n-form-item>
            <n-form-item path="formValue.educationInfo[index].master">
              <n-input
                v-model:value="formValue.educationInfo[index].master"
                placeholder="请输入学位信息"
              />
            </n-form-item>
          </div>
        </n-dynamic-input>
      </n-form-item>
    </div>
    <div class="section workExperience">
      <h2 class="title">工作经历</h2>
      <n-form-item>
        <n-dynamic-input
          v-model:value="formValue.workExperience"
          #="{ index }"
          :on-create="() => onCreate('work')"
        >
          <div style="display: flex; justify-content: space-between">
            <n-form-item path="formValue.workExperience[index].time">
              <n-input
                v-model:value="formValue.workExperience[index].time"
                placeholder="请输入年份"
              />
            </n-form-item>
            <n-form-item path="formValue.workExperience[index].workInfo">
              <n-input
                v-model:value="formValue.workExperience[index].workInfo"
                placeholder="请输入公司信息"
              />
            </n-form-item>
          </div>
        </n-dynamic-input>
      </n-form-item>
    </div>
    <div class="section projectInfo">
      <h2 class="title">项目经验</h2>
      <n-form-item>
        <n-dynamic-input
          v-model:value="formValue.projectExperience"
          #="{ index }"
          :on-create="() => onCreate('project')"
        >
          <div style="display: flex; justify-content: space-between">
            <n-form-item path="formValue.projectExperience[index].time">
              <n-input
                v-model:value="formValue.projectExperience[index].time"
                placeholder="请输入年份"
              />
            </n-form-item>
            <n-form-item path="formValue.projectExperience[index].projectInfo">
              <n-input
                v-model:value="formValue.projectExperience[index].projectInfo"
                placeholder="请输入项目信息"
              />
            </n-form-item>
          </div>
        </n-dynamic-input>
      </n-form-item>
    </div>
  </n-form>
</template>
<style scoped lang="less">
.section {
  box-shadow: 5px 1px 10px 1px gray;
  padding: 20px;
  margin-bottom: 30px;

  .title {
    margin-bottom: 20px;
  }
}
</style>
