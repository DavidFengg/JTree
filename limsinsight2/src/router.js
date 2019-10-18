import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Patients from './views/Patients.vue'
import Samples from './views/Samples.vue'
import Experiments from './views/Experiments.vue'
import Results from './views/Results.vue'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '',
      name: 'home',
      component: Home
    },
    {
      path: '/patients',
      name: 'patients',
      component: Patients
    },
    {
      path: '/samples',
      name: 'Samples',
      component: Samples
    },
    {
      path: '/experiments',
      name: 'Experiments',
      component: Experiments
    },
    {
      path: '/results',
      name: 'Results',
      component: Results
    }
  ]
})
