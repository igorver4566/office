import {GetFormTask, MakeNewTask, GetTasks, GetTaskById} from '@/api/hostsetting.js'

export default {
  state: {
    task: {
      task: {},
      items: []
    },
    tasks: [],
    form: null
  },
  mutations: {
    setForm (state, fields) {
      state.form = fields
    },
    setTasks (state, tasks) {
      state.tasks = tasks
    },
    setTask (state, task) {
      state.task = task
    }
  },
  actions: {
    getFormFields ({commit, state}) {
      commit('clearErrorOk')
      if (state.form === null) {
        GetFormTask().then((r) => {
          commit('setForm', r.data)
        }).catch((err) => {
          commit('setError', err)
        })
      }
    },
    makeTask ({commit}, payload) {
      commit('clearErrorOk')
      MakeNewTask(payload)
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setOk', response.data.data)
        } else {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    getTasks ({commit}) {
      commit('clearErrorOk')
      GetTasks()
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setTasks', response.data.data.slice(0, 10))
        } else {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    getTaskById ({commit}, id) {
      commit('clearErrorOk')
      GetTaskById(id)
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setTask', response.data.data)
        } else {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    }
  },
  getters: {
    form (state) {
      return state.form
    },
    tasks (state) {
      return state.tasks
    },
    task (state) {
      return state.task
    }
  }
}
