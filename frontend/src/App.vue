<template>
  <div id="app">
    <router-view></router-view>
    <router-view name="Head"></router-view>
    <router-view name="Main"></router-view>
    <div class="container-fluid footer">
      <router-view name="Foot"></router-view>
    </div>
   <!--  <div class="modal fade bs-example-modal-sm" id="modal" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
      <div class="modal-dialog modal-sm" role="document">
        <div class="modal-content" style="border:0">
          <div class="progress progress-striped active">
            <div class="progress-bar progress-bar-success" role="progressbar"
               aria-valuenow="60" aria-valuemin="0" aria-valuemax="100"
               style="width: 100%;">
              <span class="sr-only" style="border:0">40% 完成</span>
            </div>
          </div>
        </div>
      </div>
    </div> -->
  </div>
</template>

<script>
// $(document).ready(function() {
//   $('#modal').modal('show')
// })
export default {
  name: 'app'
}
</script>

<style>
#app {
  font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
