<template>
  <div class="bg">
    <div class="home-box">
      <Search/>
      <MarketInfo/>
      <div class="blockTxns-container">
        <IndexBlock/>
        <IndexTxns/>
      </div>
    </div>
  </div>
</template>

<script>
  import MarketInfo from "../../components/MarketInfo";
  import IndexBlock from "../../components/HomeBlocks";
  import IndexTxns from "../../components/HomeTxns";
  import Search from '../../components/Search'

  import axios from 'axios';
  import { config } from '../../utils/config'

  const { apis } = config


  export default {
    name: "Home",

    components: {Search, IndexTxns, IndexBlock, MarketInfo},

    created () {
      axios.get(`${apis.test}`).then((response) => {
        if (true) {
          axios.get(`${apis.test2}`)
        }
      })
    }

  }
</script>

<style lang="less" rel="stylesheet/less">
  .bg {
    background-image: url("../../assets/bg.png");
    background-size: 100%;
    background-repeat: no-repeat;
    .home-box {
      width: 1040px;
      margin: 0 auto;
      overflow: hidden;
      .blockTxns-container {
        padding: 0 20px 100px;
        display: flex;
        justify-content: space-between;
        margin-top: 75px;
      }
    }
  }


</style>
