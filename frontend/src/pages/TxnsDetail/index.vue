<template>
  <div class="txnsDetail-box">

    <div class="txnsDetail-header">
      <div class="my-header-container">
        <h1>Transaction Information</h1>
        <p>#{{txHash}}</p>
      </div>
    </div>

    <div class="txnsDetail-information">
      <img src="/static/img/iostWhite.png" alt="">
      <div class="txnsDetail-hash">
        <h4>TxHash:</h4>
        <p>{{txHash}}</p>
      </div>
      <div class="txnsDetail-block-time">
        <div class="txnsDetail-block">
          <h4>Block Height:</h4>
          <p><router-link :to="{path:`/block/${txnDetail.blockHeight}`}">{{txnDetail.blockHeight}}</router-link></p>
        </div>
        <div class="txnsDetail-time">
          <h4>TimeStamp:</h4>
          <p>{{txnDetail.age}}({{txnDetail.utcTime}})</p>
        </div>
      </div>
      <div class="txnsDetail-from">
        <h4>From:</h4>
        <p><router-link :to="{path:`/account/${txnDetail.from}`}">{{txnDetail.from}}</router-link></p>
      </div>


      <div class="txnsDetail-to" v-if="txnDetail.actionName != 'transfer'">
        <h4>To:</h4>
        <p v-if="txnDetail.to == 'iost.system'">{{txnDetail.to}}</p>
        <p><span>actionName：</span>{{txnDetail.actionName}}</p>

        <div class="my-pre" v-if="txnDetail.actionName == 'SetCode'">
          <span>code：</span>
          <pre>{{txnDetail.code}}</pre>
        </div>
        <div class="my-pre-normal" v-else>
          <span>data：</span><pre>{{txnDetail.data}}</pre>
        </div>

      </div>
      <div class="txnsDetail-to" v-else>
        <h4>To:</h4>
        <p><router-link :to="{path:`/account/${txnDetail.to}`}">{{txnDetail.to}}</router-link></p>
      </div>


      <div class="txnsDetail-value-gas">
        <div class="txnsDetail-value">
          <h4>Value:</h4>
          <p>{{txnDetail.amount}} IOST</p>
        </div>
        <div class="txnsDetail-gas-limit">
          <h4>Gas Limit:</h4>
          <p>{{txnDetail.gasLimit}}</p>
        </div>
        <div class="txnsDetail-gas-price">
          <h4>Gas Price:</h4>
          <p>{{txnDetail.price}}</p>
        </div>
      </div>
      <div v-if="txnDetail.code">
        <h4>Code</h4>
        <pre>{{txnDetail.code}}</pre>
      </div>

    </div>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: "Tx",
    data() {
      return {
        txHash: '',
      }
    },
    methods: {
      fetchData(r) {
        this.txHash = r.params.id

        this.$store.dispatch('getTxnDetail', this.txHash)
      }
    },

    computed: {
      ...mapState(['txnDetail'])
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
  .txnsDetail-box {
    padding-bottom: 160px;
    background: #F6F7F8;
    .txnsDetail-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        width: 1000px;
        margin: 0 auto;
        height: 120px;
        text-align: left;
        overflow: hidden;
        > h1 {
          font-size: 36px;
          line-height: 44px;
          margin: 21px 0 15px;
        }
        > p {
          font-size: 14px;
          height: 18px;
          margin: 0;
        }
      }
    }
    .txnsDetail-information {
      width: 1000px;
      margin: 24px auto 0;
      text-align: left;
      background: #FFFFFF;
      padding: 15px 50px;
      position: relative;
      box-shadow: 0 2px 3px rgba(0,0,0,0.1);
      h4 {
        font-size: 14px;
        line-height: 18px;
        font-weight: bold;
        padding-bottom: 12px;
        margin: 0;
        color: #2c2e31;
        &:last-child {
          padding-bottom: 60px;
        }
      }
      pre {
        margin-bottom: 0;
      }
      p {
        /*font-size: 18px;*/
        font-size: 14px;
        line-height: 22px;
        margin-top: 20px;
        margin-bottom: 0;
        font-weight: 300;
        a {
          color: #4b78aa;
        }
      }

      > div {
        margin-bottom: 60px;
        &:last-child {
          margin-bottom: 0;
        }
      }
      > img {
        height: 50px;
      }
      .txnsDetail-hash {
        margin-top: 35px;
      }
      .txnsDetail-to {
        .my-pre {
          margin-top: 20px;
          > span {
            display: inline-block;
            margin-bottom: 12px;
          }
        }
        .my-pre-normal {
          display: flex;
          align-items: center;
          margin-top: 20px;
          > pre {
            display: inline-block;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            max-width: 850px;
          }
        }
        span {
          font-weight: bold;
        }

      }

      .txnsDetail-hash, .txnsDetail-from, .txnsDetail-to {
        >h4 {
          border-bottom: 1px solid #f6f7f8;
        }
      }
      .txnsDetail-block-time {
        display: flex;
        .txnsDetail-block {
          width: 33.3%;
        }
        > div {
          >h4 {
            border-bottom: 1px solid #f6f7f8;
          }
        }
      }
      .txnsDetail-value-gas {
        display: flex;
        > div {
          width: 33.3%;
          >h4 {
            border-bottom: 1px solid #f6f7f8;
          }
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .txnsDetail-box {
      padding-bottom: 24px;
      .txnsDetail-header {
        .my-header-container {
          height: auto;
          width: 100%;
          flex-direction: column;
          padding: 0 25px;
          > h1 {
            font-size: 30px;
          }
          .my-pages {
            margin: 0;
          }
        }
      }
      .txnsDetail-information {
        width: 100%;
        padding: 15px 25px 0 25px;
        p {
          text-overflow: ellipsis;
          overflow: hidden;
        }
        .txnsDetail-block-time {
          .txnsDetail-block {
            width: 50%;
          }
        }
      }
    }
  }
</style>