<template>
  <div class="search-box">
    <div class="luckyBet-box">
      <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
    </div>
    <div class="search-header">
      <div class="my-header-container">
        <h1>Search</h1>
      </div>
    </div>

    <div class="my-container">
      <div class="err-search">
        <span>Oops!</span>
        <p>The search string you entered was: </p>
        <input class="my-input" v-model="searchInput" readonly>
        <p>Sorry! This is an invalid search string.</p>
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
    padding-top: 90px;
    margin: 0 auto;
    .luckyBet-box {
      background: #2C2E31;
      height: 50px;
      line-height: 50px;
      color: #F6F7F8;
      font-size: 14px;
      > img {
        width: 24px;
        height: 24px;
        margin-right: 12px;
      }
      a {
        color: #F6F7F8;
        font-size: 14px;
        line-height: 18px;
        text-decoration: none;
      }
    }

    .search-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        display: flex;
        justify-content: space-between;
        height: 85px;
        width: 1000px;
        margin: 0 auto;
        > h1 {
          margin-top: 21px;
          font-size: 36px;
          line-height: 44px;
          font-weight: bold;
        }
      }
    }

    .my-container {
      width: 1000px;
      margin: 0 auto;
      padding: 100px 0;
      min-height: 650px;
      display: flex;
      flex-direction: column;
      align-items: center;
      .my-input {
        display: block;
        width: 700px;
        background-color: #F6F7F8;
        outline: none;
        border: 0;
        padding: 6px 12px;
        margin: 20px 0;
        color: #555;
        text-align: center;

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
