<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Address <span style="font-weight: 200; font-size: 20px;">{{address}}</span></h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class=""><a href="/#/accounts">Accounts</a></li>
            <li class="active">Address</li>
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
        <table class="table account-overview">
          <thead>
          <tr>
            <th colspan="2">Overview</th>
          </tr>
          </thead>
          <tbody>
          <tr>
            <td>Balance:</td>
            <td>{{accountDetail.balance}}</td>
          </tr>
          <tr>
            <td>IOST Value:</td>
            <td></td>
          </tr>
          <tr>
            <td>Transactions:</td>
            <td>{{txnLen}}</td>
          </tr>
          </tbody>
        </table>
      </div>
      <div class="row">
        <ul class="nav nav-tabs" id="txnDetailList">
          <li role="presentation" class="active"><a href="#txns-nav-txns">Transactions</a></li>
          <li id="CCodeDisplay" style="display:none" role="presentation"><a href="#txns-contract-code">Contract Code</a></li>
        </ul>
      </div>
      <div class="tab-content">
        <div class="tab-pane active row address-txns-head" id="txns-nav-txns">
          <div style="line-height: 80px">
            <div class="col-md-8">
              <i aria-hidden="true" class="c333 fa fa-sort-amount-down"></i> Latest 25 txns from a total Of <a :href="'/#/txs?a=' + address">{{txnLen}} transactions</a>
            </div>
            <div class="col-md-4">
              <a :href="'/#/txs?a=' + address">View All</a>
            </div>
          </div>
          <table class="table account-tx-table">
            <thead>
            <tr>
              <th>TxHash</th>
              <th>Block</th>
              <th>Age</th>
              <th>From</th>
              <th>To</th>
              <th>Value</th>
              <th>[TxFee]</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="tx in accountTxnInfo.txnList">
              <td><a :href="'/#/tx/' + tx.hash">{{tx.hash}}...</a></td>
              <td><a :href="'/#/block/' + tx.blockNumber">{{tx.blockNumber}}</a></td>
              <td>{{tx.age}}</td>
              <td><a :href="'/#/account/' + tx.from">{{tx.from}}</a></td>
              <td><a :href="'/#/account/' + tx.to">{{tx.to}}</a></td>
              <td>{{tx.amount}}</td>
              <td>{{tx.gasPrice}}</td>
            </tr>
            </tbody>
          </table>
        </div>

        <div class="tab-pane" id="txns-contract-code">
          <table class="table ContractCodeTable">
            <tbody>
              <tr>
                <td>Publisher:</td>
                <td>{{ContractInfo.from}}</td>
              </tr>
              <tr>
                <td>Time:</td>
                <td>{{ContractInfo.time}}</td>
              </tr>
              <tr>
                <td>Code:</td>
                <td><pre>{{ContractInfo.code}}</pre></td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'

  export default {
    name: "AccountDetail",
    data() {
      return {
        address: this.$route.params.id,
        // account: {},
        txnList: {},
        txnLen: '',
        ContractInfo: {},
      }
    },

    computed: {
      ...mapState(['accountDetail', 'accountTxnInfo'])
    },

    methods: {

      fetchData(r) {
        this.address = r.params.id

        this.$store.dispatch('getAccountDetail', this.address)

        this.$store.dispatch('getAccountTxnInfo', this.address)


        // axios.get('https://explorer.iost.io/api/account/' + this.address).then((response) => {
        //   this.account = response.data
        // })
        // axios.get('https://explorer.iost.io/api/account/' + this.address + '/txs').then((response) => {
        //   this.txnList = response.data.txn_list
        //   this.txnLen = response.data.txn_len
        // })

        // axios.get('https://explorer.iost.io/api/tx/' + this.address).then((response) => {
        axios.get('http://47.75.223.44:8080/api/tx/' + this.address).then((response) => {
          if (response.data.code == 1) {
            $('#CCodeDisplay').hide()
            return
          }
          $('#CCodeDisplay').show()
          let blkDate = new Date(response.data.time / 1000 / 1000)
          let time = ''
          time += blkDate.getFullYear() + '-'
          time += (blkDate.getMonth()+1 < 10 ? '0'+(blkDate.getMonth()+1) : blkDate.getMonth()+1) + '-'
          time += blkDate.getDate() + ' ';
          if (blkDate.getHours() < 10) {
            time += '0' + blkDate.getHours()
          } else {
            time += blkDate.getHours()
          }
          if (blkDate.getMinutes() < 10) {
            time += ':0' + blkDate.getMinutes()
          } else {
            time += ':' + blkDate.getMinutes()
          }
          if (blkDate.getSeconds() < 10) {
            time += ':0' + blkDate.getSeconds()
          } else {
            time += ':' + blkDate.getSeconds()
          }
          response.data.time = time
          this.ContractInfo = response.data
        })
      }
    },



    watch: {
      '$route': function (r) {
        this.fetchData(r)
      }
    },
    mounted: function () {
      $("#txnDetailList a").click(function (e) {
        e.preventDefault()
        $(this).tab('show')
      })
      this.fetchData(this.$route)
    }
  }
</script>

<style>
  .ContractCodeTable {
    margin-top: 20px;
  }

  .ContractCodeTable tr:first-child td {
    border: 0;
  }
</style>
