import axios from 'axios'
export const instance = axios.create({
  baseURL: 'http://116.204.108.126:3000'
})
