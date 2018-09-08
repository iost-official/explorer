<template>
  <div class="accountDetail-box">
    <div class="luckyBet-box">
      <img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet">Play Lucky Bet !</a>
    </div>

    <div class="accountDetail-header">
      <div class="my-header-container">
        <h1>Address</h1>
        <p>{{address}}</p>
      </div>
    </div>


    <div class="accountDetail-info">
      <div class="accountDetail-img">
        <img src="../../assets/iostWhite.png" alt="">
        <div class="accountDetail-height-value">
          <div class="accountDetail-height">
            <h4>Balance:</h4>
            <p>{{accountDetail.balance}}</p>
          </div>
          <div class="accountDetail-value">
            <h4>IOST Value:</h4>
            <p></p>
          </div>
        </div>
      </div>
      <div class="accountDetail-view">
        <table class="table account-overview">
          <tbody>
          <tr>
            <td>Transactions:</td>
            <td>{{txnLen}}</td>
          </tr>
          </tbody>
        </table>
      </div>
      <div class="accountDetail-listDetail">
        <ul class="nav nav-tabs" id="txnDetailList">
          <li role="presentation" class="active"><a href="#txns-nav-txns">Transactions</a></li>
          <li id="CCodeDisplay" role="presentation"><a href="#txns-contract-code">Contract Code</a></li>
        </ul>
      </div>
      <div class="tab-content my-tab-content">
        <div class="tab-pane active row address-txns-head my-tab-pane" id="txns-nav-txns">
          <div style="line-height: 80px">
            <div class="col-md-8">
              <i aria-hidden="true" class="c333 fa fa-sort-amount-down"></i> Latest 25 txns from a total Of <a :href="'/txs?a=' + address">{{txnLen}} transactions</a>
            </div>
            <div class="col-md-4">
              <a :href="'/txs?a=' + address">View All</a>
            </div>
          </div>



          <div class="accountDetail-list">
            <ul class="my-list-header">
              <li>TxHash</li>
              <li>Block</li>
              <li>Age</li>
              <li>From</li>
              <li>To</li>
              <li>Value</li>
              <li>[TxFee]</li>
            </ul>
            <div class="list-wrap">
              <div class="list-body-wrap" v-for="tx in accountTxnInfo.txnList">
                <ul class="my-list-body">
                  <li><a :href="'/tx/' + tx.hash">{{tx.hash}}...</a></li>
                  <li><a :href="'/block/' + tx.blockNumber">{{tx.blockNumber}}</a></li>
                  <li>{{tx.age}}</li>
                  <li><a :href="'/account/' + tx.from">{{tx.from}}</a></li>
                  <li><a :href="'/account/' + tx.to">{{tx.to}}</a></li>
                  <li>{{tx.amount}}</li>
                  <li>{{tx.gasPrice}}</li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <div class="tab-pane" id="txns-contract-code">
          <table class="table ContractCodeTable">
            <tbody>
            <tr>
              <td>Publisher:</td>
              <td>{{ContractInfo.from}}</td>
            </tr>
            <tr>
              <td>Time:</td>
              <td>{{ContractInfo.time}}</td>
            </tr>
            <tr>
              <td>Code:</td>
              <td><pre>{{ContractInfo.code}}</pre></td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import { mapState } from 'vuex'

  export default {
    name: "AccountDetail",
    data() {
      return {
        address: this.$route.params.id,
        // account: {},
        txnList: {},
        txnLen: '',
        ContractInfo: {},
      }
    },

    computed: {
      ...mapState(['accountDetail', 'accountTxnInfo'])
    },

    methods: {

      fetchData(r) {
        this.address = r.params.id

        this.$store.dispatch('getAccountDetail', this.address)

        this.$store.dispatch('getAccountTxnInfo', this.address)


        // axios.get('https://explorer.iost.io/api/account/' + this.address).then((response) => {
        //   this.account = response.data
        // })
        // axios.get('https://explorer.iost.io/api/account/' + this.address + '/txs').then((response) => {
        //   this.txnList = response.data.txn_list
        //   this.txnLen = response.data.txn_len
        // })

        // axios.get('https://explorer.iost.io/api/tx/' + this.address).then((response) => {
        axios.get('http://47.75.223.44:8080/api/tx/' + this.address).then((response) => {
          if (response.data.code == 1) {
            $('#CCodeDisplay').hide()
            return
          }
          $('#CCodeDisplay').show()
          let blkDate = new Date(response.data.time / 1000 / 1000)
          let time = ''
          time += blkDate.getFullYear() + '-'
          time += (blkDate.getMonth()+1 < 10 ? '0'+(blkDate.getMonth()+1) : blkDate.getMonth()+1) + '-'
          time += blkDate.getDate() + ' ';
          if (blkDate.getHours() < 10) {
            time += '0' + blkDate.getHours()
          } else {
            time += blkDate.getHours()
          }
          if (blkDate.getMinutes() < 10) {
            time += ':0' + blkDate.getMinutes()
          } else {
            time += ':' + blkDate.getMinutes()
          }
          if (blkDate.getSeconds() < 10) {
            time += ':0' + blkDate.getSeconds()
          } else {
            time += ':' + blkDate.getSeconds()
          }
          response.data.time = time
          this.ContractInfo = response.data
        })
      }
    },



    watch: {
      '$route': function (r) {
        this.fetchData(r)
      }
    },
    mounted: function () {
      $("#txnDetailList a").click(function (e) {
        e.preventDefault()
        $(this).tab('show')
      })
      this.fetchData(this.$route)
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .accountDetail-box {
    padding-top: 90px;
    margin: 0 auto;
    background: #F6F7F8;
    .luckyBet-box {
      background: #2C2E31;
      height: 50px;
      line-height: 50px;
      color: #F6F7F8;
      font-size: 14px;
      > img {
        width: 24px;
        height: 24px;
        margin-right: 12px;
      }
      a {
        color: #F6F7F8;
        font-size: 14px;
        line-height: 18px;
        text-decoration: none;
      }
    }
    .accountDetail-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        width: 1000px;
        height: 120px;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        > h1 {
          margin: 21px 0 15px;
          color: #273238;
          font-size: 36px;
          line-height: 44px;
        }
        > p {
          margin: 0;
          font-size: 14px;
          line-height: 18px;
        }
      }
    }
    .accountDetail-info {
      width: 1000px;
      margin: 24px auto 0;
      background: #FFFFFF;
      height: 805px;
      text-align: left;
      .accountDetail-img {
        > img {
          height: 50px;
        }
        padding: 15px 100px 0 50px;
        box-shadow: 0 5px 5px 0 rgba(0, 0, 0, .1);
      }
      h4 {
        font-size: 14px;
        line-height: 18px;
        color: #2C2E31;
        padding-bottom: 12px;
        padding-top: 20px;
        border-bottom: 1px solid #F1F1F1;
      }
      .accountDetail-height-value {
        display: flex;
        .accountDetail-height, .accountDetail-value {
          width: 50%;
        }
      }
      p {
        font-size: 18px;
        line-height: 22px;
        color: #2C2E31;
        a {
          font-size: 18px;
          line-height: 22px;
          color: #2C2E31;
        }
      }

      .my-tab-content {
        .my-tab-pane {
          margin: 0;
          .accountDetail-list {
            padding: 0 20px;
            .my-list-header {
              font-size: 16px;
              font-weight: 500;
              /*> li {*/
                /*list-style: none;*/
                /*text-align: left;*/
                /*color: #333333;*/
                /*a {*/
                  /*color: #333333;*/
                  /*text-decoration: none;*/
                /*}*/
                /*&:first-child {*/
                  /*width: 400px;*/
                  /*text-overflow: ellipsis;*/
                  /*overflow: hidden;*/
                /*}*/
                /*&:nth-child(2) {*/
                  /*width: 100px;*/

                /*}*/
                /*&:nth-child(3) {*/
                  /*width: 200px;*/
                /*}*/
                /*&:nth-child(4) {*/
                  /*width: 130px;*/
                /*}*/
              /*}*/
            }
            .list-wrap {
              max-height: 400px;
              overflow: auto;
              &::-webkit-scrollbar {
                width: 0;
              }
              .list-body-wrap {
                width: 100%;
                &:hover {
                  box-shadow: 0 8px 30px 0 rgba(0, 0, 0, .15);
                }
                .my-list-body {
                  /*> li {*/
                    /*list-style: none;*/
                    /*text-align: left;*/
                    /*color: #333333;*/
                    /*a {*/
                      /*color: #333333;*/
                      /*text-decoration: none;*/
                    /*}*/
                    /*&:first-child {*/
                      /*width: 400px;*/
                      /*text-overflow: ellipsis;*/
                      /*overflow: hidden;*/
                    /*}*/
                    /*&:nth-child(2) {*/
                      /*width: 100px;*/

                    /*}*/
                    /*&:nth-child(3) {*/
                      /*width: 200px;*/
                    /*}*/
                    /*&:nth-child(4) {*/
                      /*width: 130px;*/
                    /*}*/
                  /*}*/
                }

              }
            }

            ul {
              display: flex;
              padding: 15px 0;
              width: 1000px;
              margin: 0 auto;
              li {
                list-style: none;
                text-align: left;
                color: #333333;
                a {
                  color: #333333;
                  text-decoration: none;
                }
                &:first-child {
                  width: 200px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }
                &:nth-child(2) {
                  width: 100px;

                }
                &:nth-child(3) {
                  width: 100px;
                }
                &:nth-child(4) {
                  width: 150px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }
                &:nth-child(5) {
                  width: 150px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }
                &:nth-child(6) {
                  width: 130px;
                }
                &:nth-child(7) {
                  width: 130px;
                }
              }
            }
          }
        }
      }

    }
  }
</style>
