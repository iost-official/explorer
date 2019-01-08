<template>
  <div id="app">
    <router-view name="Head"></router-view>
    <div class="my-main">
      <router-view name="Main"></router-view>
    </div>
    <router-view name="Foot"></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style lang="less" rel="stylesheet/less">
  @font-face{
    font-family: 'Montserrat';
    src: url('/static/fonts/Montserrat-Light.woff.ttf');
  }
  #app {
    position: relative;
    height: 100%;
    font-family: Montserrat, "Helvetica Neue",Helvetica,Arial,sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    .my-main {
      padding-top: 90px;
      min-height: 884px;
    }
  }

</style>
