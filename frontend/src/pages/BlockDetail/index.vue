<template>
  <div class="blockDetail-box">
    <div class="luckyBet-box">
      <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet">Play Lucky Bet !</a>
    </div>

    <div class="blockDetail-header">
      <div class="my-header-container">
        <h1>Block Information</h1>
        <p>#{{blockHeight}}</p>
      </div>
    </div>



    <div class="blockDetail-information">
      <img src="../../assets/iostWhite.png" alt="">

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
        <p><a :href="'/#/account/' + blockDetail.witness">{{blockDetail.witness}}</a></p>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'

  export default {
    name: "Block",
    data() {
      return {
        blockHeight: '',
        // age: '',
        // utcAge: '',
        // txnLen: '',
        // txnHash: '',
        // parentHash: '',
        // blockHash: '',
        // witness: '',
      }
    },
    methods: {
      fetchData(r) {
        this.blockHeight = r.params.id
        this.$store.dispatch('getBlockDetail', this.blockHeight)

        // axios.get('https://explorer.iost.io/api/block/' + this.blockHeight).then((response) => {
        //   this.blockHeight = response.data.height
        //   this.age = response.data.age
        //   this.utcAge = response.data.utc_time
        //   this.txnLen = response.data.txn
        //   this.parentHash = response.data.parent_hash
        //   this.blockHash = response.data.block_hash
        //   this.witness = response.data.witness
        // })
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
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .blockDetail-box {
    padding-top: 90px;
    margin: 0 auto;
    background: #F6F7F8;
    padding-bottom: 160px;
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
        font-weight: bold;
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
          color: #4C6889;
        }
      }

      p {
        font-size: 18px;
        line-height: 22px;
        color: #2C2E31;
        margin-top: 20px;
        margin-bottom: 0;
        a {
          font-size: 18px;
          line-height: 22px;
          color: #2C2E31;
        }
      }
    }
  }
</style>
