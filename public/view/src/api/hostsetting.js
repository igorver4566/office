import axios from 'axios'
import Cookies from 'js-cookie'

export const Settings = {
  host: 'http://softsetters-office.ru:8086/',
  auth: {headers: {
    Authorization: 'Bearer ' + Cookies.get('Authorization')
  }}
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
  return axios.get(Settings.host + 'api/tasks/make', Settings.auth)
}

export function MakeNewTask (obj) {
  return axios.post(Settings.host + 'api/tasks/make', obj, Settings.auth)
}

export function MakeNewSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/make', obj, Settings.auth)
}

export function EditSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/edit', obj, Settings.auth)
}

export function ChangeStatusSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/editStatus', obj, Settings.auth)
}

export function TrueTimeSubTask (obj) {
  return axios.post(Settings.host + 'api/subtasks/editTime', obj, Settings.auth)
}

export function GetTasks () {
  return axios.get(Settings.host + 'api/tasks', Settings.auth)
}

export function GetTaskById (id) {
  return axios.get(Settings.host + 'api/task/' + id, Settings.auth)
}

export function GetSubTasks () {
  return axios.get(Settings.host + 'api/subtasks', Settings.auth)
}

export function SetCookie (token) {
  Cookies.set('Authorization', token)
}

export function GetCookie () {
  return axios.get(Settings.host + 'api/check-token/' + Cookies.get('Authorization'), Settings.auth)
}

export function GetMessages (name) {
  return axios.get(Settings.host + 'api/chat/' + name, Settings.auth)
}

export function sendMsg (msg) {
  return axios.post(Settings.host + 'api/chat/' + msg.chat, msg.msg, Settings.auth)
}
