<template>
  <div class="blocks">
    <div class="list-banner">
      <div class="list-title pull-left">
        <i class="fa fa-cubes"></i> Blocks
      </div>
      <a href="/#/blocks" class="pull-right">View All</a>
    </div>
    <ul class="list-body">
      <li v-for="block in blockList">
        <div class="list-body-left pull-left">
          <a :href="'#/block/' + block.height">Block {{block.height}}</a>
          <span>{{block.age}}</span>
        </div>
        <div class="list-body-right pull-left">
          Minted By
          <p><a :href="'/#/account/' + block.witness" style="display: inline">{{block.witness}}</a></p>
          <p><a :href="'/#/txs/?b=' + block.height" style="text-align: left;">{{block.txn}} txns</a></p>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "IndexBlock",
    data() {
      return {
        blockList: []
      }
    },
    mounted: function () {
      axios.get('https://explorer.iost.io/api/indexBlocks').then((response) => {
        this.blockList = response.data
      })
    }
  }
</script>
