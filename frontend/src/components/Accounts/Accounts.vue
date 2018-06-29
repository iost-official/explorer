<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">All Accounts</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Accounts</li>
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
        <h4 v-if="totalPage * 25 < 500 " class="pull-left">{{totalLen}} accounts found</h4>
        <h4 v-else class="pull-left">{{totalLen}} accounts found (showing the last 500 records)</h4>
        <nav aria-label="..." class="pull-right">
          <ul class="pagination">
            <li><a href="/#/accounts" aria-label="First"><span aria-hidden="true">«</span></a></li>
            
            <li>
              <a v-if="page == 1" href="javascript:void(0)" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
              <a v-else :href="'/#/accounts?p=' + (page-1)" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
            </li>

            <li><a href="#">page <b>{{page}}</b> of <b>{{totalPage}}</b></a></li>
            <li>
              <a v-if="page == totalPage" href="javascript:void(0)" aria-label="Next">
                <span aria-hidden="true">></span>
              </a>
              <a v-else :href="'/#/accounts?p=' + (page+1)">
                <span aria-hidden="true">></span>
              </a>
            </li>
            
            <li><a :href="'/#/accounts?p=' + totalPage" aria-label="Last"><span aria-hidden="true">»</span></a></li>
          </ul>
        </nav>
      </div>
      <div class="row">
        <table class="table explorer-table">
          <thead>
          <tr>
            <th>Rank</th>
            <th>Address</th>
            <th>Balance</th>
            <th>TxCount</th>
          </tr>
          </thead>
          <tbody>
            <tr v-for="(account, index) in accountList">
              <td>{{index+1}}</td>
              <td><a :href="'/#/account/' + account.address">{{account.address}}</a></td>
              <td>{{account.balance.toFixed(2)}}</td>
              <td>{{account.tx_count}}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'Accounts',
    data() {
      return {
        accountList: [],
        page: '',
        totalPage: '',
        totalLen: '',
      }
    },
    methods: {
      fetchData(r) {
        if (!r.query.p) {
          this.page = 1
        } else {
          this.page = parseInt(r.query.p)
        }

        axios.get('https://explorer.iost.io/api/accounts?p=' + this.page).then((response) => {
          this.accountList = response.data.account_list
          this.totalPage = response.data.page_last
          this.totalLen = response.data.total_len
        })
      }
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
