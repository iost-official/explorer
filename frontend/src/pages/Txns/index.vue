<template>
  <div class="txns-box">
    <LuckyBet/>

    <div class="txns-header">
      <div class="my-header-container">
        <div class="txns-info">
          <div>
            <h1 v-if="address != ''">
              Transactions
              <p>Address：{{address}}</p>
            </h1>
            <h1 v-else-if="blk != ''">
              Transactions
              <p>Block：{{blk}}</p>
            </h1>
            <h1 v-else>
              Transactions
            </h1>
          </div>
          <h4>{{txnInfo.totalLen}} transactions found (showing the last 500 records)</h4>
        </div>
        <ul class="my-pages">
          <li>
            <router-link :to="{path:`/txs?a=${address}&b=${blk}`}">
              <span><img src="../../assets/arrow-douleft.png" alt=""></span>
            </router-link>
          </li>

          <li>
            <a v-if="page == 1" href="javascript:void(0)">
              <span><img src="../../assets/arrow-left.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/txs?p=${(page-1)}&a=${address}&b=${blk}`}">
              <span><img src="../../assets/arrow-left.png" alt=""></span>
            </router-link>
          </li>

          <li><a class="page-auto" href="#"><b>{{page}}</b> / <b>{{txnInfo.pageLast}}</b></a></li>
          <li>
            <a v-if="page == txnInfo.pageLast" href="javascript:void(0)">
              <span><img src="../../assets/arrow-right.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/txs?p=${(page+1)}&a=${address}&b=${blk}`}">
              <span><img src="../../assets/arrow-right.png" alt=""></span>
            </router-link>
          </li>
          <li>
            <router-link :to="{path:`/txs?p=${txnInfo.pageLast}&a=${address}&b=${blk}`}">
              <span><img src="../../assets/arrow-douright.png" alt=""></span>
            </router-link>
          </li>
        </ul>
      </div>
    </div>

    <div class="txns-list">
      <ul class="my-list-header">
        <li>Height</li>
        <li>TxHash</li>
        <li>From</li>
        <li></li>
        <li>To</li>
        <li>Age</li>
        <li>Value</li>
      </ul>
      <div class="list-wrap">
        <div class="list-body-wrap" v-for="tx in txnInfo.txsList">
          <ul class="my-list-body">
            <li>
              <router-link :to="{path:`/block/${tx.blockHeight}`}">{{tx.blockHeight}}</router-link>
            </li>
            <li>
              <router-link :to="{path:`/tx/${tx.txHash}`}">{{tx.txHash}}</router-link>
            </li>
            <li>
              <span v-if="address == tx.from">{{tx.from}}</span>
              <router-link v-else :to="{path:`/account/${tx.from}`}">{{tx.from}}</router-link>
            </li>
            <li>
              <img v-if="tx.state == 0" src="../../assets/failure.png" alt="">
              <img v-else src="../../assets/success.png" alt="">
            </li>
            <li>
              <router-link :to="{path:`/account/${tx.to}`}">{{tx.to}}...</router-link>
            </li>
            <li>{{tx.age}}</li>
            <li>{{tx.amount}} IOST</li>
          </ul>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
  import LuckyBet from '../../components/LuckyBet'
  import { mapState } from 'vuex'

  export default {
    name: "Txns",
    data() {
      return {
        page: '',
        address: '',
        blk: '',
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

        this.$store.dispatch('getTxnInfo',{page:this.page, address: this.address, blk:this.blk})
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
    },

    components: {
      LuckyBet
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .txns-box {
    padding-bottom: 100px;
    .txns-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        height: 120px;
        width: 1000px;
        margin: 0 auto;
        display: flex;
        justify-content: space-between;
        text-align: left;
        .txns-info {
          display: flex;
          flex-direction: column;
          justify-content: center;
          padding: 20px 0 22px;
          font-size: 0;
          h1 {
            margin: 0;
            padding: 0;
            font-size: 36px;
            line-height: 44px;
            font-weight: bold;
            height: 44px;
            > p {
              /*font-size: 12px;*/
              font-size: 14px;
              color: #999;
              font-weight: 400;
              line-height: 1;
              height: 14px;
            }

          }
          h4 {
            margin: 15px 0 0;
            font-size: 14px;
            line-height: 18px;
            color: #2C2E31;
          }
        }

        .my-pages {
          display: inline-block;
          margin: 64px 0 0;
          padding: 0;
          li {
            list-style: none;
            display: inline;
            a {
              padding: 5px;
              height: 32px;
              color: #2c2e31;
              text-align: center;
              position: relative;
              float: left;
              margin-left: -1px;
              line-height: 1.42857143;
              text-decoration: none;
              background-color: #fff;
              border: 1px solid #ddd;
              &:hover, &:focus {
                z-index: 2;
                color: #23527c;
                background-color: #eee;
                border-color: #ddd;
              }
              &.page-auto {
                width: 112px;
              }
              span {
                padding: 0;
                img {
                  width: 20px;
                  height: 20px;
                }
              }
            }
          }
        }

      }

    }

    .txns-list {
      min-height: 500px;
      margin-top: 2px;
      &::-webkit-scrollbar {
        width: 0;
      }

      .my-list-header {
        /*font-size: 12px;*/
        font-size: 14px;
        line-height: 15px;
        font-weight: 500;
        > li {
          color: #2C2E31;
        }
      }
      .list-body-wrap {
        width: 100%;
        &:hover {
          box-shadow: 0 2px 10px 0 rgba(0, 0, 0, .08);
        }
        .list-wrap {
          max-height: 1000px;
          overflow: auto;
          .my-list-body {

          }
        }
      }

      ul {
        display: flex;
        padding: 24px 0 21px;
        width: 1000px;
        margin: 0 auto;
        font-size: 14px;
        line-height: 18px;
        border-bottom: 1px solid #f6f7f8;
        li {
          list-style: none;
          text-align: left;
          color: #2c2e31;
          a {
            color: #4b78aa;
            /*font-size: 12px;*/
            font-size: 14px;
            line-height: 15px;
          }
          &:first-child {
            width: 80px;
          }
          &:nth-child(2) {
            width: 220px;
            padding-right: 30px;
            overflow: hidden;
            text-overflow: ellipsis;
          }
          &:nth-child(3) {
            width: 180px;
            padding-right: 0;
            overflow: hidden;
            text-overflow: ellipsis;
          }
          &:nth-child(4) {
            width: 80px;
            text-align: center;
            > img {
              width: 18px;
              height: 18px;
            }
          }
          &:nth-child(5) {
            width: 200px;
            padding-right: 30px;
            overflow: hidden;
            text-overflow: ellipsis;
          }
          &:nth-child(6) {
            width: 130px;
          }

          &:nth-child(7) {
            width: 60px;
          }
        }
      }
    }

  }

  @media screen and (max-width:480px) {
    .txns-box {
      .txns-header {
        .my-header-container {
          height: auto;
          width: 100%;
          flex-direction: column;
          padding: 0 25px;
          .txns-info {
            h1 {
              height: auto;
            }
          }
          .my-pages {
            margin: 0;
          }
        }
      }
      .txns-list {
        overflow: auto;
        padding: 0 25px;
      }
    }
  }
</style>
