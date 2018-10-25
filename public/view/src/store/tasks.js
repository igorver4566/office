import {GetFormTask, MakeNewTask, GetTasks} from '@/api/hostsetting.js'

export default {
  state: {
    tasks: [],
    form: null
  },
  mutations: {
    setForm (state, fields) {
      state.form = fields
    },
    setTasks (state, tasks) {
      state.tasks = tasks
    }
  },
  actions: {
    getFormFields ({commit, state}) {
      commit('clearErrorOk')
      if (state.form === null) {
        return new Promise((resolve, reject) => {
          GetFormTask().then((r) => {
            commit('setForm', r.data)
            return resolve(r.data)
          }).catch((err) => {
            commit('setError', err)
            return reject(err)
          })
        })
      }
    },
    makeTask ({commit}, payload) {
      commit('clearErrorOk')
      return new Promise((resolve, reject) => {
        MakeNewTask(payload)
        .then((response) => {
          if (response.data.ok === 'true') {
            commit('setOk', response.data.data)
          } else {
            commit('setError', response.data.data)
          }
          return resolve()
        })
        .catch(err => {
          commit('setError', err)
          return reject(err)
        })
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
    }
  },
  getters: {
    form (state) {
      return state.form
    },
    tasks (state) {
      return state.tasks
    }
  }
}
