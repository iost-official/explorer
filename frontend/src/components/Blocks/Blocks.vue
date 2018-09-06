<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Blocks</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="#/">Home</a></li>
            <li class="active">Blocks</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" id="banner-luckybet">
        <i class="far fa-calendar-alt"></i> Latest Activity: <a href="/#/luckyBet">Play Lucky Bet !</a>
      </div>
    </div>

    <div class="container">
      <div class="row explorer-brief">
        <nav aria-label="..." class="pull-right">
          <ul class="pagination">
            <li><a href="/#/blocks" aria-label="First"><span aria-hidden="true">«</span></a></li>

            <li>
              <a v-if="page == 1" href="javascript:void(0)" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
              <a v-else :href="'/#/blocks?p=' + (page-1)" aria-label="Previous">
                <span aria-hidden="true"><</span>
              </a>
            </li>

            <li><a href="#">page <b>{{page}}</b> of <b>{{blockInfo.pageLast}}</b></a></li>
            <li>
              <a v-if="page == blockInfo.pageLast" href="javascript:void(0)" aria-label="Next">
                <span aria-hidden="true">></span>
              </a>
              <a v-else :href="'/#/blocks?p=' + (page+1)">
                <span aria-hidden="true">></span>
              </a>
            </li>
            <li><a :href="'/#/blocks?p=' + blockInfo.pageLast" aria-label="Last"><span aria-hidden="true">»</span></a></li>
          </ul>
        </nav>
      </div>
      <div class="row">
        <table class="table explorer-table">
          <thead>
          <tr>
            <th>Height</th>
            <th>Age</th>
            <th>Txn</th>
            <th>Witness</th>
            <th>GasLimit</th>
            <th>Avg.GasPrice</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="block in blockInfo.blockList">
            <td><a :href="'#/block/' + block.height">{{block.height}}</a></td>
            <td>{{block.age}}</td>
            <td>{{block.txn}}</td>
            <td><a :href="'#/account/' + block.witness">{{block.witness}}</a></td>
            <td>{{block.totalGasLimit}}</td>
            <td>{{block.avgGasPrice}}</td>
          </tr>

          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'


  export default {
    name: "Blocks",
    data() {
      return {
        // blockList: [],
        page: '',
        // totalPage: '',
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

        // axios.get('https://explorer.iost.io/api/blocks?p=' + this.page).then((response) => {
        // // axios.get('http://47.75.223.44:8080/api/blocks?p=' + this.page).then((response) => {
        //
        //   this.blockList = response.data.block_list
        //   this.totalPage = response.data.page_last
        // })
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
    }
  }
</script>
