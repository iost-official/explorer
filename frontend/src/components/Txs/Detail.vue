<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Transaction <span style="font-weight: 200; font-size: 16px">#{{txHash}}</span></h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class=""><a href="/#/txs/">Transactions</a></li>
            <li class="active">Transaction Information</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" id="banner-luckybet">
        <i class="far fa-calendar-alt"></i> Latest Activity: <a href="/#/luckyBet">Play Lucky Bet !</a>
      </div>
    </div>

    <div class="container block-information">
      <div class="row">
        <ul class="nav nav-tabs">
          <li role="presentation" class="active"><a href="#">Overview</a></li>
        </ul>
      </div>
      <div class="row">
        <table class="table">
          <thead>
          <tr>
            <th colspan="2">Transaction Information</th>
          </tr>
          </thead>
          <tbody>
          <tr>
            <td>TxHash:</td>
            <td>{{txHash}}</td>
          </tr>
          <tr>
            <td style="min-width: 110px;">Block Height:</td>
            <td><a :href="'/#/block/' + txnDetail.blockHeight">{{txnDetail.blockHeight}}</a></td>
          </tr>
          <tr>
            <td>TimeStamp:</td>
            <td>{{txnDetail.age}}({{txnDetail.utcTime}})</td>
          </tr>
          <tr>
            <td>From:</td>
            <td><a :href="'/#/account/' + txnDetail.from">{{txnDetail.from}}</a></td>
          </tr>
          <tr>
            <td>To:</td>
            <td><a :href="'/#/account/' + txnDetail.to">{{txnDetail.to}}</a></td>
          </tr>
          <tr>
            <td>Value:</td>
            <td>{{txnDetail.amount}} IOST</td>
          </tr>
          <tr>
            <td>Gas Limit:</td>
            <td>{{txnDetail.gasLimit}}</td>
          </tr>
          <tr>
            <td>Gas Price:</td>
            <td>{{txnDetail.price}}</td>
          </tr>
          <tr>
            <td>Code</td>
            <td><pre>{{txnDetail.code}}</pre></td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'

  export default {
    name: "Tx",
    data() {
      return {
        txHash: '',
        // blockHeight: '',
        // from: '',
        // to: '',
        // amount: '',
        // gasLimit: '',
        // price: '',
        // age: '',
        // utcTime: '',
        // code: ''
      }
    },
    methods: {
      fetchData(r) {
        this.txHash = r.params.id

        this.$store.dispatch('getTxnDetail', this.txHash)

        // axios.get('https://explorer.iost.io/api/tx/' + this.txHash).then((response) => {
        //   this.blockHeight = response.data.block_height
        //   this.from = response.data.from
        //   this.to = response.data.to
        //   this.amount = response.data.amount
        //   this.gasLimit = response.data.gas_limit
        //   this.price = response.data.price
        //   this.age = response.data.age
        //   this.utcTime = response.data.utc_time
        //   this.code = response.data.code
        // })
      }
    },

    computed: {
      ...mapState(['txnDetail'])
    },

    watch: {
      '$route': function (r) {
        this.fetchData(r)
      }
    },
    mounted: function () {
      this.fetchData(this.$route)
    }
  }
</script>
