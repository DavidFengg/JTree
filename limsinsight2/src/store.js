import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    uniquePatients: [],
    uniqueSamples: [],
    uniqueExperiments: [],
    uniqueResults: []
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

    addUniqueResults(state, results) {
      state.uniqueResults = results;
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

    addUniqueResults({ commit }, results) {
      commit('addUniqueResults', results);
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
    },

    results: (state) => {
      return state.uniqueResults;
    }
  }
})
