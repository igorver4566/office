import {GetMessages, sendMsg} from '@/api/hostsetting.js'

export default {
  state: {
    messages: []
  },
  mutations: {
    setMessages (state, data) {
      state.messages = data
    }
  },
  actions: {
    getMessages ({commit}, payload) {
      commit('clearErrorOk')
      GetMessages(payload)
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setMessages', response.data.data)
        } else {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    sendMsg ({commit}, payload) {
      commit('clearErrorOk')
      sendMsg(payload)
      .then((response) => {
        if (response.data.ok !== 'true') {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    }
  },
  getters: {
    messages (state) {
      return state.messages
    }
  }
}
