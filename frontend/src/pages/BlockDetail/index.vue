<template>
  <div class="blockDetail-box">
    <LuckyBet/>

    <div class="blockDetail-header">
      <div class="my-header-container">
        <h1>Block Information</h1>
        <p>#{{blockHeight}}</p>
      </div>
    </div>

    <div class="blockDetail-information">
      <img src="/static/img/iostWhite.png" alt="">

      <div class="blockDetail-height-txns">
        <div class="blockDetail-height">
          <h4>Height:</h4>
          <p>{{blockHeight}}</p>
        </div>
        <div class="blockDetail-txns">
          <h4>Transactions:</h4>
          <p><router-link :to="{path:`/txs/?b=${blockHeight}`}">{{blockDetail.txn}}</router-link></p>
        </div>
      </div>
      <div class="blockDetail-time">
        <h4>TimeStamp:</h4>
        <p>{{blockDetail.age}} ({{blockDetail.utcTime}})</p>
      </div>
      <div class="blockDetail-hash">
        <h4>Hash:</h4>
        <p>{{blockDetail.blockHash}}</p>
      </div>
      <div class="blockDetail-parent">
        <h4>Parent Hash:</h4>
        <p>{{blockDetail.parentHash}}</p>
      </div>
      <div class="blockDetail-witness">
        <h4>Witness:</h4>
        <p><router-link :to="{path:`/account/${blockDetail.witness}`}">{{blockDetail.witness}}</router-link></p>
      </div>
    </div>
  </div>
</template>

<script>
  import LuckyBet from '../../components/LuckyBet'
  import { mapState } from 'vuex'

  export default {
    name: "Block",
    data() {
      return {
        blockHeight: '',
      }
    },
    methods: {
      fetchData(r) {
        this.blockHeight = r.params.id
        this.$store.dispatch('getBlockDetail', this.blockHeight)
      }
    },

    computed: {
      ...mapState(['blockDetail'])
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
  .blockDetail-box {
    padding-bottom: 160px;
    background: #F6F7F8;
    .blockDetail-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        width: 1000px;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        height: 120px;
        > h1 {
          font-size: 36px;
          line-height: 44px;
          margin: 21px 0 15px;
          font-weight: bold;
        }
        > p {
          font-size: 14px;
          height: 18px;
          margin: 0;
        }
      }
    }
    .blockDetail-information {
      width: 1000px;
      margin: 24px auto 0;
      background: #FFFFFF;
      text-align: left;
      padding: 15px 50px 0 50px;
      box-shadow: 0 2px 3px rgba(0,0,0,0.1);
      > img {
        height: 50px;
      }
      h4 {
        font-size: 14px;
        line-height: 18px;
        font-weight: bolder;
        padding-bottom: 12px;
        margin: 0;
        color: #2c2e31;
        border-bottom: 1px solid #F1F1F1;
      }
      > div {
        margin-bottom: 60px;
        &:last-child {
          margin-bottom: 0;
        }
      }
      .blockDetail-height-txns {
        display: flex;
        margin-top: 35px;
        .blockDetail-height, .blockDetail-txns {
          width: 50%;
        }
      }
      .blockDetail-witness {
        padding-bottom: 60px;
        a {
          color: #4b78aa;
        }
      }

      p {
        /*font-size: 18px;*/
        font-size: 14px;
        line-height: 22px;
        color: #2C2E31;
        margin-top: 20px;
        margin-bottom: 0;
        a {
          /*font-size: 18px;*/
          font-size: 14px;
          line-height: 22px;
          color: #2C2E31;
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .blockDetail-box {
      padding-bottom: 24px;
      .blockDetail-header {
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
      .blockDetail-information {
        width: 100%;
        padding: 15px 25px 0 25px;
        p {
          text-overflow: ellipsis;
          overflow: hidden;
        }
      }
    }
  }
</style>
