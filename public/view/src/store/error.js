export default {
  state: {
    error: null,
    ok: null
  },
  mutations: {
    setOk (state, ok) {
      state.ok = ok
    },
    setError (state, err) {
      state.error = err
    },
    clearErrorOk (state) {
      state.error = null
      state.ok = null
    }
  },
  actions: {},
  getters: {
    error (state) {
      return state.error
    },
    ok (state) {
      return state.ok
    }
  }
}
