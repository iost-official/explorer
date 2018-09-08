<template>
  <div class="homeBlocks-box">
    <div class="my-list-banner">
      <div class="my-list-title">
        BLOCKS
      </div>
      <router-link to="/blocks">View All</router-link>
    </div>
    <ul class="my-list-body">
      <li v-for="block in indexBlockList">
        <img src="../../assets/block.png" alt="">
        <div class="my-list-wrap">
          <router-link :to="{path:`/block/${block.height}`}">Block {{block.height}}</router-link>
          <p>Minted By</p>
          <router-link :to="{path:`/block/${block.height}`}">{{block.witness}}</router-link>
          <router-link :to="{path:`/txs/?b=${block.height}`}" style="text-align: left;">{{block.txn}} txns</router-link>
          <p>{{block.age}}</p>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: "IndexBlock",

    computed: {
      ...mapState(['indexBlockList'])
    },

    mounted () {
      this.$store.dispatch('getIndexBlockList')
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .homeBlocks-box {
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
      padding-top: 8px;
      padding-left: 0;
      li {
        list-style: none;
        text-align: left;
        padding: 16px 0 16px 30px;
        display: flex;
        /*margin-bottom: 4px;*/
        height: 160px;
        /*height: 200px;*/

        > img {
          width: 24px;
          height: 24px;
        }
        &:nth-child(2n) {
          background: #F6F7F8;
        }
        &:nth-child(2n+1) {
          background: #FFFFFF;
        }
      }
      .my-list-wrap {
        margin-left: 26px;
        margin-top: 2px;
        line-height: 1;
        font-size: 0;
        > a {
          display: inline-block;
          text-decoration: none;
          color: #4C6889;
          &:first-of-type {
            font-size: 14px;
            line-height: 18px;
            font-weight: bold;
          }
          &:nth-of-type(2), &:nth-of-type(3) {
            font-size: 12px;
            line-height: 16px;
            margin-top: 8px;
          }
          &:nth-of-type(2) {
            width: 380px;
            text-overflow: ellipsis;
            overflow: hidden;
          }
        }
        > p{
          font-size: 12px;
          line-height: 15px;
          color: #2C2E31;
          margin-bottom: 0;
          &:nth-of-type(1) {
            margin-top: 20px;
          }
          &:nth-of-type(2) {
            margin-top: 8px;
          }
        }
      }
    }
  }

</style>
