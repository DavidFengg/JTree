import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    uniquePatients: [],
    uniqueSamples: [],
    uniqueExperiments: []
  },

  mutations: {
    addUniquePatients(state, patients) {
      state.uniquePatients = patients;
    },

    addUniqueSamples(state, samples) {
      state.uniqueSamples = samples;
    },

    addUniqueExperiments(state, experiments) {
      state.uniqueExperiments = experiments;
    },
  },

  actions: {
    addUniquePatients({ commit }, patients) {
      commit('addUniquePatients', patients);
    },

    addUniqueSamples({ commit }, samples) {
      commit('addUniqueSamples', samples);
    },

    addUniqueExperiments({ commit }, experiments) {
      commit('addUniqueExperiments', experiments);
    },
  },

  getters: {
    patients: (state) => {
      return state.uniquePatients;
    },

    samples: (state) => {
      return state.uniqueSamples;
    },

    experiments: (state) => {
      return state.uniqueExperiments;
    }
  }
})
