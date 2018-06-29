<template>
  <div class="txns">
    <div class="list-banner">
      <div class="list-title pull-left">
        <i class="fa fa-list-alt"></i> Transactions
      </div>
      <a href="/#/txs" class="pull-right">View All</a>
    </div>
    <ul class="list-body txn-list-body">
      <li v-for="(txn, index) in txnList" class="transaction-detail" :class="'rainbow-border-' + index">
        <table class="table">
          <tr>
            <td><i class="fa fa-tasks"></i></td>
            <td>TX#</td>
            <td><a :href="'#/tx/' + txn.tx_hash">{{txn.tx_hash}}</a></td>
          </tr>
          <tr>
            <td></td>
            <td>From</td>
            <td><a :href="'#/account/' + txn.from">{{txn.from}}</a></td>
          </tr>
          <tr>
            <td></td>
            <td>To</td>
            <td><a :href="'#/account/' + txn.to">{{txn.to}}</a></td>
          </tr>
          <tr>
            <td></td>
            <td>Amount</td>
            <td>{{txn.amount}} IOST</td>
          </tr>
        </table>
      </li>
    </ul>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "IndexTxns",
    data() {
      return {
        txnList: [],
      }
    },
    mounted: function () {
      axios.get('https://explorer.iost.io/api/indexTxns').then((response) => {
        this.txnList = response.data
      })
    }
  }
</script>

<style>
  .txn-list-body {
    height: 1000px;
  }

  .transaction-detail {
    padding-top: 0px !important;
  }

  .transaction-detail i {
    width: 30px;
  }

  .transaction-detail tr > td:first-child {
    width: 40px;
    height: 25px;
    text-align: center;
  }

  .transaction-detail tr > td:nth-child(2) {
    width: 70px;
  }

  .transaction-detail tr > td > a {
    width: 450px;
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;
  }
</style>
