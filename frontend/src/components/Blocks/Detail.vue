<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Block <span style="font-weight: 200; font-size: 25px">#{{blockHeight}}</span></h1>
          <ol class="breadcrumb pull-right">
            <li><a href="#">Home</a></li>
            <li class=""><a :href="'#/blocks'">Blocks</a></li>
            <li class="active">Block Information</li>
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
        <table class="table" style="text-align: left">
          <thead>
          <tr>
            <th colspan="2">Block Information</th>
          </tr>
          </thead>
          <tbody>
          <tr>
            <td>Height:</td>
            <td>{{blockHeight}}</td>
          </tr>
          <tr>
            <td>TimeStamp:</td>
            <td>{{blockDetail.age}} ({{blockDetail.utcTime}})</td>
          </tr>
          <tr>
            <td>Transactions:</td>
            <td><a :href="'/#/txs/?b=' + blockHeight">{{blockDetail.txn}}</a></td>
          </tr>
          <tr>
            <td>Hash:</td>
            <td>{{blockDetail.blockHash}}</td>
          </tr>
          <tr>
            <td>Parent Hash:</td>
            <td>{{blockDetail.parentHash}}</td>
          </tr>
          <tr>
            <td>Witness:</td>
            <td><a :href="'/#/account/' + blockDetail.witness">{{blockDetail.witness}}</a></td>
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
    name: "Block",
    data() {
      return {
        blockHeight: '',
        // age: '',
        // utcAge: '',
        // txnLen: '',
        // txnHash: '',
        // parentHash: '',
        // blockHash: '',
        // witness: '',
      }
    },
    methods: {
      fetchData(r) {
        this.blockHeight = r.params.id
        this.$store.dispatch('getBlockDetail', this.blockHeight)

        // axios.get('https://explorer.iost.io/api/block/' + this.blockHeight).then((response) => {
        //   this.blockHeight = response.data.height
        //   this.age = response.data.age
        //   this.utcAge = response.data.utc_time
        //   this.txnLen = response.data.txn
        //   this.parentHash = response.data.parent_hash
        //   this.blockHash = response.data.block_hash
        //   this.witness = response.data.witness
        // })
      }
    },

    computed: {
      ...mapState(['blockDetail'])
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
