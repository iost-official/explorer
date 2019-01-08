<template>
  <div class="form-search-box">
    <input type="text" class="form-control" placeholder="Search..." v-model.trim="searchInput" @keydown.enter="searchData">
    <img src="/static/img/search.png" alt="" @click="searchData">
  </div>
</template>

<script>
  import axios from 'axios';
  import { config } from '../../utils/config'
  const { apis } = config
  export default {
    data() {
      return {
        searchInput: ''
      }
    },

    methods: {
      searchData () {
        if (!this.searchInput) return
        axios.get(`${apis.search}${this.searchInput}`).then((response) => {
          var type = response.data.data.type
          if (type == "block") {
            if (response.data.data.text) {
              this.$router.push({
                path: '/block/' + response.data.text
              })
            } else {
              this.$router.push({
                path: '/block/' + this.searchInput
              })
            }
          } else if (type == "account") {
            this.$router.push({
              path: '/account/' + this.searchInput
            })
          } else if (type == "tx") {
            this.$router.push({
              path: '/tx/' + this.searchInput
            })
          } else {
            this.$router.push({
              path: '/search/' + this.searchInput
            })
          }
          this.searchInput = ''
          this.isShow = false
        })
        return false
      }

    }

  }
</script>

<style lang="less" rel="stylesheet/less" scoped>
  .form-search-box {
    margin: 165px auto 0;
    width: 780px;
    position: relative;
    .form-control {
      display: inline-block;
      width: 100%;
      height: 50px;
      padding: 6px 12px;
      font-size: 14px;
      line-height: 1.5;
      color: #555;
      background-color: #fff;
      border: 1px solid #ccc;
      border-radius: 4px;
      box-shadow: inset 0 1px 1px rgba(0,0,0,.075);
    }
    > img {
      width: 22px;
      position: absolute;
      top: 14px;
      right: 14px;
    }
  }

  @media screen and (max-width:480px) {
    .form-search-box {
      margin: 140px auto 0;
      width: 300px;
    }
  }
</style>