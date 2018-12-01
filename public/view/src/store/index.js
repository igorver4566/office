import Vue from 'vue'
import Vuex from 'vuex'
import tasks from './tasks'
import subTasks from './subTasks'
import user from './user'
import errors from './error'
import chat from './chat'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    tasks,
    subTasks,
    user,
    chat,
    errors
  }
})
