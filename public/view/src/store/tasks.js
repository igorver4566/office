import {GetFormTask, MakeNewTask} from '@/api/hostsetting.js'

export default {
  state: {
    tasks: [],
    form: null
  },
  mutations: {
    setForm (state, fields) {
      state.form = fields
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
        console.log(payload)
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
    }
  },
  getters: {
    form (state) {
      return state.form
    }
  }
}
