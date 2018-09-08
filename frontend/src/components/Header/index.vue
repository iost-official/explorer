<template>
  <div class="header-box" :class="{active: currentColor2}">
    <div class="my-container">
      <div class="logo-box">
        <router-link class="logo-link" to="/">
          <img src="../../assets/logo.png" alt="">
        </router-link>
      </div>
      <ul class="my-nav-box">
        <li class="my-nav-item" :class="{active: currentColor == 0}">
          <router-link to="/">HOME</router-link>
        </li>
        <li class="my-nav-item" :class="{active: currentColor == 1}">
          <router-link to="/blocks">BLOCK</router-link>
        </li>
        <li class="my-nav-item" :class="{active: currentColor == 2}">
          <router-link to="/txs">TRANSACTION</router-link>
        </li>

        <li class="my-nav-item" :class="{active: currentColor == 3}">
          <router-link to="/accounts">ACCOUNTS</router-link>
        </li>
        <li class="my-nav-item" :class="{active: currentColor == 4}">
          <router-link to="/applyIOST">REQUEST TEST IOST</router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Head",
    data() {
      return {
        currentColor: 0,
        currentColor2: false
      }
    },
    watch: {
      '$route' (to) {
        const arr = ['/block','/tx','/account','/applyIOST']
        let arr2 = arr.filter((a, index) => {
          // console.log(to.path.search(a))
          return to.path.search(a)+1
        })
        let result = arr2.join('')
        if (result == '/block') {
          this.currentColor = 1
        } else if (result == '/tx') {
          this.currentColor = 2
        } else if (result == '/account') {
          this.currentColor = 3
        } else if (result == '/applyIOST') {
          this.currentColor = 4
        } else {
          this.currentColor = 0
        }
      }
    },
    
    created () {
      if (window.location.pathname.indexOf('block')) {
        this.currentColor = 1
      }else if (window.location.pathname.indexOf('tx')) {
        this.currentColor = 2
      }else if (window.location.pathname.indexOf('account')) {
        this.currentColor = 3
      }else if (window.location.pathname.indexOf('applyIOST?r')) {
        this.currentColor = 4
      } else if (window.location.pathname == '/') {
        this.currentColor2 = true
        this.currentColor = 0
      }

      if (window.location.pathname.indexOf('applyIOST?r')) {
        this.currentColor = 4
      }
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
      background-color: rgba(44,46,49, 0.75);
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
        .logo-link{
          > img {
            width: 130px;
          }
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
          a {
            color: #8F9A9C;
            text-decoration: none;
            font-size: 14px;
            line-height: 18px;
          }
          &:nth-child(2) {
            position: relative;
            &:hover {
              .my-dropdown {
                display: block;
              }
            }
          }
        }
      }


    }
  }
</style>
