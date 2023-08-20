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
    master: string
  }>
  workExperience: Array<{
    time: string
    workInfo: string
    position:string
  }>
  projectExperience: Array<{
    time: string
    projectInfo: string
  }>
  skills: string
}
