import axios from 'axios'
import Cookies from 'js-cookie'

export const Settings = {
  host: 'http://localhost:8080/'
}

export function Registration (name, pass, mail) {
  return axios.post(Settings.host + 'api/register', {
    login: name,
    password: pass,
    email: mail
  })
}

export function Login (name, pass) {
  return axios.post(Settings.host + 'api/login', {
    login: name,
    password: pass
  })
}

export function GetFormTask () {
  return axios.get(Settings.host + 'api/tasks/make')
}

export function MakeNewTask (obj) {
  console.log(obj)
  return axios.post(Settings.host + 'api/tasks/make', obj)
}

export function SetCookie (token) {
  Cookies.set('Authorization', token)
}

export function GetCookie () {
  if (Cookies.get('Authorization')) {
    return axios.get(Settings.host + 'api/check-token/' + Cookies.get('Authorization'))
  } else {
    throw new Error('Ошибка проверки токена')
  }
}
