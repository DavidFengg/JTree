import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    uniquePatients: [],
    uniqueSamples: []
  },

  mutations: {
    addUniquePatients(state, patients) {
      state.uniquePatients = patients;
    },

    addUniqueSamples(state, samples) {
      state.uniqueSamples = samples;
    },
  },

  actions: {
    addUniquePatients({ commit }, patients) {
      commit('addUniquePatients', patients);
    },

    addUniqueSamples({ commit }, samples) {
      commit('addUniqueSamples', samples);
    }
  },

  getters: {
    patients: (state) => {
      return state.uniquePatients;
    },

    samples: (state) => {
      return state.uniqueSamples;
    }
  }
})
