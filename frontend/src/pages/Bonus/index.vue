<template>
	<div class="voteWithdraw-box">

  	<div class="voteWithdraw-header">
      <div class="my-header-container">
				<h1>Bonus Withdraw</h1>
      </div>
    </div>
  	<div class="my-container">
  		<h3>投票奖励：</h3>
		<table class="table bonus-table">
	        <tbody>
	          <tr>
	            <td style="width: 300px;">节点名：</td>
	            <td>{{account}}</td>
	          </tr>
	          <tr>
	            <td>奖励：</td>
	            <td>{{bonusVal}}</td>
	          </tr>
	          <tr>
	            <td></td>
	            <td><a href="javascript:void(0)" class="btn btn-success" @click="withdrawVote">领取</a></td>
	          </tr>
	        </tbody>
	      </table>

	    <h3>造块奖励：</h3>
		<table class="table bonus-table">
	        <tbody>
	          <tr>
	            <td style="width: 300px;">节点名：</td>
	            <td>{{account}}</td>
	          </tr>
	          <tr>
	            <td>奖励：</td>
	            <td>{{bpVal}}</td>
	          </tr>
	          <tr>
	            <td></td>
	            <td><a href="javascript:void(0)" class="btn btn-success" @click="withdrawBp">领取</a></td>
	          </tr>
	        </tbody>
	      </table>
  	</div>
  </div>
</template>

<script type="text/javascript">
	import Swal from 'sweetalert2';
	import axios from 'axios';

	export default {
		name: 'Bonus',
		data() {
			return {
				account: '-',
				bonusVal: '-',
				bpVal: '-',
				iost: null
			}
		},

		methods: {
			delay(n) {
				n = n || 500;
				return new Promise(done => {
					setTimeout(() => {
						done();
					}, n);
				});
			},

			withdrawVote: function() {
				if (this.bonusVal <= 0) return;
				const tx = this.iost.callABI("vote_producer.iost", "candidateWithdraw", [this.account]);
				tx.setGas(1, 2000000);

				self = this;
				var txHash;
          		this.iost.signAndSend(tx)
            		.on('pending', function(txid) {
              			txHash = txid;
              			Swal.fire({
                			title: '正在发送',
			                html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txid}">${txid}</a>`,
			                onBeforeOpen: () => {
                  				Swal.showLoading() 
                			},
                			allowOutsideClick: () => !Swal.isLoading()
              			})
              			console.log("txid:", txid)
            		})
            		.on('success', function(result) {
              			Swal.close()
              			Swal.fire({
                			type: 'success',
                			title: '领取成功',
                			html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txHash}">${txHash}</a>`,
              			})
              			self.bonusVal = 0;
              			console.log("res:", result, txHash)
            		})
            		.on('failed', function(failed) {
              			Swal.close()
              			Swal.fire({
                			type: 'error',
                			title: '领取失败',
                			html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txHash}">${txHash}</a>`
              			})
              			console.log(failed)
            		})
            },

            withdrawBp: function() {
            	if (this.bpVal <= 0) return;
            	const tx = this.iost.callABI("bonus.iost", "exchangeIOST", [this.account, '0']);
				tx.setGas(1, 2000000);

				self = this;
				var txHash;
          		this.iost.signAndSend(tx)
            		.on('pending', function(txid) {
              			txHash = txid;
              			Swal.fire({
                			title: '正在发送',
			                html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txid}">${txid}</a>`,
			                onBeforeOpen: () => {
                  				Swal.showLoading() 
                			},
                			allowOutsideClick: () => !Swal.isLoading()
              			})
              			console.log("txid:", txid)
            		})
            		.on('success', function(result) {
              			Swal.close()
              			Swal.fire({
                			type: 'success',
                			title: '领取成功',
                			html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txHash}">${txHash}</a>`,
              			})
              			self.bonusVal = 0;
              			console.log("res:", result, txHash)
            		})
            		.on('failed', function(failed) {
              			Swal.close()
              			Swal.fire({
                			type: 'error',
                			title: '领取失败',
                			html: `txid: <a target=_blank href="https://www.iostabc.com/tx/${txHash}">${txHash}</a>`
              			})
              			console.log(failed)
            		})
            }
		},

		created() {
			const s = document.createElement('script');
			s.type = 'text/javascript';
			s.src = 'https://cdn.jsdelivr.net/npm/iost@0.1.2/iost.min.js';
			document.body.appendChild(s)
		},

		async mounted() {
			await this.delay(600);
			console.log("fuck it");
			if (typeof IWalletJS === 'undefined') {
				Swal.fire({
					type: 'error',
					text: '无法加载钱包，请确认Chrome插件已正确安装',
				})
			}
			self = this;

			IWalletJS.enable().then(function (account) {
				console.log(account)
				if(!account) {
					Swal.fire({
						type: 'error',
						text: '无法加载钱包，请确认私钥已正确导入'
					})
					return;
				}
				self.account = account;

				axios.get(`https://api.iost.io/getCandidateBonus/${account}/1`)
		        	.then((response) => {
		        		self.bonusVal = response.data.bonus / 2;
		          })

		       	axios.get(`http://api.iost.io/getTokenBalance/${account}/contribute/1`)
		       		.then((response) => {
		       			self.bpVal = response.data.balance;
		       		})

		       	self.iost = IWalletJS.newIOST(IOST);
		    })
		}
	}
</script>

<style lang="less" rel="stylesheet/less">
	.voteWithdraw-box {
		padding-bottom: 160px;
		background: #F6F7F8;
		.voteWithdraw-header {
			background: #F6F7F8;
			box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
			.my-header-container {
				display: flex;
				height: 85px;
				width: 1000px;
				margin: 0 auto;
				> h1 {
					margin-top: 21px;
					font-size: 36px;
					line-height: 44px;
					font-weight: bold;
				}
			}
		}
		.my-container {
			width: 1000px;
			margin: 24px auto 0;
			text-align: left;
			background: #FFFFFF;
			padding: 26px 120px 86px 80px;
			position: relative;
			box-shadow: 0 2px 3px rgba(0,0,0,0.1);
			.err-msg {
				margin-bottom: 60px;
				margin-left: 100px;
			}
			> div {
				&:nth-child(2) {
					margin-bottom: 60px;
					align-items: center;
				}
				&:nth-child(3) {
					margin-bottom: 75px;
				}
			}

			> button {
				width: 700px;
				height: 54px;
				margin-left: 100px;
				background-color: rgba(44,46,49,0.5);
				box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
				border: none;
				color: #FFFFFF;
				font-weight: bold;
				&.active {
					background-color: rgba(44,46,49,1);
				}
			}
		}
	}

</style>