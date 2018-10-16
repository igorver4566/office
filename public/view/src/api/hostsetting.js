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

export function Login (name, pass, mail) {
  return axios.post(Settings.host + 'api/login', {
    login: name,
    password: pass
  })
}

export function SetCookie (token) {
  Cookies.set('Authorization', token)
}

export function GetCookie () {
  return Cookies.get('api.example.com')
}
