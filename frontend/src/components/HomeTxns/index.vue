<template>
  <div class="hometxns-box">
    <div class="my-list-banner">
      <div class="my-list-title">
        TRANSACTIONS
      </div>
      <router-link to="/txs">View All</router-link >
    </div>
    <ul class="my-list-body">
      <li v-for="(txn, index) in indexTxnList" :key="index">
        <div class="my-list-wrap">
          <div class="my-list1">
            <p><img src="../../assets/txns.png" alt=""></p>
            <p>TX#</p>
            <p>From</p>
            <p>To</p>
            <p>Amount</p>
          </div>
          <div class="my-list2">
            <p>Block {{txn.blockHeight}}</p>
            <router-link :to="{path:`/tx/${txn.txHash}`}">{{txn.txHash}}</router-link>
            <router-link :to="{path:`/account/${txn.from}`}">{{txn.from}}</router-link>
            <router-link :to="{path:`/account/${txn.to}`}">{{txn.to}}</router-link>
            <p>{{txn.amount}} IOST</p>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'


  export default {
    name: "IndexTxns",
    data() {
      return {
        // indexTxnList: [],
      }
    },
    // mounted: function () {
    //   axios.get('https://explorer.iost.io/api/indexTxns').then((response) => {
    //
    //     this.indexTxnList = response.data
    //   })
    // }

    computed: {
      ...mapState(['indexTxnList'])
    },

    mounted () {
      this.$store.dispatch('getIndexTxnList')
    }
  }
</script>

<style lang="less" rel="stylesheet/less">

  .hometxns-box {
    width: 485px;
    border: 1px solid #F1F1F1;
    box-shadow: 0 10px 40px 0 rgba(0, 0, 0, .1);
    .my-list-banner {
      background: #F6F7F8;
      padding: 0 30px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      height: 54px;
      box-shadow: 0 2px 3px rgba(0,0,0,0.1);
      .my-list-title {
        font-size: 20px;
        line-height: 24px;
        color: #2C2E31;
        font-weight: bolder;
      }
      > a{
        font-size: 14px;
        line-height: 18px;
        color: #4C6889;
        text-decoration: none;
      }
    }
    .my-list-body {
      padding-left: 0;
      margin-top: 8px;
      li {
        list-style: none;
        padding: 16px 0 16px 30px;
        height: 160px;
        &:nth-child(2n) {
          background: #F6F7F8;
        }
        .my-list-wrap {
          display: flex;
          text-align: left;
          .my-list1 {
            > p {
              font-size: 12px;
              line-height: 15px;
              color: #2C2E31;
              margin-top: 8px;
              margin-bottom: 0;
              > img {
                width: 24px;
                height: 24px;
              }
              &:first-child {
                margin-top: 0;
              }
              &:nth-child(2) {
                margin-top: 18px;
              }
              &:nth-child(3) {
                margin-top: 8px;
              }
              &:nth-of-type(4) {
                margin-top: 8px;
              }
            }
          }
          .my-list2 {
            margin-left: 10px;
            font-size: 0;
            > p {
              font-size: 14px;
              line-height: 18px;
              font-weight: bold;
              color: #4C6889;
              margin-top: 4px;
              margin-bottom: 0;
              &:nth-of-type(2) {
                margin-top: 7px;
                font-size: 12px;
                font-weight: 400;
              }
            }
            > a {
              font-size: 12px;
              line-height: 15px;
              color: #4C6889;
              display: inline-block;
              text-overflow: ellipsis;
              overflow: hidden;
              max-width: 340px;
              &:nth-of-type(1) {
                margin-top: 20px;
              }
              &:nth-of-type(2),&:nth-of-type(3) {
                margin-top: 8px;
              }
            }
          }
        }
      }

    }
  }
</style>
