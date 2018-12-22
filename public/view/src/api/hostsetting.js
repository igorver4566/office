import axios from 'axios'
import Cookies from 'js-cookie'

export const Settings = {
  host: 'http://softsetters-office.ru:8086/'
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
  return axios.post(Settings.host + 'api/tasks/make', obj)
}

export function MakeNewSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/make', obj)
}

export function EditSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/edit', obj)
}

export function ChangeStatusSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/editStatus', obj)
}

export function GetTasks () {
  return axios.get(Settings.host + 'api/tasks')
}

export function GetTaskById (id) {
  return axios.get(Settings.host + 'api/task/' + id)
}

export function GetSubTasks () {
  return axios.get(Settings.host + 'api/subtasks')
}

export function SetCookie (token) {
  Cookies.set('Authorization', token)
}

export function GetCookie () {
  return axios.get(Settings.host + 'api/check-token/' + Cookies.get('Authorization'))
}

export function GetMessages (name) {
  return axios.get(Settings.host + 'api/chat/' + name)
}

export function sendMsg (msg) {
  return axios.post(Settings.host + 'api/chat/' + msg.chat, msg.msg)
}
