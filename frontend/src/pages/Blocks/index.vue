<template>
  <div class="blocks-box">
    <LuckyBet/>

    <div class="blocks-header">
      <div class="my-header-container">
        <div class="blocks-info">
          <h1>Blocks</h1>
          <h4>showing the last 30 records</h4>
        </div>
        <ul class="my-pages">
          <li>
            <router-link to="/blocks">
              <span><img src="/static/img/arrow-douleft.png" alt=""></span>
            </router-link>
          </li>

          <li>
            <a v-if="page == 1" href="javascript:void(0)">
              <span><img src="/static/img/arrow-left.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/blocks?p=${(page-1)}`}">
              <span><img src="/static/img/arrow-left.png" alt=""></span>
            </router-link>
          </li>

          <li><a class="page-auto" href="#"><b>{{page}}</b> / <b>{{blockInfo.pageLast}}</b></a></li>
          <li>
            <a v-if="page == blockInfo.pageLast" href="javascript:void(0)">
              <span><img src="/static/img/arrow-right.png" alt=""></span>
            </a>
            <router-link v-else :to="{path:`/blocks?p=${(page+1)}`}">
              <span><img src="/static/img/arrow-right.png" alt=""></span>
            </router-link>
          </li>
          <li>
            <router-link :to="{path:`/blocks?p=${blockInfo.pageLast}`}">
              <span aria-hidden="true"><img src="/static/img/arrow-douright.png" alt=""></span>
            </router-link>
          </li>
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
  import LuckyBet from '../../components/LuckyBet'
  import { mapState } from 'vuex'

  export default {
    name: "Blocks",
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

        this.$store.dispatch('getBlockInfo',this.page)
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
    },

    components: {
      LuckyBet
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .blocks-box {
    padding-bottom: 100px;
    .blocks-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        height: 120px;
        width: 1000px;
        margin: 0 auto;
        display: flex;
        justify-content: space-between;
        text-align: left;
        .blocks-info {
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
              font-size: 12px;
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
    .blocks-list {
      min-height: 366px;
      margin-top: 2px;
      .my-list-header {
        /*font-size: 12px;*/
        font-size: 14px;
        line-height: 15px;
        font-weight: 500;
        > li {
          color: #2C2E31;
        }
      }
      .list-wrap {
        .list-body-wrap {
          width: 100%;
          &:hover {
            box-shadow: 0 2px 10px 0 rgba(0, 0, 0, .08);
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
          color: #2c2e31;
          a {
            font-size: 14px;
            line-height: 18px;
            color: #4b78aa;
          }
          &:first-child {
            width: 100px;
          }
          &:nth-child(2) {
            width: 450px;
            padding-right: 50px;
            overflow: hidden;
            text-overflow: ellipsis;
          }
          &:nth-child(3) {
            width: 100px;
          }
          &:nth-child(4) {
            width: 180px;
          }
          &:nth-child(5) {
            width: 120px;
          }
          &:nth-child(6) {
            width: 70px;
          }
        }
      }
    }

  }

  @media screen and (max-width:480px) {
    .blocks-box {
      .blocks-header {
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
      .blocks-list {
        overflow: auto;
        padding: 0 25px;
      }
    }
  }
</style>
