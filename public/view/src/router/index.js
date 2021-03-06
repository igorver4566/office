import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home/Home'
import Login from '@/components/Auth/Login'
import Register from '@/components/Auth/Register'
import MakeTask from '@/components/Tasks/MakeTask'
import Task from '@/components/Tasks/Task'
import AllTasks from '@/components/Tasks/AllTasks'
import AllSubTasks from '@/components/Tasks/AllSubTasks'
import Workers from '@/components/Workers/Workers'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },

    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/make_task',
      name: 'make_task',
      component: MakeTask
    },
    {
      path: '/tasks',
      name: 'tasks',
      component: AllTasks
    },
    {
      path: '/subtasks',
      name: 'subtask',
      component: AllSubTasks
    },
    {
      path: '/task/:id',
      name: 'task',
      props: true,
      component: Task
    },
    {
      path: '/workers',
      name: 'workers',
      component: Workers
    }
  ],
  mode: 'history'
})
