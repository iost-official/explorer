<template>
  <div class="header-box" :class="{active: currentTheme}">
    <div class="my-container">
      <div class="logo-box">
        <router-link class="logo-link" to="/">
          <img src="../../assets/logo.png" alt="">
        </router-link>
      </div>
      <ul class="my-nav-box">
        <li class="my-nav-item" v-show="!isShow" :class="{active: currentColor == 0}">
          <router-link to="/">HOME</router-link>
        </li>
        <li class="my-nav-item" v-show="!isShow" :class="{active: currentColor == 1}">
          <router-link to="/blocks">BLOCK</router-link>
        </li>
        <li class="my-nav-item" v-show="!isShow" :class="{active: currentColor == 2}">
          <router-link to="/txs">TRANSACTION</router-link>
        </li>

        <li class="my-nav-item" v-show="!isShow" :class="{active: currentColor == 3}">
          <router-link to="/accounts">ACCOUNTS</router-link>
        </li>
        <li class="my-nav-item" v-show="!isShow" :class="{active: currentColor == 4}">
          <router-link to="/applyIOST">REQUEST TEST IOST</router-link>
        </li>

        <li class="my-nav-item searchActive" v-show="isShow">
          <input type="text" placeholder="Search..."
                 @click.stop="" v-model.trim="searchInput" @keydown.enter="searchData">
        </li>
        <li class="my-nav-item" v-show="!currentTheme">
          <img src="../../assets/search.png" alt="" @click="openSearch"/>
        </li>
      </ul>

    </div>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "Head",
    data() {
      return {
        currentColor: 0,
        currentTheme: true,
        isShow: false,
        searchInput: ''

      }
    },
    watch: {
      '$route' (to) {
        const arr = ['/block','/tx','/account','/applyIOST','/feedback','/Search','/luckyBet']
        let arr2 = arr.filter((a, index) => {
          return to.path.search(a)+1
        })
        let result = arr2.join('')
        if (result == '/block') {
          this.currentColor = 1
          this.currentTheme = false
        } else if (result == '/tx') {
          this.currentColor = 2
          this.currentTheme = false
        } else if (result == '/account') {
          this.currentColor = 3
          this.currentTheme = false
        } else if (result == '/applyIOST') {
          this.currentColor = 4
          this.currentTheme = false
        }else if (result == '/feedback') {
          this.currentColor = 5
          this.currentTheme = false
        }else if (result == '/Search') {
          this.currentColor = 6
          this.currentTheme = false
        } else if (result == '/luckyBet') {
          this.currentColor = 7
          this.currentTheme = false
        }else {
          this.currentColor = 0
          this.currentTheme = true
        }
      }
    },

    methods: {
      onScroll () {
        if (location.pathname != '/') {
          return
        }
        const heightTop = document.documentElement.scrollTop || document.body.scrollTop;
        if (heightTop >= 60) {
          this.currentTheme = false

        } else {
          this.currentTheme = true

        }
      },

      openSearch (e) {
        if(e.stopPropagation){
          e.stopPropagation();
        }
        this.isShow = !this.isShow
        if (this.isShow) {
          console.log(132131)
          setTimeout(() => {
            document.querySelector('input').focus()
          })
        }
      },

      onSearch () {
        this.isShow = false
      },

      searchData () {
        if (!this.searchInput) return
        axios.get('http://47.75.223.44:8080/api/search/' + this.searchInput).then((response) => {
          var type = response.data.data.type
          if (type == "block") {
            if (response.data.text) {
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
              path: '/Search/' + this.searchInput
            })
          }
          this.searchInput = ''
          this.isShow = false
        })
        return false
      }

    },
    
    created () {
      const arr = ['/block','/tx','/account','/applyIOST','/feedback','/Search','/luckyBet']
      let arr2 = arr.filter((a, index) => {
        return location.pathname.search(a)+1
      })
      let result = arr2.join('')
      if (result == '/block') {
        this.currentColor = 1
        this.currentTheme = false
      } else if (result == '/tx') {
        this.currentColor = 2
        this.currentTheme = false
      } else if (result == '/account') {
        this.currentColor = 3
        this.currentTheme = false
      } else if (result == '/applyIOST') {
        this.currentColor = 4
        this.currentTheme = false
      }else if (result == '/feedback') {
        this.currentColor = 5
        this.currentTheme = false
      }else if (result == '/Search') {
        this.currentColor = 6
        this.currentTheme = false
      }else if (result == '/luckyBet') {
        this.currentColor = 7
        this.currentTheme = false
      }  else {
        this.currentColor = 0
        this.currentTheme = true
      }

    },

    mounted () {
      window.addEventListener('scroll', this.onScroll);
      window.addEventListener('click', this.onSearch);
    }
  }
</script>

<style lang="less" rel="stylesheet/less" scoped>
  .header-box {
    position: fixed;
    width: 100%;
    height: 90px;
    display: flex;
    z-index: 100;
    background: #2c2e31;
    box-shadow: 0 2px 3px rgba(0,0,0, 0.1);
    border: 1px solid #2c2e31;
    &.active {
      background-color: rgba(44,46,49, 0.5);
    }
    .my-container{
      width: 1000px;
      height: 100%;
      margin: 0 auto;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .logo-box {
        position: relative;
        display: flex;
        align-items: center;
        margin-left: -19px;
        .logo-link{
          > img {
            width: 130px;
          }
        }
      }
      .my-search {
        padding-left: 30px;
        margin-right: -20px;
        .my-input {

        }
      }
      .my-nav-box {
        display: flex;
        margin-bottom: 0;
        height: 100%;

        .my-nav-item {
          display: flex;
          align-items: center;
          list-style: none;
          margin-left: 35px;
          height: 100%;
          &.active {
            a {
              color: #FFFFFF;
            }
          }
          &.searchActive {
            margin-right: -25px;
          }
          a {
            color: #8F9A9C;
            text-decoration: none;
            font-size: 14px;
            line-height: 18px;
          }
          > img {
            width: 18px;
            cursor: pointer;
          }

          > input {
            width: 450px;
            border: none;
            border-bottom: 1px solid rgba(143,145,156,0.5);
            background-color: #2c2e31;
            color: #FFFFFF;
            padding-left: 5px;
            &:focus {
              outline: none;
            }
          }
        }
      }
    }
  }
</style>
