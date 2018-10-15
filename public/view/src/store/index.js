import Vue from 'vue'
import Vuex from 'vuex'
import tasks from './tasks'
import user from './user'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    tasks,
    user
  }
})
