<template>
  <div class="marketInfo">
    <p>$ {{ marketInfo.price }}</p>
    <p>{{marketInfo.btcPrice}} BTC / {{marketInfo.ethPrice}} ETH</p>
    <!--<p><span>24h Change:</span><span :class="{up: isUp, down: isDown}">{{ marketInfo.percentChange24h }}%</span></p>-->
    <p><span>24h Change:</span><span :class="[marketInfo.percentChange24h > 0 ? 'up': 'down']">{{ marketInfo.percentChange24h }}%</span></p>
    <p><span>24h Volume:</span><span>$ {{ marketInfo.volume24h }}</span></p>
    <p><span>Market Cap:</span><span>$ {{ marketInfo.marketCap }}</span></p>
    <p>update {{ marketInfo.lastUpdate }}</p>
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
    // mounted: function () {
    //   axios.get('https://explorer.iost.io/api/market').then((response) => {
    //
    //       var marketInfo = response.data
    //     console.log(marketInfo)
    //     this.price = marketInfo.price
    //     this.change = marketInfo.percent_change_24h
    //     this.volume = marketInfo.volume_24h
    //     this.cap = marketInfo.market_cap
    //     this.btcPrice = marketInfo.btc_price
    //     this.ethPrice = marketInfo.eth_price
    //     this.update = marketInfo.last_update
    //     if (this.change > 0) {
    //       this.isUp = true
    //     } else {
    //       this.isDown = true
    //     }
    //
    //     this.$emit('marketRender', 'done')
    //   })
    // }


    computed: {
      ...mapState(['marketInfo'])
    },

    mounted () {
      this.$store.dispatch('getMarketInfo')
    }
  }
</script>

<style>
  .marketInfo p {
    text-align: left;
    color: #000;
  }

  .marketInfo p:first-child {
    font-size: 50px;
    font-weight: 500;
    line-height: 60px;
    border-bottom: 1px solid #f7f7f7;
    margin-bottom: 0px;
  }

  .marketInfo p:nth-child(2) {
    border-bottom: 1px solid #f7f7f7;
    font-size: 20px;
    line-height: 40px;
  }

  .marketInfo p:nth-child(n+3) {
    font-size: 16px;
    margin-top: 10px;
    margin-bottom: 0;
  }

  .marketInfo p:last-child {
    text-align: right;
    color: #a7a7a7;
    font-size: 14px;
    margin-top: 15px;
  }

  .marketInfo span:last-child {
    float: right;
  }

  .up {
    color: #0f0;
  }

  .down {
    color: #f00;
  }
</style>
