<template>
  <div class="luckyBet-box">
    <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
  </div>
</template>

<script>
  export default {
    name: "LuckyBet"
  }
</script>

<style lang="less" rel="stylesheet/less" scoped>
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
</style>