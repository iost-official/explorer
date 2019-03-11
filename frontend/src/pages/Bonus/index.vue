<template>
	<div class="voteWithdraw-box">

  	<div class="voteWithdraw-header">
      <div class="my-header-container">
				<h1>Bonus Withdraw</h1>
      </div>
    </div>
  	<div class="my-container">
		<table class="table bonus-table">
	        <tbody>
	          <tr>
	            <td>节点名：</td>
	            <td>{{account}}</td>
	          </tr>
	          <tr>
	            <td>分红：</td>
	            <td>{{bonusVal}}</td>
	          </tr>
	          <tr>
	            <td></td>
	            <td><a href="javascript:void(0)" class="btn btn-success" @click="withdraw">领取</a></td>
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
			}
		},

		methods: {
			withdraw: function() {

            }
		},

		mounted() {
			self = this;

			IWalletJS.enable().then(function (account) {
				if(!account) {
					return;
				}
				self.account = account;

				axios.get(`http://api.iost.io/getCandidateBonus/${account}/1`)
		        	.then((response) => {
		        		self.bonus = response.data.bonus;
		          })
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
			padding: 56px 120px 86px 80px;
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