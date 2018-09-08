<template>
  <div id="app">
    <router-view name="Head"></router-view>
    <router-view name="Main"></router-view>
    <router-view name="Foot"></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
  @font-face{
    font-family: 'Montserrat';
    src: url('./assets/font/Montserrat-Light.woff.ttf');
  }

  #app {
  font-family: Montserrat, "Helvetica Neue",Helvetica,Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
