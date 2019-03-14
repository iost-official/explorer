<template>
  <div class="success-box">

  	<div class="success-header">
			<div class="my-header-container">
				<h1>Request successfully processed</h1>
			</div>
		</div>

		<div class="success-container1" v-if="privKey != ''">
			<p class="success-title">Save your Account Info</p>
			<p class="success-address"><span>Account Name:</span> {{accountName}}</p>
			<p class="success-private"><span>Private Key:</span> {{privKey}}</p>
			<p class="success-tips">**Do not lose it!** It cannot be recovered if you lose it.</p>
			<p class="success-tips">**Do not share it!** Your funds will be stolen if you use this private key on a malicious/phishing site.</p>
		</div>

    <div class="success-container2">
			<div class="container2-hash">
				<h4>TxHash:</h4>
				<p><router-link :to="{path:`/tx/${txHash}`}" target="_blank">{{txHash}}</router-link></p>
			</div>
			<div class="container2-address">
				<h4>Account Name:</h4>
				<p>{{accountName}}</p>
			</div>
			<div class="container2-email">
				<h4>Email:</h4>
				<p>{{email}}</p>
			</div>

    </div>
  </div>
</template>

<script>
  export default {
		name: 'ApplyIOSTSuccess',
		data() {
			return {
				address: this.$route.params.address,
				accountName: this.$route.params.accountName,
				email: this.$route.params.email,
				txHash: this.$route.params.txHash,
				privKey: this.$route.params.privKey,
			}
		},
		created: function() {
			// if (!this.$route.params.txHash) {
			// 	this.$router.push({
			// 		path: '/applyIOST',
			// 	})
			// }
		}
  }
</script>

<style lang="less" rel="stylesheet/less">
	.success-box {
		padding-bottom: 160px;
		background: #F6F7F8;
		.success-header {
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

		.success-container1 {
			width: 1000px;
			margin: 24px auto 0;
			text-align: left;
			background: #FFFFFF;
			padding: 50px 100px;
			position: relative;
			box-shadow: 0 2px 3px rgba(0,0,0,0.1);
			p {
				/*font-size: 18px;*/
				font-size: 14px;
				line-height: 22px;
				margin-bottom: 0;
				font-weight: 300;
				text-align: center;
				a {
					color: #4678aa;
				}
			}
			.success-title {
				font-size: 24px;
				line-height: 29px;
				color: #2C2E31;
				margin-bottom: 50px;
			}
			.success-address, .success-private {
				height: 54px;
				margin-bottom: 40px;
				color: #000000;
				font-weight: bold;
				line-height: 54px;
				> span {
					font-weight: normal;
				}
			}
			.success-address {
				background-color: #98D5BD;
			}
			.success-private {
				height: auto;
				line-height: 27px;
				background-color: #5EBFD9;
				padding: 14px 0;

			}
			.success-tips {
				font-size: 14px;
				color: #A94442;
			}
		}

		.success-container2 {
			width: 1000px;
			margin: 24px auto 0;
			text-align: left;
			background: #FFFFFF;
			padding: 15px 50px 60px 50px;
			position: relative;
			box-shadow: 0 2px 3px rgba(0,0,0,0.1);
			> div {
				margin-bottom: 60px;
				&:last-child {
					margin-bottom: 0;
				}
				h4 {
					font-size: 14px;
					line-height: 18px;
					font-weight: bold;
					padding-bottom: 12px;
					margin: 0;
					color: #2c2e31;
					border-bottom: 1px solid #f6f7f8;
					&:last-child {
						padding-bottom: 60px;
					}
				}
				p {
					/*font-size: 18px;*/
					font-size: 14px;
					line-height: 22px;
					margin-top: 20px;
					margin-bottom: 0;
					font-weight: 300;
				}
			}
		}
	}
</style>