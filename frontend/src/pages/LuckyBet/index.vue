<template>
  <div>
    <div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Lucky Bet</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Bet</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" id="luckyBetIntro">
        <div class="alert alert-danger" role="alert">
          <p>Lucky Bet is operating on the v0.5 Everest testnet of the IOST Blockchain. It provides a way for users to participate and try out the IOST testnet. Everest is still in its early testing phase. In order to launch a stable Mainnet, IOST engineers will be constantly pressure testing, adjusting, fixing and upgrading the testnet. This may result in instability, periods where the testnet is offline, and other abnormalities.</p>

          <!--<p>If you encounter any issues or bugs, please email team@iost.io or report your issue via the following link: <a href="https://explorer.iost.io/#/feedback">https://explorer.iost.io/#/feedback</a></p>-->
          <p>If you encounter any issues or bugs, please email team@iost.io or report your issue via the following link:<router-link to="/feedback">https://explorer.iost.io/#/feedback</router-link>
          </p>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" style="padding-top: 10px;">
        <div class="col-md-2" style="padding-left: 0">
          <div class="panel panel-default bet-pannel betLuckyNumberPannel">
            <div class="panel-heading">
              <h3 class="panel-title">Last Lucky Number</h3>
            </div>
            <div class="panel-body">
              <p v-if="top5ResultList.length > 0">{{top5ResultList[0].BlockHeight.toString()[top5ResultList[0].BlockHeight.toString().length - 1]}}<span> ({{top5ResultList[0].BlockHeight}})</span></p>
            </div>
          </div>
          <div class="panel panel-default bet-pannel">
            <div class="panel-heading">
              <h3 class="panel-title">Current Block</h3>
            </div>
            <div class="panel-body">
              <div id="betBlocks">
                27681
              </div>
            </div>
          </div>
          <div class="panel panel-default bet-pannel">
            <div class="panel-heading">
              <h3 class="panel-title">Overview</h3>
            </div>
            <div class="panel-body">
              <div id="betOverView">
                <p><a href="javascript:void(0)" @click="betSearchModal">Bet Search</a></p>
                <p><a href="javascript:void(0)" @click="todayRanking">Today's Ranking</a></p>
                <p><a href="">History Payment</a></p>
              </div>
            </div>
          </div>
        </div>
        <div class="col-md-10">
          <div class="panel panel-default bet-pannel last5ResultPannel">
            <div class="panel-heading">
              <h3 class="panel-title">Game Rules</h3>
            </div>
            <div class="panel-body">
              <p>1. Participants can choose to bet between 1-5 Test IOST on a number of their choice from 0 to 9.</p>
              <p>2. Every time a bet is placed, 0.01 Test IOST will be deducted as a transaction fee.</p>
              <p>3. After 100 participants have placed their bets, the network will generate a winning number.</p>
              <p>4. The winning number is determined by the last digit of the block height when the 100th user participated.</p>
              <p>5. The total prize pool is 95% of all the bets placed, with the remaining 5% used for Gas.</p>
              <p>6. The total prize pool will be distributed among the winners based on their respective bets.</p>
            </div>
          </div>
          <div class="panel panel-default bet-pannel last5ResultPannel">
            <div class="panel-heading">
              <h3 class="panel-title">Last 5 Bet Results</h3>
            </div>
            <div class="panel-body">
              <table class="table table-striped">
                <thead>
                <tr>
                  <th>Round</th>
                  <th>Block Height</th>
                  <th>Lucky Number</th>
                  <th>Total Winners</th>
                  <th>Total Rewards</th>
                  <th>Time</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(betInfo, index) in top5ResultList">
                  <td><a href="javascript:void(0)" @click="showRoundInfo(betInfo.round, betInfo.BlockHeight, betInfo.WinUserNumber, betInfo.TotalRewards, betInfo.win_time)">{{betInfo.round}}</a></td>
                  <td>{{betInfo.BlockHeight}}</td>
                  <td>{{betInfo.BlockHeight.toString()[betInfo.BlockHeight.toString().length-1]}}</td>
                  <td>{{betInfo.WinUserNumber}}</td>
                  <td>{{Math.round(betInfo.TotalRewards/1000000)/100}}</td>
                  <td>{{betInfo.win_time}}</td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="panel panel-default bet-pannel joinBetPannel">
            <div class="panel-heading">
              <h3 class="panel-title">Join Lucky Bet !</h3>
            </div>
            <div class="panel-body">
              <div class="" role="alert" id="errBetAlert">{{errMsg}}</div>
              <form class="form-horizontal">
                <div class="form-group">
                  <label class="col-sm-2 control-label">Amount</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" placeholder="1-5" v-model.trim()="betAmount">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 control-label">Lucky Number</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" placeholder="0-9" v-model.trim()="luckyNumber">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 control-label">Address</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" placeholder="address" v-model.trim()="address">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 control-label">Private Key</label>
                  <div class="col-sm-10">
                    <input type="password" class="form-control" placeholder="account private key" v-model.trim()="privKey">
                  </div>
                </div>
                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <div class="g-recaptcha" id="recap2" data-sitekey="6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"></div>
                  </div>
                </div>
                <div class="form-group">
                  <div class="col-sm-offset-2 col-sm-10">
                    <button type="button" class="btn btn-primary" id="joinBetButton" @click="joinBet" style="min-width:90px; text-align: left;">{{betButtonMsg}}</button>
                    <a :href="'/#/applyIOST'">Do not have test IOST?</a>
                  </div>
                </div>
              </form>
            </div>
          </div>

          <div class="panel panel-default bet-pannel myBetPannel">
            <div class="panel-heading">
              <h3 class="panel-title">My Bets</h3>
            </div>
            <div class="panel-body">
              <table class="table">
                <thead>
                <tr>
                  <th>Round</th>
                  <th>Amount</th>
                  <th>Lucky Number</th>
                  <th>Time</th>
                  <th>Result</th>
                </tr>
                </thead>
                <tbody id="myBetTBody">
                <tr v-for="(betInfo, index) in betResultList">
                  <td v-if="betInfo.result==null">pending...</td>
                  <td v-else>{{betInfo.result.round}}</td>
                  <td>{{(betInfo.bet_amount/100000000).toFixed(2)}}</td>
                  <td>{{betInfo.lucky_number}}</td>
                  <td>{{betInfo.bet_time}}</td>
                  <!--<td v-if="betInfo.result==null">pending...</td>-->
                  <!--<td v-else-if="betInfo.result.is_win==false">not win</td>-->
                  <!--<td v-else>-->
                    <!--win <span style="color:#f00">{{betInfo.result.Amount/100000000).toFixed(2)}}</span> IOST-->
                  <!--</td>-->
                  <td v-if="betInfo.result==null && betInfo.status == 0">waiting for confirm</td>
                  <td v-else-if="betInfo.result==null && betInfo.status == 1">pending</td>
                  <td v-else-if="betInfo.result==null && betInfo.status == -1">confirm failed</td>
                  <td v-else-if="!betInfo.result.is_win">not win</td>
                  <td v-else>
                  win <span style="color:#f00">{{(betInfo.result.Amount/100000000).toFixed(2)}}</span> IOST
                  </td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="modal fade bs-example-modal-lg" id="roundDetail" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">Lucky Bet</h4>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-10"></div>
              <div class="col-md-1"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="modal fade bs-example-modal-lg" id="todayRanking" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">Today's Ranking</h4>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-10"></div>
              <div class="col-md-1"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="modal fade bs-example-modal-lg" id="searchBet" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">Bet Search</h4>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-10">
                <form class="form-inline">
                  <div class="form-group">
                    <label class="sr-only" for="exampleInputEmail3">Address</label>
                    <input type="text" class="form-control" placeholder="Address" id="betSearchInput" style="min-width: 640px">
                  </div>
                  <button type="button" class="btn btn-default" id="betSearch">Search</button>
                </form>
                <div id="searchResults"></div>
              </div>
              <div class="col-md-1"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import base58 from 'bs58';
  import swal from 'sweetalert2'

  export default {
    name: 'LuckyBetPages',
    data() {
      return {
        maxBetBlkHeight: 0,
        betAmount: '',
        luckyNumber: '',
        address: '',
        privKey: '',
        errMsg: '',
        betButtonMsg: 'Place Bet',
        betResultList: [],
        top5ResultList: [],
        balance: '',
      }
    },
    methods: {
      showRoundInfo: function(round, blkHeight, winUserNumber, totalRewards, winTime) {
        // axios.get('https://explorer.iost.io/api/luckyBet/round/' + round).then((response) => {
        axios.get('/api/luckyBet/round/' + round).then((response) => {
          let roundInfoList = response.data

          $("#roundDetail").on('show.bs.modal', function(event) {
            let modal = $(this)
            modal.find('.modal-title').text("Lucky Bet Round " + round)
            let commRoundHtml = '<table class="table">'
            commRoundHtml += '<thead><tr><th colspan="2">Round Information</th></tr></thead>'
            commRoundHtml += '<tbody>'
            commRoundHtml += '<tr><td>Round:</td><td>' + round + '</td></tr>'
            commRoundHtml += '<tr><td>Block Height:</td><td>' + blkHeight + '</td></tr>'
            commRoundHtml += '<tr><td>Lucky Number:</td><td style="color:#f00; font-weight:600;">' + blkHeight.toString()[blkHeight.toString().length - 1] + '</td></tr>'
            commRoundHtml += '<tr><td>Total Winners:</td><td>' + winUserNumber + '</td></tr>'
            commRoundHtml += '<tr><td>Total Rewards:</td><td>' + totalRewards + '</td></tr>'
            commRoundHtml += '<tr><td>Time:</td><td>' + winTime + '</td></tr>'
            commRoundHtml += '</tbody>'
            commRoundHtml += '</table>'

            let winUserListHtml = '<table class="table">'
            if (roundInfoList && roundInfoList.length > 0) {
              winUserListHtml += '<thead><tr><th colspan="2">Winner List</th></tr></thead>'
              winUserListHtml += '<tbody>'
              winUserListHtml += '<tr><th>Address</th><th>Win Amount</th></tr>'
              for (let i = 0, len=roundInfoList.length; i < len; i++) {
                winUserListHtml += '<tr><td><a href=' + roundInfoList[i]._id + '"/#/account/" target="_blank">' + roundInfoList[i]._id + '</a></td><td>' + roundInfoList[i].totalWinIOST.toFixed(2) + '</td></tr>'
              }
              winUserListHtml += '</tbody>'
            } else {
              winUserListHtml += '<thead><tr><th colspan="2">No Winner In this Round</th></tr></thead>'
            }
            winUserListHtml += '</table>'

            modal.find('.modal-body .col-md-10').html(commRoundHtml + winUserListHtml)
          })
          $("#roundDetail").modal()
        })
      },
      joinBet: function() {
        swal.close()
        if (this.betAmount <= 0 || this.betAmount > 5 || this.betAmount % 1 !== 0) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid amount'
          return false
        }

        if (this.luckyNumber == '' || this.luckyNumber < 0 || this.luckyNumber > 9 || this.luckyNumber % 1 !== 0) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid luckyNumber'
          return false
        }

        if (this.address.length < 45) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid IOST address'
          return false
        }

        if (this.privKey.length !=87 || this.privKey.length !=88) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid IOST private Key'
          return false
        }

        var grecap = grecaptcha.getResponse()
        if (grecap.length == 0) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid google reCaptcha check'
          return false
        }

        let index = 1
        this.betButtonMsg = 'Sending'
        let interval = window.setInterval(() => {
          if (index == 0) {
            this.betButtonMsg = 'Sending'
          } else if (index == 1) {
            this.betButtonMsg = 'Sending.'
          } else if (index == 2) {
            this.betButtonMsg = 'Sending..'
          } else if (index == 3) {
            this.betButtonMsg = 'Sending...'
            index = 0
          }
          index++

        }, 1000)

        $('#joinBetButton').attr('disabled', true)
        $('#errBetAlert').removeClass('alert alert-danger')
        this.errMsg = ''
        swal({
          title: 'Placing Bet...',
          onOpen: () => {
            swal.showLoading()
          },
          allowOutsideClick: () => !swal.isLoading()
        })
        grecaptcha.reset()

        var params = new URLSearchParams();
        params.append('address', this.address)
        params.append('privKey', this.privKey)
        params.append('betAmount', this.betAmount*100000000)
        params.append('luckyNumber', this.luckyNumber)
        params.append('gcaptcha', grecap)

        // axios.post('https://explorer.iost.io/api/luckyBet', params).then((response) => {
        axios.post('/api/luckyBet', params).then((response) => {
          $('#joinBetButton').attr('disabled', false)
          window.clearInterval(interval)
          this.betButtonMsg = 'Place Bet'
          var retCode = response.data.ret
          var txHash = response.data.tx_hash
          swal.close()
          if (retCode != 0) {
            if (retCode == 6) {
              swal({
                type: 'error',
                title: response.data.msg,
                html: '<p style="font-size:15px;">balance: ' + response.data.balance + '</p>',
                closeOnClickOutside: false,
              })
            } else {
              swal({
                type: 'error',
                title: 'network error',
                closeOnClickOutside: false,
              })
            }
          } else {
            swal({
              title: 'bet success!',
              html: '<h5>txHash: <a href=' + txHash +'"/#/tx/" target="_blank">' + txHash + '</a></h5>',
            })
            this.cronAddressBet(this.address)
          }
        })
      },
      formatTime: function(timeStr) {
        let blkDate = new Date(timeStr * 1000)
        let time = ''
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
        return time
      },
      cronLatest: function() {
        // axios.get('https://explorer.iost.io/api/luckyBet/latestBetInfo').then((response) => {
        // axios.get('http://13.114.55.155/api/luckyBet/latestBetInfo').then((response) => {
        axios.get('/api/luckyBet/latestBetInfo').then((response) => {
          if (!response.data) {
            return
          }
          let blkHeight = response.data.block_height
          let blkTime = response.data.block_time
          let blkHtml =  '<p>' + blkHeight + '</p>'
          blkHtml +=  '<p>' + this.formatTime(blkTime) + '</p>'
          $('#betBlocks').html(blkHtml)
          let latestList = response.data.latest_bet_list
          if (latestList != null && latestList.length > 0) {
            this.top5ResultList = latestList
            for (let i = 0, len = this.top5ResultList.length; i < len; i++) {
              this.top5ResultList[i].win_time = this.formatTime(this.top5ResultList[i].win_time)
              this.top5ResultList[i].TotalRewards = this.top5ResultList[i].TotalRewards.toFixed(2)
            }
          }
        })
      },
      cronAddressBet: function(address) {
        if (this.address.length < 45) {
          $('#errBetAlert').addClass('alert alert-danger')
          this.errMsg = 'invalid IOST address'
          return false
        }

        // axios.get('https://explorer.iost.io/api/luckyBet/addressBet/' + this.address + '?t=t').then((response) => {
        // axios.get('http://13.114.55.155/api/luckyBet/addressBet/' + this.address + '?t=t').then((response) => {
        axios.get('/api/luckyBet/addressBet/' + this.address + '?t=t').then((response) => {
          let betResList = response.data.address_bet_list
          if (betResList == null || betResList.length == 0) {
            return
          }
          for (let i = 0, len = betResList.length; i < len; i++) {
            betResList[i].bet_time = this.formatTime(betResList[i].bet_time)
            if (betResList[i].result) {
              betResList[i].result.Amount = betResList[i].result.Amount.toFixed(2)
            }
          }
          this.betResultList = betResList
          this.balance =  response.data.balance
        })
      },
      todayRanking: function() {
        // axios.get('https://explorer.iost.io/api/luckyBet/todayRanking').then((response) => {
        axios.get('/api/luckyBet/todayRanking').then((response) => {
          let rankList = response.data


          $("#todayRanking").on('show.bs.modal', function(event) {
            let modal = $(this)
            if (!rankList || rankList.length == 0) {
              modal.find('.modal-body .col-md-10').html("<p>No Rankings Currently Available...</p>")
              return
            }


            let rankingHtml = '<table class="table">'
            rankingHtml += '<thead><tr><th>Rank</th><th>Address</th><th>Total Win</th><th>Participate Times</th></tr></thead>'
            rankingHtml += '<tbody>'
            for (let i = 0, len = rankList.length; i < len; i++) {
              rankingHtml += '<tr>'
              rankingHtml +=  '<td>' + (i+1) + '</td>'
              rankingHtml +=  '<td><a href=' + rankList[i]._id + '"/#/account/" target="_blank">' + rankList[i]._id + '</a></td>'
              rankingHtml +=  '<td>' + rankList[i].netEarn.toFixed(2) + ' IOST</td>'
              rankingHtml +=  '<td>' + rankList[i].totalWinTimes + '</td>'
              rankingHtml += '</tr>'
            }
            rankingHtml += '</tbody>'
            rankingHtml += '</table>'

            modal.find('.modal-body .col-md-10').html(rankingHtml)
          })
          $("#todayRanking").modal()
        })
      },
      betSearchModal: function() {
        $("#searchBet").on('show.bs.modal', function(event) {
          let modal = $(this)
          $("#betSearch").click(function() {
            let addressInput = $('#betSearchInput').val()
            if (addressInput.length != 44 && addressInput.length != 45) {
              return false
            }
            modal.find('#searchResults').html("<p>searching...</p>")
            // axios.get('https://explorer.iost.io/api/luckyBet/addressBet/' + addressInput).then((response) => {
            // axios.get('http://13.114.55.155/api/luckyBet/addressBet/' + addressInput).then((response) => {
            axios.get('/api/luckyBet/addressBet/' + addressInput).then((response) => {
              let betResList = response.data.address_bet_list
              if (betResList == null || betResList.length == 0) {
                return
              }
              let searchHtml = '<table class="table table-striped">'
              searchHtml += '<thead><tr><th>Round</th><th>Amount</th><th>Lucky Number</th><th>Time</th><th>Result</th></tr></thead>'
              searchHtml += '<tbody>'

              for (let i = 0, len = betResList.length; i < len; i++) {
                let blkDate = new Date(betResList[i].bet_time * 1000)
                let time = ''
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
                betResList[i].bet_time = time
                if (betResList[i].result) {
                  betResList[i].result.Amount = betResList[i].result.Amount.toFixed(2)
                }

                searchHtml += '<tr>'
                if (betResList[i].result == null) {
                  searchHtml += '<td>pending</td>'
                } else {
                  searchHtml += '<td>' + betResList[i].result.round + '</td>'
                }
                searchHtml += '<td>' + betResList[i].bet_amount + '</td>'
                searchHtml += '<td>' + betResList[i].lucky_number + '</td>'
                searchHtml += '<td>' + betResList[i].bet_time + '</td>'
                if (betResList[i].result == null) {
                  searchHtml += '<td>pending</td>'
                } else if (betResList[i].result.is_win == false){
                  searchHtml += '<td>not win</td>'
                } else {
                  searchHtml += '<td>win <span style="color:#f00">' + betResList[i].result.Amount + '</span> IOST</td>'
                }
              }
              searchHtml += '</tbody>'
              searchHtml += '</table>'


              modal.find('#searchResults').html(searchHtml)
            })
          })
        })
        $("#searchBet").modal()
      },
    },
    mounted: function() {
      window.setTimeout(function() {
        if ($('.g-recaptcha').html().length == 0) {
          // swal({
          //   title: 'Failed to load google reCaptcha !',
          //   text: '',
          //   confirmButtonText: 'try again'
          // }).then((result) => {
          //   location.reload()
          // })
          grecaptcha.render('recap2',{"sitekey": "6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"})


        }
      }, 3000)
      let interval = window.setInterval(() => {
        this.cronLatest()
        if (this.address) {
          this.cronAddressBet(this.address)
        }
      }, 3000)
    },
    created: function() {
      // if (window.location.href.substr(-2) !== '?r') {
      //   window.location = window.location.href + '?r';
      //   location.reload()
      // }
      // axios.get('https://explorer.iost.io/api/luckyBetBlockInfo').then((response) => {
      axios.get('/api/luckyBetBlockInfo').then((response) => {
        if (response.data.ret == 0) {
          let blkList = response.data.top_6_blk
          this.maxBetBlkHeight = blkList[0].height
          let blkHtml = ''
          for (let i = 0, len=blkList.length; i < len; i++) {
            let blkDate = new Date(blkList[i].timestamp * 1000)
            let time = ''
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
            blkHtml +=  '<p>' + blkList[i].height + '</p>'
            blkHtml +=  '<p>' + time + '</p>'
            break
          }
          $('#betBlocks').html(blkHtml)
        }
      })
      this.cronLatest()
    },
    watch: {
      address: function() {
        this.cronAddressBet(this.address)
      }
    }
  }
