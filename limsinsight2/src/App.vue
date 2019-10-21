<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/patients">Patients</router-link> |
      <router-link to="/samples">Samples</router-link> |
      <router-link to="/experiments">Experiments</router-link> |
      <router-link to="/results">Results</router-link> |
      <router-link to="/resultdetails">Result Details</router-link> |

    </div>
    <router-view/>
  </div>
</template>

<script>
import APIService from './services/APIService';
 
export default {
  // load all data from tables upon creation
  mounted() {
    APIService.getPatients().then(data => {
      this.$store.dispatch('addUniquePatients', data);
    });

    APIService.getSamples().then(data => {
      this.$store.dispatch('addUniqueSamples', data);
    });

    APIService.getExperiments().then(data => {
      this.$store.dispatch('addUniqueExperiments', data);
    });

    APIService.getResults().then(data => {
      this.$store.dispatch('addUniqueResults', data);
    });
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
