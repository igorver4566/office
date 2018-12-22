import {GetSubTasks, MakeNewSubTask, EditSubTask, ChangeStatusSubTask} from '@/api/hostsetting.js'

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
    },
    getAllSubTasks ({commit}) {
      commit('clearErrorOk')
      GetSubTasks()
      .then((response) => {
        if (response.data.ok === 'true') {
          commit('setSubTasks', response.data.data)
        } else {
          commit('setError', response.data.data)
        }
      })
      .catch(err => {
        commit('setError', err)
      })
    },
    makeSubTask ({commit}, payload) {
      commit('clearErrorOk')
      MakeNewSubTask(payload)
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
    editSubTask ({commit}, payload) {
      commit('clearErrorOk')
      EditSubTask(payload)
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
    changeStatusSubTask ({commit}, payload) {
      commit('clearErrorOk')
      ChangeStatusSubTask(payload)
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
    }
  },
  getters: {
    subTasks (state) {
      return state.subTasks
    }
  }
}