</script>

<style>
  .bet-pannel {
    text-align: left;
  }

  .betLuckyNumberPannel .panel-body {
    padding: 5px;
  }

  .betLuckyNumberPannel p {
    color: #f00;
    font-size: 36px;
    font-weight: 600;
    margin-left: 25%;
  }

  .betLuckyNumberPannel p span {
    color: #666;
    font-size: 25px;
  }

  #betBlocks p:first-child {
    font-size: 36px;
    font-weight: 600;
    color: #f00;
    text-align: center;
  }

  #betBlocks p:not(:first-child){
    text-align: right;
    font-size: 16px;
    color: #666;
  }

  .last5ResultPannel {
    min-height: 200px;
  }

  .last5ResultPannel .pannel-body {
    padding-top: 0;
  }

  .joinBetPannel a {
    margin-left: 10px;
  }

  .last5ResultPannel table th, .last5ResultPannel table td {
    text-align: center;
  }

  .last5ResultPannel table > tbody > tr > td:nth-child(3) {
    font-size: 16px;
    color: #f00;
    font-weight: 600;
  }

  #roundDetail table:first-child tr th, #roundDetail table:nth-child(2) thead th {
    font-weight: 800;
    font-size: 18px;
  }

  #roundDetail table:first-child tr > td:first-child {
    font-weight: 500;
  }

  #roundDetail table:first-child td {
    border-top: 0;
    border-bottom: 1px solid #ddd;
  }

  .myBetPannel table td, .myBetPannel table th {
    text-align: center;
  }

  #searchResults {
    margin-top: 10px;
  }

  #searchResults p {
    font-size: 16px;
  }
</style>