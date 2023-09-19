import axios from 'axios'
export const instance = axios.create({
  baseURL: 'http://localhost:3000/api/v1'
})
const whitePath = ['/login', '/register']
instance.interceptors.request.use(
  (config) => {
    // 在发送请求之前做些什么
    const pathname = location.pathname
    const token = localStorage.getItem('token')
    if (localStorage.getItem('token')) {
      if (!whitePath.includes(pathname)) {
        config.headers.Authorization = token
      }
    }
    return config
  },
  (error) => {
    //对请求错误做些什么
    return Promise.reject(error)
  }
)
instance.interceptors.response.use((response) => {
  const { status } = response

  if (status === 200) {
    return Promise.resolve(response)
  } else {
    return Promise.reject(response)
  }
})
