<template>
  <div class="BPList-box">
    <div class="BPList-header">
      <div class="my-header-container">
        <div class="BPList-header-info">
          <h1>BPList</h1>
          <h4>showing total {{bpList.length}} BPs</h4>
        </div>
      </div>
    </div>

    <div class="bp-list">
      <ul class="my-list-header">
        <li>Name</li>
        <li>PubKey</li>
        <li>Location</li>
        <li>Url</li>
        <li>Online</li>
      </ul>
      <div class="list-wrap">
        <div class="list-body-wrap" v-for="bp in bpList">
          <ul class="my-list-body">
            <li>{{bp.name}}</li>
            <li>{{bp.pubkey}}</li>
            <li v-if="bp.loc == ''">-</li>
            <li v-else>{{bp.loc}}</li>
            <li v-if="bp.url == ''">-</li>
            <li v-else><a :href="bp.url" target="_blank">{{bp.name}}</a></li>
            <li v-if="bp.online == true" class="online"></li>
            <li v-else class="offline"></li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "BPList",
    data() {
      return {
        bpList: [],
      }
    },
    methods: {
      fetchData() {
        axios.get('/api/BPList').then((response) => {
          this.bpList = response.data.data
        })
      }
    },

    mounted: function () {
      this.fetchData()
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .BPList-box {
    padding-bottom: 100px;
    .BPList-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        height: 120px;
        width: 1000px;
        margin: 0 auto;
        display: flex;
        justify-content: space-between;
        text-align: left;
      }
    }

    .bp-list {
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
          text-align: center;
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

    .online:before {
      content: '\25CF';
      font-size: 20px;
      color: #0f0;
    }

    .offline:before {
      content: '\25CF';
      font-size: 20px;
      color: #f00;
    }
  }
</style>
