import {Registration, Login} from '@/api/hostsetting.js'

export default {
  state: {
    user: null,
    userError: null
  },
  mutations: {
    setToken (state, token) {
      state.user = token
    },
    setError (state, err) {
      state.userError = err
    },
    clearError (state) {
      state.userError = null
    }
  },
  actions: {
    register ({commit}, payload) {
      commit('clearError')
      return new Promise((resolve, reject) => {
        Registration(payload.login, payload.password, payload.email)
        .then(function (response) {
          if (response.data.ok === 'true') {
            commit('setToken', response.data.token)
          } else {
            commit('setError', response.data.token)
            return reject(response.data.token)
          }
          return resolve()
        })
        .catch(err => {
          commit('setError', err)
          return reject(err)
        })
      })
    },
    auth ({commit}, payload) {
      commit('clearError')
      return new Promise((resolve, reject) => {
        Login(payload.login, payload.password)
        .then(function (response) {
          if (response.data.ok === 'true') {
            commit('setToken', response.data.token)
          } else {
            commit('setError', response.data.token)
            return reject(response.data.token)
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
    user (state) {
      return state.user
    },
    error (state) {
      return state.userError
    }
  }
}
