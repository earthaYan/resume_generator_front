import { ref } from 'vue'
import { defineStore } from 'pinia'
export interface formFieldsType {
  title: string
  base_info: {
    name: string
    age: number
    experience_years: number
    email: string
    phone: string
    brief_introduction: string
  }
  career_target: {
    position: string
    city: string
    salary: string
    status: string
  }
  education: Array<{
    time: string
    school: string
    degree: string
  }>
  work_experience: Array<{
    time: string
    position: string
    company: string
  }>
  project_experience: Array<{
    time: string
    project_info: string
  }>
  skills: Array<{
    name: string
    percent: string
  }>
}

export const useFormStore = defineStore('formValue', () => {
  const formValue = ref<formFieldsType>({
    title: '简历1',
    base_info: {
      name: '李四',
      age: 26,
      experience_years: 3,
      email: '256790@qq.com',
      phone: '12345789870',
      brief_introduction: '一句话介绍自己'
    },
    career_target: {
      position: '前端开发',
      city: '上海',
      salary: '1-2k',
      status: '在职'
    },
    work_experience: [
      {
        time: '2020-2023',
        position: '前端开发',
        company: 'xx公司'
      }
    ],
    education: [
      {
        time: '2020-2023',
        school: '大学',
        degree: '学士'
      }
    ],
    project_experience: [
      {
        time: '2020-2023',
        project_info: '项目信息项目内容'
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
  function setDefaultInfo(data: formFieldsType) {
    formValue.value = data
  }
  return {
    formValue,
    setDefaultInfo
  }
})
