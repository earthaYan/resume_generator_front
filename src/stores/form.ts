import { ref} from 'vue'
import { defineStore } from 'pinia'
export interface formFieldsType {
    baseInfo: {
      name: string
      age:number
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
      master: string
    }>
    workExperience: Array<{
      time: string
      position:string
      company:string
    }>
    projectExperience: Array<{
      time: string
      projectInfo: string
    }>
    skills: Array<{
      name:string
      percent:string
    }>
  }
  
export const useFormStore=defineStore('formValue',()=>{
    const formValue = ref<formFieldsType>({
        baseInfo: {
          name: '',
          age: 26,
          experienceYears: 1,
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
            position: '',
            company:''
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
        skills: [{
          name:'JavaScript',
          percent:'熟练'
        }]
      })
    return {
        formValue
    }
})