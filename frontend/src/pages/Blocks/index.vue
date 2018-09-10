<template>
  <div class="blocks-box">
    <div class="luckyBet-box">
      <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
    </div>

    <div class="blocks-header">
      <div class="my-header-container">
        <h1>Blocks</h1>
        <ul class="pagination my-pages">
          <li><router-link to="/blocks" aria-label="First"><span aria-hidden="true"><img src="../../assets/arrow-douleft.png" alt=""></span></router-link></li>

          <li>
            <a v-if="page == 1" href="javascript:void(0)" aria-label="Previous">
              <span aria-hidden="true"><img src="../../assets/arrow-left.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/blocks?p=${(page-1)}`}" aria-label="Previous">
              <span aria-hidden="true"><img src="../../assets/arrow-left.png" alt=""></span>
            </router-link>
          </li>

          <li><a href="#">page <b>{{page}}</b> of <b>{{blockInfo.pageLast}}</b></a></li>
          <li>
            <a v-if="page == blockInfo.pageLast" href="javascript:void(0)" aria-label="Next">
              <span aria-hidden="true"><img src="../../assets/arrow-right.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/blocks?p=${(page+1)}`}">
              <span aria-hidden="true"><img src="../../assets/arrow-right.png" alt=""></span>
            </router-link>
          </li>
          <li><router-link :to="{path:`/blocks?p=${blockInfo.pageLast}`}" aria-label="Last"><span aria-hidden="true"><img src="../../assets/arrow-douright.png" alt=""></span></router-link></li>
        </ul>
      </div>
    </div>

    <div class="blocks-list">
      <ul class="my-list-header">
        <li>Height</li>
        <li>Witness</li>
        <li>GasLimit</li>
        <li>Avg.GasPrice</li>
        <li>Age</li>
        <li>Txn</li>
      </ul>
      <div class="list-wrap">
        <div class="list-body-wrap" v-for="block in blockInfo.blockList">
          <ul class="my-list-body">
            <li><router-link :to="{path:`/block/${block.height}`}">{{block.height}}</router-link></li>
            <li><router-link :to="{path:`/account/${block.witness}`}">{{block.witness}}</router-link></li>
            <li>{{block.totalGasLimit}}</li>
            <li>{{block.avgGasPrice}}</li>
            <li>{{block.age}}</li>
            <li>{{block.txn}}</li>
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
    name: "Blocks",
    data() {
      return {
        // blockList: [],
        page: '',
        // totalPage: '',
      }
    },
    methods: {
      fetchData(r) {
        if (!r.query.p) {
          this.page = 1
        } else {
          this.page = parseInt(r.query.p)
        }

        this.$store.dispatch('getBlockInfo',this.page)

        // axios.get('https://explorer.iost.io/api/blocks?p=' + this.page).then((response) => {
        // // axios.get('http://47.75.223.44:8080/api/blocks?p=' + this.page).then((response) => {
        //
        //   this.blockList = response.data.block_list
        //   this.totalPage = response.data.page_last
        // })
      }
    },

    computed: {
      ...mapState(['blockInfo'])
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
  .blocks-box {
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
    .blocks-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        display: flex;
        justify-content: space-between;
        height: 85px;
        width: 1000px;
        margin: 0 auto;
        > h1 {
          margin-top: 21px;
          font-size: 36px;
          line-height: 44px;
          font-weight: bold;
        }
        .my-pages {
          margin-top: 31px;
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
      }
    }
    .blocks-list {
      min-height: 366px;
      margin-top: 2px;
      .my-list-header {
        font-size: 12px;
        line-height: 15px;
        font-weight: 500;
        > li {
          color: #2C2E31;
        }
      }
      &::-webkit-scrollbar {
        width: 0;
      }
      .list-wrap {
        max-height: 1000px;
        overflow: auto;
        .list-body-wrap {
          width: 100%;
          &:hover {
            box-shadow: 0 8px 30px 0 rgba(0, 0, 0, .15);
          }
        }
      }

      ul {
        display: flex;
        padding: 24px 0 21px;
        margin: 0 auto;
        width: 1000px;
        border-bottom: 1px solid #f6f7f8;

        li {
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
            width: 500px;
          }
          &:nth-child(3) {
            width: 100px;
          }
          &:nth-child(4) {
            width: 110px;
          }
          &:nth-child(5) {
            width: 120px;
          }
          &:nth-child(6) {
            width: 50px;
          }
        }
      }
    }

  }
</style>
