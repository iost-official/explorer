<template>
  <div class="marketInfo-box">
    <img class="my-img" src="../../assets/sixTri.png" alt="">
    <div class="my-marketInfo">
      <img class="marketInfo-logo" src="../../assets/iostMarket.png" alt="">
      <div class="marketInfo-price">
        <p><span>$</span> {{ marketInfo.price }}</p>
        <p>{{marketInfo.btcPrice}} BTC / {{marketInfo.ethPrice}} ETH</p>
      </div>
      <div class="marketInfo-time">
        <div class="time-name">
          <p>24h Change:</p>
          <p>24h Volume:</p>
          <p>Market Cap:</p>
          <p>update:</p>
        </div>
        <div class="time-value">
          <p :class="[marketInfo.percentChange24h > 0 ? 'up': 'down']">{{ marketInfo.percentChange24h }}%</p>
          <p>$ {{ marketInfo.volume24h }}</p>
          <p>$ {{ marketInfo.marketCap }}</p>
          <p>{{ marketInfo.lastUpdate }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'


  export default {
    name: "marketInfo",
    data() {
      return {

      }
    },

    computed: {
      ...mapState(['marketInfo'])
    },

    mounted () {
      this.$store.dispatch('getMarketInfo')
    }
  }
</script>

<style lang="less" rel="stylesheet/less" scoped>
  .marketInfo-box {
    width: 1000px;
    height: 220px;
    margin: 100px auto 0;
    .my-img {
      width: 1080px;
      height: 300px;
      position: absolute;
      left: 0;
      right: 0;
      margin: 0 auto;
      top: 285px;
    }
    .my-marketInfo {
      display: flex;
      position: relative;
      z-index: 10;
      .marketInfo-logo {
        margin-top: 36px;
        margin-left: 24px;
        width: 143px;
        height: 142px;
      }
      .marketInfo-price {
        margin-left: 62px;
        p {
          text-align: left;
          margin-bottom: 0;
          &:nth-child(1) {
            margin-top: 45px;
            font-size: 72px;
            line-height: 1;
            color: #000000;
            font-weight: 300;
            span {
              font-size: 18px;
            }
          }
          &:nth-child(2) {
            margin-top: 29px;
            font-size: 18px;
            line-height: 22px;
            color: #4C6889;
          }
        }
      }
      .marketInfo-time {
        width: 200px;
        margin-left: 130px;
        margin-top: 55px;
        display: flex;
        justify-content: space-between;
        text-align: left;
        font-size: 14px;
        line-height: 18px;
        color: #2C2E31;
        .time-name, .time-value {
          > p {
            margin-bottom: 15px;
            &:last-child {
              margin-bottom: 0;
            }
          }
        }
        .time-name {
          text-align: right;
        }
        .time-value {
          .up {
            color: #0f0;
          }
          .down {
            color: #f00;
          }
        }
      }
    }

  }
</style>
