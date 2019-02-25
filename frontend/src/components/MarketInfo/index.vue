<template>
  <div class="marketInfo-box">
    <img class="my-img" src="/static/img/sixTri.png" alt="">
    <div class="my-marketInfo">
      <img class="marketInfo-logo" src="/static/img/iostMarket.png" alt="">
      <div class="marketInfo-price" v-if="marketInfo.price && marketInfo.btcPrice">
        <p><span>$</span> {{ marketInfo.price }}</p>
        <p>{{marketInfo.btcPrice}} BTC / {{marketInfo.ethPrice}} ETH</p>
      </div>
      <div class="marketInfo-price" v-else>
        <img src="/static/img/loading.gif" alt="">
      </div>

      <div class="marketInfo-time" v-if="marketInfo.volume24h && marketInfo.marketCap && marketInfo.lastUpdate">
        <div class="time-name">
          <!-- <p>24h Change:</p> -->
          <p>24h Volume:</p>
          <p>Market Cap:</p>
          <p>update:</p>
        </div>
        <div class="time-value">
          <!-- <p :class="[marketInfo.percentChange24h > 0 ? 'up': 'down']">{{ marketInfo.percentChange24h }}%</p> -->
          <p>$ {{ marketInfo.volume24h }}</p>
          <p>$ {{ marketInfo.marketCap }}</p>
          <p>{{ marketInfo.lastUpdate }}</p>
        </div>
      </div>
      <div class="marketInfo-time" v-else>
        <img src="/static/img/loading.gif" alt="">
      </div>
    </div>
  </div>
</template>

<script>
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
          /*font-weight: bold;*/
          &:nth-child(1) {
            margin-top: 45px;
            font-size: 72px;
            line-height: 1;
            color: #000000;
            font-weight: 300;
            span {
              /*font-size: 18px;*/
              font-size: 14px;
            }
          }
          &:nth-child(2) {
            margin-top: 29px;
            /*font-size: 18px;*/
            font-size: 14px;
            line-height: 22px;
            color: #2c2e31;
          }
        }

        > img {
          width: 310px;
          height: 200px;
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
            font-weight: bold;
            &:last-child {
              margin-bottom: 0;
              font-weight: normal;
            }
          }
        }
        .time-name {
          text-align: right;
        }
        .time-value {
          .up {
            color: #49C5B6;
          }
          .down {
            color: #FF4D4D;
          }
        }
        > img {
          width: 310px;
          height: 200px;
          margin-top: -55px;
          margin-left: -55px;
        }
      }
    }

  }

  @media screen and (max-width:480px) {
    .marketInfo-box {
      width: 100%;
      height: auto;
      margin: 30px auto 0;
      .my-img {
        display: none;
      }
      .my-marketInfo {
        flex-direction: column;
        align-items: center;
        border: 1px solid #F6F7F8;
        margin: 0 10px;
        box-shadow: 0 10px 40px 0 rgba(0, 0, 0, .1);
        .marketInfo-logo {
          display: none;
        }
        .marketInfo-price {
          margin-left: 0;
          p {
            text-align: center;
            &:nth-child(1) {
              margin-top: 24px;
            }
            &:nth-child(2) {
              margin-top: 24px;
            }
          }
        }
        .marketInfo-time {
          margin-left: 0;
          margin-bottom: 24px;
          margin-top: 24px;
        }
      }

    }
  }
</style>
