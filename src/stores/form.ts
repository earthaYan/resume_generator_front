import { ref } from 'vue'
import { defineStore } from 'pinia'
export interface formFieldsType {
  baseInfo: {
    name: string
    age: number
    experienceYears: number
    email: string
    phone: string
    briefIntroduction: string
  }
  careerObjective: {
    position: string
    city: string
    salary: string
    status: string
  }
  educationInfo: Array<{
    time: string
    school: string
    degree: string
  }>
  workExperience: Array<{
    time: string
    position: string
    company: string
  }>
  projectExperience: Array<{
    time: string
    projectInfo: string
  }>
  skills: Array<{
    name: string
    percent: string
  }>
}

export const useFormStore = defineStore('formValue', () => {
  const formValue = ref<formFieldsType>({
    baseInfo: {
      name: '李四',
      age: 26,
      experienceYears: 3,
      email: '256790@qq.com',
      phone: '12345789870',
      briefIntroduction: '一句话介绍自己'
    },
    careerObjective: {
      position: '前端开发',
      city: '上海',
      salary: '1-2k',
      status: '在职'
    },
    workExperience: [
      {
        time: '2020-2023',
        position: '前端开发',
        company: 'xx公司'
      }
    ],
    educationInfo: [
      {
        time: '2020-2023',
        school: '大学',
        degree: '学士'
      }
    ],
    projectExperience: [
      {
        time: '2020-2023',
        projectInfo: '项目信息项目内容'
      }
    ],
    skills: [
      {
        name: 'JavaScript',
        percent: '熟练'
      },
      {
        name: 'Node',
        percent: '熟练'
      }
    ]
  })
  return {
    formValue
  }
})
