import {Registration, Login, SetCookie, GetCookie, GetWorkers} from '@/api/hostsetting.js'
import {WeatherLocation} from '@/api/weather.js'

export default {
  state: {
    user: null,
    userLogin: null,
    userName: null,
    location: 'St. Petersburg',
    weather: 0,
    workers: []
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
    },
    setWorkers (state, data) {
      state.workers = data
    }
  },
  actions: {
    register ({commit}, payload) {
      commit('clearErrorOk')
      Registration(payload.login, payload.password, payload.email)
      .then(function (r) {
        if (r.data.ok === 'true') {
          commit('setToken', r.data.data.token)
          commit('setUserFields', r.data.data)
          SetCookie(r.data.data.token)
        } else {
          commit('setError', r.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    auth ({commit}, payload) {
      commit('clearErrorOk')
      Login(payload.login, payload.password)
      .then(function (r) {
        if (r.data.ok === 'true') {
          commit('setToken', r.data.data.token)
          commit('setUserFields', r.data.data)
          SetCookie(r.data.data.token)
        } else {
          commit('setError', r.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    checkToken ({commit}) {
      commit('clearErrorOk')
      GetCookie().then((r) => {
        if (r.data.ok === 'true') {
          commit('setToken', r.data.data.token)
          commit('setUserFields', r.data.data)
        } else {
          commit('setError', r.data.data)
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
    },
    GetAllWorkers ({commit}, payload) {
      commit('clearErrorOk')
      GetWorkers(payload.dt_start, payload.dt_end)
      .then(function (r) {
        if (r.data.ok === 'true') {
          commit('setWorkers', r.data.data)
        } else {
          commit('setError', r.data.data)
        }
      })
      .catch(err => {
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
    },
    workers (state) {
      return state.workers
    }
  }
}
