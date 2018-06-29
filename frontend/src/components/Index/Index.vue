<template>
  <div>
    <div class="container">
      <div class="row" id="banner-luckybet">
        <i class="far fa-calendar-alt"></i> Latest Activity: <a href="/#/luckyBet">Play Lucky Bet !</a>
      </div>
    </div>
    <div class="container" style="margin-top:10px">
      <div class="row">
        <div class="col-md-6">
          <market-info></market-info>
          <IndexBlock></IndexBlock>
        </div>
        <div class="col-md-6">
          <IndexTxns></IndexTxns>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import MarketInfo from "./marketInfo";
  // import TxnCharts from "./chart";
  import IndexBlock from "./blocks";
  import IndexTxns from "./txns";

  export default {
    name: "Index",
    components: {IndexTxns, IndexBlock, MarketInfo}
  }
</script>
