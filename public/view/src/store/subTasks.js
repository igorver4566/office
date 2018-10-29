import {GetSubTasks} from '@/api/hostsetting.js'

export default {
  state: {
    subTasks: []
  },
  mutations: {
    setSubTasks (state, data) {
      state.subTasks = data
    }
  },
  actions: {
    getSubTasks ({commit}) {
      commit('clearErrorOk')
      GetSubTasks()
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setSubTasks', response.data.data.slice(0, 15))
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
    subTasks (state) {
      return state.subTasks
    }
  }
}