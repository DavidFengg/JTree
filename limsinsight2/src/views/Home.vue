<template>
  <div class="home">
    <!-- <NavBar/> -->
    <HelloWorld msg="Test"/>
  </div>
</template>

<script>
// @ is an alias to /src
// import NavBar from '@/components/NavBar.vue'
import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'home',
  components: {
    HelloWorld,
    // NavBar
  }
}
</script>
