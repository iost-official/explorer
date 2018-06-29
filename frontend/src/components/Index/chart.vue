<template>
  <div class="charts" :id="id" :option="option"></div>
</template>

<script>
  import axios from 'axios';
  import HighCharts from 'highcharts'

  export default {
    name: "txnCharts",
    data() {
      return {
        id: "txnCharts",
        option: {
          title: {
            text: '14 day IOST Transactions History'
          },

          xAxis: {
            labels: {
              format: "{value:%m-%d}",
              rotation: -30,
            },
            tickmarkPlacement: 'on',
          },

          yAxis: {
            title: {
              text: null
            }
          },

          series: [{
            name: 'Transactions',
            data: []
          }],

          tooltip: {
            formatter: function () {
              var newDate = new Date()
              newDate.setTime(this.x)
              return newDate.toDateString() + '<br>' +
                '<span style="color:#7cb5ec">Transactions: </span><b>' + this.y + '</b>'
            }
          },

          credits: {
            enabled: false
          }
        }
      }
    },
    mounted() {
      axios.get('https://explorer.iost.io/api/indexBlocks').then((response) => {
        var jsonData = {"2018-05-24":20569,"2018-06-04":28753,"2018-05-25":23059,"2018-06-03":22407,"2018-06-02":34091,"2018-05-23":16423,"2018-06-01":22233,"2018-05-28":24396,"2018-05-29":55395,"2018-05-26":35262,"2018-06-06":45847,"2018-05-27":17640,"2018-06-05":20908,"2018-05-31":19050,"2018-05-30":26296}
        var xData = []
        var yData = []
        for (var i in jsonData) {
          xData.push(Date.parse(i))
          yData.push(jsonData[i])
        }
        this.option.xAxis.categories = xData
        this.option.series[0].data = yData
        HighCharts.chart(this.id, this.option)
      })
    }
  }
</script>

<style>
  #txnCharts {
    height: 300px;
  }
</style>
