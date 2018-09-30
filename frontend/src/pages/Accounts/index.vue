<template>
  <div class="accounts-box">
    <LuckyBet/>

    <div class="accounts-header">
      <div class="my-header-container">
        <div class="accounts-header-info">
          <h1>All Accounts</h1>
          <div class="accounts-info">
            <h4 v-if="accountInfo.pageLast * 25 < 500 ">{{accountInfo.totalLen}} accounts found</h4>
            <h4 v-else>{{accountInfo.totalLen}} accounts found (showing the last 500 records)</h4>
          </div>
        </div>
        <ul class="my-pages">
          <li>
            <router-link to="/accounts">
              <span><img src="/static/img/arrow-douleft.png" alt=""></span>
            </router-link>
          </li>

          <li>
            <a v-if="page == 1" href="javascript:void(0)">
              <span><img src="/static/img/arrow-left.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/accounts?p=${(page-1)}`}">
              <span><img src="/static/img/arrow-left.png" alt=""></span>
            </router-link>
          </li>

          <li><a href="#" class="page-auto"><b>{{page}}</b> / <b>{{accountInfo.pageLast}}</b></a></li>
          <li>
            <a v-if="page == accountInfo.pageLast" href="javascript:void(0)">
              <span><img src="/static/img/arrow-right.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/accounts?p=${(page+1)}`}">
              <span><img src="/static/img/arrow-right.png" alt=""></span>
            </router-link>
          </li>

          <li>
            <router-link :to="{path:`/accounts?p=${accountInfo.pageLast}`}">
              <span><img src="/static/img/arrow-douright.png" alt=""></span>
            </router-link>
          </li>
        </ul>
      </div>
    </div>

    <div class="accounts-list">
      <ul class="my-list-header">
        <li>Rank</li>
        <li>Address</li>
        <li>Balance</li>
        <li>TxCount</li>
      </ul>
      <div class="list-wrap">
        <div class="list-body-wrap" v-for="(account, index) in accountInfo.accountList">
          <ul class="my-list-body">
            <li>{{index+1}}</li>
            <li><router-link :to="{path:`/account/${account.address}`}">{{account.address}}</router-link></li>
            <li>{{(account.balance/100000000).toFixed(2)}}</li>
            <li>{{account.txCount}}</li>
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
    name: 'accounts',
    data() {
      return {
        page: '',
      }
    },
    methods: {
      fetchData(r) {
        if (!r.query.p) {
          this.page = 1
        } else {
          this.page = parseInt(r.query.p)
        }

        this.$store.dispatch('getAccountInfo',this.page)
      }
    },

    computed: {
      ...mapState(['accountInfo'])
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
  .accounts-box {
    padding-bottom: 100px;
    .accounts-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        display: flex;
        justify-content: space-between;
        height: 120px;
        width: 1000px;
        margin: 0 auto;
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
        .accounts-header-info {
          padding: 20px 0 22px;
          > h1 {
            font-size: 36px;
            line-height: 44px;
            font-weight: bold;
            margin: 0;

          }
          .accounts-info {
            margin-top: 15px;
            > h4 {
              text-align: left;
              color: #2C2E31;
              font-size: 14px;
              line-height: 18px;
              margin: 0;
            }
          }
        }
      }
    }


    .accounts-list {
      margin-top: 2px;
      &::-webkit-scrollbar {
        width: 0;
      }

      .my-list-header {
        font-size: 12px;
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
        }
      }

      ul {
        display: flex;
        padding: 24px 0 21px;
        margin: 0 auto;
        width: 1000px;
        border-bottom: 1px solid #f6f7f8;
        li {
          font-size: 14px;
          line-height: 18px;
          list-style: none;
          text-align: left;
          color: #2c2e31;
          a {
            font-size: 14px;
            line-height: 18px;
            color: #4b78aa;
            text-decoration: none;
          }
          &:first-child {
            width: 90px;
          }
          &:nth-child(2) {
            width: 560px;

          }
          &:nth-child(3) {
            width: 200px;
          }
          &:nth-child(4) {
            width: 130px;
          }
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .accounts-box {
      padding-bottom: 24px;
      .accounts-header {
        .my-header-container {
          height: auto;
          width: 100%;
          flex-direction: column;
          padding: 0 25px;
          .my-pages {
            margin: 0;
          }
        }
      }
      .accounts-list {
        overflow: auto;
        padding: 0 25px;
      }
    }
  }
</style>