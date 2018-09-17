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
          <div class="list-block">
            <img src="../../assets/txns.png" alt="">
            <p>Tx</p>
          </div>
          <div class="list-tx">
            <p>TX#</p>
            <router-link :to="{path:`/tx/${txn.txHash}`}">{{txn.txHash}}</router-link>
          </div>
          <div class="list-from">
            <p>From</p>
            <router-link :to="{path:`/account/${txn.from}`}">{{txn.from}}</router-link>
          </div>
          <div class="list-to">
            <p>To</p>
            <router-link :to="{path:`/account/${txn.to}`}">{{txn.to}}</router-link>
          </div>
          <div class="list-amount">
            <p>Amount</p>
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
    background-color: #FFFFFF;
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
        color: #4b78aa;
      }
    }
    .my-list-body {
      padding-left: 0;
      padding-top: 8px;
      margin-bottom: 0;
      padding-bottom: 10px;

      li {
        list-style: none;
        padding: 16px 0 22px 30px;
        &:nth-child(2n) {
          background: #F6F7F8;
          .my-list-wrap {
            .list-block {
              border-bottom: 1px solid #FFFFFF;
            }
          }
        }
        &:nth-child(2n+1) {
          background: #FFFFFF;
          .my-list-wrap {
            .list-block {
              border-bottom: 1px solid #f6f7f8;
            }
          }
        }
        .my-list-wrap {
          text-align: left;
          > div {
            display: flex;
            align-items: center;
            &.list-block {
              padding-bottom: 6px;
              > img {
                width: 24px;
                height: 24px;
              }
              > p {
                font-size: 14px;
                line-height: 18px;
                font-weight: bold;
                color: #2C2E31;
                margin-top: 4px;
                margin-bottom: 0;
                padding-left: 26px;
              }
            }
            &.list-tx {
              padding-top: 9px;
            }
            &.list-tx, &.list-from, &.list-to {
              justify-content: space-between;
              padding-bottom: 14px;
              padding-right: 41px;
              > a {
                font-size: 12px;
                line-height: 15px;
                color: #4b78aa;
                display: inline-block;
                text-overflow: ellipsis;
                overflow: hidden;
                width: 360px;
              }
              > p {
                font-size: 12px;
                line-height: 15px;
                color: #2C2E31;
                margin-bottom: 0;
              }
            }
            &.list-amount {
              > p {
                font-size: 12px;
                line-height: 15px;
                color: #2C2E31;
                margin-bottom: 0;
                &:nth-child(2) {
                  margin-left: 4px;
                }
              }
            }
          }
        }
      }

    }
  }
  @media screen and (max-width:480px) {
    .hometxns-box {
      width: 100%;
    }
  }
</style>
