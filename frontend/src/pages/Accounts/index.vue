<template>
  <div class="accounts-box">
    <div class="luckyBet-box">
      <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
    </div>

    <div class="accounts-header">
      <div class="my-header-container">
        <div class="accounts-header-info">
          <h1>All Accounts</h1>
          <div class="accounts-info">
            <h4 v-if="accountInfo.pageLast * 25 < 500 ">{{accountInfo.totalLen}} accounts found</h4>
            <h4 v-else>{{accountInfo.totalLen}} accounts found (showing the last 500 records)</h4>
          </div>
        </div>
        <ul class="pagination my-pages">
          <li><a href="/accounts" aria-label="First"><span aria-hidden="true"><img src="../../assets/arrow-douleft.png" alt=""></span></a></li>

          <li>
            <a v-if="page == 1" href="javascript:void(0)" aria-label="Previous">
              <span aria-hidden="true"><img src="../../assets/arrow-left.png" alt=""></span>
            </a>
            <a v-else :href="'/accounts?p=' + (page-1)" aria-label="Previous">
              <span aria-hidden="true"><img src="../../assets/arrow-left.png" alt=""></span>
            </a>
          </li>

          <li><a href="#">page <b>{{page}}</b> of <b>{{accountInfo.pageLast}}</b></a></li>
          <li>
            <a v-if="page == accountInfo.pageLast" href="javascript:void(0)" aria-label="Next">
              <span aria-hidden="true"><img src="../../assets/arrow-right.png" alt=""></span>
            </a>
            <a v-else :href="'/accounts?p=' + (page+1)">
              <span aria-hidden="true"><img src="../../assets/arrow-right.png" alt=""></span>
            </a>
          </li>

          <li><a :href="'/accounts?p=' + accountInfo.pageLast" aria-label="Last"><span aria-hidden="true"><img src="../../assets/arrow-douright.png" alt=""></span></a></li>
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
            <li>{{account.balance.toFixed(2)}}</li>
            <li>{{account.txCount}}</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'


  export default {
    name: 'accounts',
    data() {
      return {
        // accountList: [],
        page: '',
        // totalPage: '',
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

        this.$store.dispatch('getAccountInfo',this.page)


        // axios.get('https://explorer.iost.io/api/accounts?p=' + this.page).then((response) => {
        //   this.accountList = response.data.account_list
        //   this.totalPage = response.data.page_last
        //   this.totalLen = response.data.total_len
        // })
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
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .accounts-box {
    padding-top: 90px;
    margin: 0 auto;
    .luckyBet-box {
      background: #2C2E31;
      height: 50px;
      line-height: 50px;
      color: #F6F7F8;
      font-size: 14px;
      > img {
        width: 24px;
        height: 24px;
        margin-right: 12px;
      }
      a {
        color: #F6F7F8;
        font-size: 14px;
        line-height: 18px;
        text-decoration: none;
      }
    }
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
          margin-top: 60px;
          li {
            a {
              padding: 5px;
              height: 32px;
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
          > h1 {
            font-size: 36px;
            line-height: 44px;
            font-weight: bold;

          }
          .accounts-info {
            > h4 {
              text-align: left;
              color: #2C2E31;
              font-size: 14px;
              line-height: 18px;
            }
          }
        }
      }
    }


    .accounts-list {
      min-height: 500px;
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
          box-shadow: 0 8px 30px 0 rgba(0, 0, 0, .15);
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
          color: #4C6889;
          a {
            font-size: 14px;
            line-height: 18px;
            color: #4C6889;
            text-decoration: none;
          }
          &:first-child {
            width: 100px;
          }
          &:nth-child(2) {
            width: 550px;

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
</style>