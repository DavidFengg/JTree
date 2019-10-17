import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    uniquePatients: []
  },
  mutations: {
    addUniquePatients(state, patients) {
      state.uniquePatients = patients;
    }
  },
  actions: {
    addUniquePatients ({ commit }, patients) {
      commit('addUniquePatients', patients);
    }
  },
  getters: {
    patients: (state) => {
      return state.uniquePatients;
    }
  }
})
