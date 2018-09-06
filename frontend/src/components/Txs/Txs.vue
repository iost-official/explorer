<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 v-if="address != ''" class="pull-left" style="display: block">
            Transactions <span style="font-size: 16px; color: #999; font-weight: 400">:: Address {{address}}</span>
          </h1>
          <h1 v-else-if="blk != ''" class="pull-left" style="display: block;">
            Transactions <span style="font-size: 16px; color: #999; font-weight: 400">:: Block {{blk}}</span>
          </h1>
          <h1 v-else class="pull-left" style="display: block">
            Transactions <span style="font-size: 16px; color: #999; font-weight: 400"></span>
          </h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Transactions</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" id="banner-luckybet">
        <i class="far fa-calendar-alt"></i> Latest Activity: <a href="/#/luckyBet">Play Lucky Bet !</a>
      </div>
    </div>

    <div class="container">
      <div class="row explorer-brief">
        <h4 class="pull-left">{{txnInfo.totalLen}} transactions found (showing the last 500 records)</h4>
        <nav aria-label="..." class="pull-right">
          <ul class="pagination">
            <li><a :href="'/#/txs?a=' + address + '&b=' + blk" aria-label="First"><span aria-hidden="true">«</span></a></li>

            <li>
              <a v-if="page == 1" href="javascript:void(0)" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
              <a v-else :href="'/#/txs?p=' + (page-1) + '&a=' + address + '&b=' + blk" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
            </li>

            <li><a href="#">page <b>{{page}}</b> of <b>{{txnInfo.pageLast}}</b></a></li>
            <li>
              <a v-if="page == txnInfo.pageLast" href="javascript:void(0)" aria-label="Next">
                <span aria-hidden="true">></span>
              </a>
              <a v-else :href="'/#/txs?p=' + (page+1) + '&a=' + address + '&b=' + blk">
                <span aria-hidden="true">></span>
              </a>
            </li>
            <li><a :href="'/#/txs?p=' + txnInfo.pageLast + '&a=' + address + '&b=' + blk" aria-label="Last"><span aria-hidden="true">»</span></a></li>
          </ul>
        </nav>
      </div>
      <div class="row">
        <table class="table explorer-table txs-table">
          <thead>
          <tr>
            <th>TxHash</th>
            <th>Block</th>
            <th>Age</th>
            <th>From</th>
            <th></th>
            <th>To</th>
            <th>Value</th>
          </tr>
          </thead>
          <tbody>
            <tr v-for="tx in txnInfo.txsList">
              <td><a :href="'/#/tx/' + tx.txHash">{{tx.txHash}}</a></td>
              <td><a :href="'/#/block/' + tx.blockHeight">{{tx.blockHeight}}</a> </td>
              <td>{{tx.age}}</td>
              <td>
                <span v-if="address == tx.from">{{tx.from}}</span>
                <a v-else :href="'/#/account/' + tx.from">{{tx.from}}</a>
              </td>
              <td>
                <span v-if="address != ''" class="label label-success rounded">&nbsp; IN &nbsp;</span>
                <i v-else class="fas fa-arrow-right" style="color:#0f0"></i>
              </td>
              <td><a :href="'/#/account/' + tx.to">{{tx.to}}...</a></td>
              <td>{{tx.amount}} IOST</td>
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
    name: "Txs",
    data() {
      return {
        // txnList: [],
        page: '',
        // totalPage: '',
        address: '',
        blk: '',
        // totalLen: '',
      }
    },
    methods: {
      fetchData(r) {
        if (!r.query.p) {
          this.page = 1
        } else {
          this.page = parseInt(r.query.p)
        }

        this.address = ''
        if (this.$route.query.a) {
          this.address = this.$route.query.a
        }
        this.blk = ''
        if (this.$route.query.b) {
          this.blk = this.$route.query.b
        }


        this.$store.dispatch('getTxnInfo',this.page, this.address, this.blk)

        // axios.get('https://explorer.iost.io/api/txs?p=' + this.page + '&a=' + this.address + '&b=' + this.blk).then((response) => {
        //   this.txnList = response.data.txs_list
        //   this.totalPage = response.data.page_last
        //   this.totalLen = response.data.total_len
        // })
      }
    },

    computed: {
      ...mapState(['txnInfo'])
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
