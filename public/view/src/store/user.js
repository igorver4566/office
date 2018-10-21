import {Registration, Login, SetCookie, GetCookie} from '@/api/hostsetting.js'

export default {
  state: {
    user: null
  },
  mutations: {
    setToken (state, token) {
      state.user = token
    }
  },
  actions: {
    register ({commit}, payload) {
      commit('clearErrorOk')
      return new Promise((resolve, reject) => {
        Registration(payload.login, payload.password, payload.email)
        .then(function (response) {
          if (response.data.ok === 'true') {
            commit('setToken', response.data.token)
            SetCookie(response.data.token)
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
      commit('clearErrorOk')
      return new Promise((resolve, reject) => {
        Login(payload.login, payload.password)
        .then(function (response) {
          if (response.data.ok === 'true') {
            commit('setToken', response.data.token)
            SetCookie(response.data.token)
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
    checkToken ({commit}) {
      GetCookie().then((r) => {
        if (r.data.ok === 'true') {
          commit('setToken', r.data.data)
        }
      }).catch(() => {})
    }
  },
  getters: {
    user (state) {
      return state.user
    }
  }
}
