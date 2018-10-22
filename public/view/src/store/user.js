import {Registration, Login, SetCookie, GetCookie} from '@/api/hostsetting.js'
import {WeatherLocation} from '@/api/weather.js'

export default {
  state: {
    user: null,
    userLogin: null,
    userName: null,
    location: 'St. Petersburg',
    weather: 0
  },
  mutations: {
    setToken (state, token) {
      state.user = token
    },
    setLocation (state, data) {
      state.location = data.location.city
      state.weather = data.item.condition.temp
    },
    setUserFields (state, data) {
      state.userLogin = data.login
      state.userName = data.name
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
      commit('clearErrorOk')
      GetCookie().then((r) => {
        if (r.data.ok === 'true') {
          commit('setToken', r.data.data.token)
          commit('setUserFields', r.data.data)
        }
      }).catch(err => {
        commit('setError', err)
      })
    },
    userLocation ({commit}) {
      commit('clearErrorOk')
      WeatherLocation().then((r) => {
        commit('setLocation', r)
      }).catch(err => {
        commit('setError', err)
      })
    }
  },
  getters: {
    user (state) {
      return state.user
    },
    userName (state) {
      return state.userName
    },
    userLogin (state) {
      return state.userLogin
    },
    location (state) {
      return state.location
    },
    weather (state) {
      return state.weather
    }
  }
}
