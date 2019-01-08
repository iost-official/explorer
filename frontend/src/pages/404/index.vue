<template>
  <div class="container iost-404">
    <div class="dialog">
      <h1>Whoops!</h1>
      <p>We were unable to find the page you were looking for</p>
    </div>
    <img src="/static/img/404.png" class="image">
  </div>
</template>

<script>
  export default {
    name: 'NotFound404'
  }
</script>

<style>
  .dialog {
    float:right;
    text-align: left;
    width: 45%;
    margin: 5% auto 0;
  }

  .iost-404 h1 {
    font-size: 5em;
    color: #393939;
    line-height: 1.5em;
  }

  .iost-404 p {
    font-size: 1.4em;
    color: #393939;
    padding-right: 5%;
  }

  .iost-404 .image {
    float: left;
    width:50%;
    margin: 5% 0 0 5%;

  }
</style>
