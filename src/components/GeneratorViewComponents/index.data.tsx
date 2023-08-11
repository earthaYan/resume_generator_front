import type { FormRules } from 'naive-ui'

export const generateEductionColumns = () => {
  return [
    {
      title: '起止时间',
      key: 'time'
    },
    {
      title: '学校信息',
      key: 'school'
    }
  ]
}
export const generatePersonalColumns = () => {
  return [
    {
      title: '出生时间',
      key: 'birth'
    },
    {
      title: '常住地',
      key: 'address'
    }
  ]
}
export const generateJobColumns = () => {
  return [
    {
      title: '起止时间',
      key: 'time'
    },
    {
      title: '公司信息',
      key: 'company'
    },
    {
      title: '职位',
      key: 'position'
    }
  ]
}
const phoneExp = new RegExp('^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])d{8}$')
export const formRules: FormRules = {
  baseInfo: {
    name: {
      required: true,
      message: '不能为空',
      trigger: ['input', 'blur']
    },
    phone: [
      {
        required: true,
        validator: (_, value: string) => {
          if (!value) {
            return new Error('不能为空')
          } else if (!phoneExp.test(value)) {
            return new Error('输入合法的手机号码')
          }
          return true
        },
        trigger: ['input', 'blur']
      }
    ]
  }
}
