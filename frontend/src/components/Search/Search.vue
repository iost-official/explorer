<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Search</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Search</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" style="margin-top: 100px; margin-bottom: 300px">
        <div class="col-md-8 col-md-offset-2 err-search">
          <span>Oops!</span>
          <br/>
          <p>The search string you entered was: </p>
          <pre>{{searchInput}}</pre>
          Sorry! This is an invalid search string.
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Search',
    data() {
      return {
        searchInput: this.$route.params.id,
      }
    },
    watch: {
      '$route': function (r) {
        this.searchInput = r.params.id
      }
    }
  }
</script>

<style>
  .err-search span {
    color: #555;
    display: block;
    font-size: 35px;
    font-weight: 200;
  }
</style>
