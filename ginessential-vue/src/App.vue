<template>
  <div id="app">
    <!-- <nav>
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </nav> -->
    <NavBar />
    <!-- 页面主体放入container中 -->
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import NavBar from './views/layout/NavBar.vue';

export default {
  components: { NavBar },
  data() {},
};
</script>
<style lang="scss"></style>
