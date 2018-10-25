import Vue from 'vue'
import Vuex from 'vuex'
import tasks from './tasks'
import subTasks from './subTasks'
import user from './user'
import errors from './error'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    tasks,
    subTasks,
    user,
    errors
  }
})
