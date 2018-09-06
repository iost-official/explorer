<template>
  <div>
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/"><img src="../assets/logo.svg"></a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
            <li><a href="/#/">Home</a></li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Blockchain <span class="caret"></span></a>
              <ul class="dropdown-menu">
                <li><a href="/#/txs">View Txns</a></li>
                <li><a href="/#/blocks">View Blocks</a></li>
              </ul>
            </li>
            <li><a href="/#/accounts">Accounts</a></li>
            <li><a href="/#/applyIOST" ref="link">Request Test IOST</a></li>
          </ul>
          <form class="navbar-form navbar-right">
            <input type="text" class="form-control" placeholder="Search..." v-model.trim="searchInput" @keydown.enter="searchData()">
            <input type="text" style="display:none;">
          </form>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "Head",
    data() {
      return {
        searchInput: ''
      }
    },
    methods: {
      searchData: function () {
        if (!this.searchInput) return
        // axios.get('https://explorer.iost.io/api/search/' + this.searchInput).then((response) => {
        axios.get('http://47.75.223.44:8080/api/search/' + this.searchInput).then((response) => {
          console.log(response.data)
          var type = response.data.type
          console.log(type)
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
              path: '/search/' + this.searchInput
            })
          }
        })
        return false
      }
    }
  }
</script>
