<template>
  <div class="search-box">
    <div class="search-header">
      <div class="my-header-container">
        <h1>Search</h1>
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

<style lang="less" rel="stylesheet/less">
  .search-box {
    padding-top: 60px;

    .search-header {
      .my-header-container {
        width: 980px;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        > h1 {
          margin: 10px 0;
        }
      }
    }
  }
  .err-search span {
    color: #555;
    display: block;
    font-size: 35px;
    font-weight: 200;
  }
</style>
