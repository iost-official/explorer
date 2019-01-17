<template>
  <div class="applyIost-box">

		<div class="applyIost-header">
      <div class="my-header-container">
				<h1>Create</h1>
      </div>
    </div>

		<div class="my-tips-container">
			<div class="" id="luckyBetIntro">
				<p>Lucky Bet is operating on the v1.0 Everest testnet of the IOST Blockchain. It provides a way for users to participate and try out the IOST testnet. Everest is still in its early testing phase. In order to launch a stable Mainnet, IOST engineers will be constantly pressure testing, adjusting, fixing and upgrading the testnet. This may result in instability, periods where the testnet is offline, and other abnormalities.</p>
				<p>If you encounter any issues or bugs, please email team@iost.io or report your issue via the following link:</p>
				<router-link to="/feedback">https://explorer.iost.io/feedback</router-link>
				<p class="cn">此页面用于创建IOST主网账号，单IP每天最多创建5个账号。</p>
				<p>如果你发现了测试网络的任何Bug，欢迎发邮件给team@iost.io，或在以下网址提交给我们改进：</p>
				<!--<a href="https://explorer.iost.io/#/feedback">https://explorer.iost.io/#/feedback</a>-->
				<router-link to="/feedback">https://explorer.iost.io/feedback</router-link>
			</div>
		</div>

		<div class="my-container">
			<div class="my-err-msg" role="alert" id="errAlert">{{errMsg}}</div>
			<div class="my-group">
				<p class="">Account PubKey</p>
				<input class="my-input col-sm-8" placeholder="Account PubKey" id="applyAddress" v-model.trim()="address">
			</div>
			
			<div class="my-group my-check">
				<input type="checkbox" id="generateAddressCheck" v-model="checked">
				<p>Generate KeyPair for me</p>
			</div>

			<div class="my-group">
				<p class="">Account Name</p>
				<input class="my-input col-sm-8" placeholder="Account Name" id="applyName" v-model.trim()="accountName">
			</div>

			<div class="my-group">
				<p>Email</p>
				<input type="email" class="my-input" placeholder="Email" v-model.trim()="email">
			</div>

			<div class="my-group my-recaptcha">
				<div class="g-recaptcha" id="recap" data-sitekey="6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"></div>
			</div>

			<button type="button" :class="{active: isOK}" id="RequestBuuton" @click="apply">Create</button>
		</div>

		<div class="modal fade bs-example-modal-sm" id="applyIOSTModal" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
	  <div class="modal-dialog modal-sm" role="document">
	    <div class="modal-content">
	      <div class="modal-body">
	      	<p></p>
        	<p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Creating Account<span> {{dott}}</span></p>
        	<p></p>
     	  </div>
	    </div>
	  </div>
	</div>
  </div>
</template>

<script>
import axios from 'axios';
import swal from 'sweetalert2'
import base58 from 'bs58'
import elliptic from 'elliptic'
import { config } from '../../utils/config'

const { apis } = config

export default {
	name: 'ApplyIOST',
	data() {
		return {
			address: '',
			addressx: '',
			auto: 0,
			privKey: '',
			pubKey: '',
			amount: '',
			email: '',
			mobile: '',
			verify: '',
			mobileButtonMsg: 'Send',
			ltime: 0,
			errMsg: '',
			checked: false,
			dott: '',
			accountName: '',
		}
	},

	computed: {
		isOK () {
			if (this.address != '' && this.email != '' && this.verify != '' && this.mobile != '') {
				return true
			} else {
				return false
			}
		}
	},

	methods: {
	    // 替换后的
	    randomBytes: function(size) {
			let rawBytes = new Uint8Array(size)
			let bytes

			crypto.getRandomValues(rawBytes)
			bytes = Buffer.from(rawBytes.buffer)

			return bytes
	    },

		bytesToHex: function(bytes) {
			return bytes.toString('hex')
		},

		getApplyParam(grecap) {
			var params = new URLSearchParams();
			if (this.auto == 1) {
				// params.append('address', this.addressx)
				params.append('address', this.address)
			} else {
				params.append('address', this.address)
			}
			params.append('account', this.accountName)
			params.append('email', this.email)
			params.append('gcaptcha', grecap)
			return params
		},

		apply: function() {
			if (this.address.length < 35) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid account pubKey'
				return false
			}

			if (this.accountName.length < 5 || this.accountName.length > 11) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid account name, must 5-11 characters'
				return false
			}

			var regExp = /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/
			if (this.email.length == 0 || !regExp.test(this.email)) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid email address'
				return false
			}

			var grecap = grecaptcha.getResponse()
			if (grecap.length == 0) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid google reCaptcha check'
				return false
			}

			swal({
				html: '',
				title: 'Creating Account...',
				onOpen: () => {
					swal.showLoading()
				},
				allowOutsideClick: () => !swal.isLoading()
			})
			let dotNum = 1
			let interval = window.setInterval(() => {
				if (dotNum == 6) {
					this.dott = '......'
					
					dotNum = 0
				} else if (dotNum == 5) {
					this.dott = '.....'
				} else if (dotNum == 4) {
					this.dott = '....'
				} else if (dotNum == 3) {
					this.dott = '...'
				} else if (dotNum == 2) {
					this.dott = '..'
				} else if (dotNum == 1) {
					this.dott = '.'
				} else if (dotNum == 0) {
					this.dott = ''
				}
				dotNum++
			}, 1000)

			axios.post(`${apis.applyIOST}`, this.getApplyParam(grecap)).then((response) => {
				let retCode = response.data.ret
				let txHash = response.data.msg
				if (retCode != 0) {
					swal.close()
					$('#applyIOSTModal').modal('hide')
					$('#errAlert').addClass('alert alert-danger')
					this.errMsg = response.data.msg
				} else {
					swal.close()
					$('#applyIOSTModal').modal('hide')
					let pushParams = {}
					if (this.address.length > 0) {
						pushParams.address = this.address
					} else {
						// pushParams.address = this.addressx
						pushParams.address = this.address
					}
					pushParams.email = this.email
					pushParams.accountName = this.accountName
					pushParams.privKey = this.privKey
					pushParams.txHash = txHash
					this.$router.push({
						name: 'ApplyIOSTSuccess',
						params: pushParams
					})
				}
			})
		}
	},

	watch: {
		checked: function() {
			if (this.checked == false) {
				this.privKey = ''
				this.pubKey = ''
				this.address = ''
				$('#applyAddress').attr('disabled', false)
			} else {
				$('#applyAddress').attr('disabled', true)

				let bytes = this.randomBytes(32)
		        const EdDSA = elliptic.eddsa;
		        const ec = new EdDSA('ed25519');
		        const key = ec.keyFromSecret(bytes);

		        let pubKey = key.pubBytes()

		        let privKey = [...bytes]
		        privKey.push(...pubKey)

				this.privKey = base58.encode(Buffer.from(privKey))
				this.pubKey = base58.encode(Buffer.from(pubKey))
				this.address = base58.encode(Buffer.from(pubKey))

				console.log(this.pubKey)
				console.log(this.privKey)
				console.log(privKey)

		        this.auto = 1
			}
		}
	},

	mounted: function() {
		$("#phone").intlTelInput()
		let rc = document.getElementsByClassName('g-recaptcha')[0].children
		setTimeout(() => {
			grecaptcha.render('recap',{"sitekey": "6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"})
			if (rc.length == 0) {
				swal({
					title: 'Failed to load google reCaptcha !',
					text: '',
					confirmButtonText: 'try again'
				}).then((result) => {
					grecaptcha.render('recap',{"sitekey": "6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"})
				})
			}
		},2000)
	}
}
</script>


<style lang="less" rel="stylesheet/less">
	.applyIost-box {
		padding-bottom: 160px;
		background: #F6F7F8;
		.applyIost-header {
			background: #F6F7F8;
			box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
			.my-header-container {
				width: 1000px;
				height: 85px;
				margin: 0 auto;
				text-align: left;
				overflow: hidden;
				> h1 {
					margin-top: 21px;
					font-size: 36px;
					line-height: 44px;
					font-weight: bold;
				}
			}
		}
		.my-tips-container {
			width: 1000px;
			margin: 24px auto 0;
			padding: 30px 50px 38px 50px;
			background: #F2DEDE;
			border: 1px solid #EBCCD1;
			box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
			text-align: left;
			font-weight: 600;
			p {
				font-size: 14px;
				line-height: 22px;
				color: #A94442;
				&.cn {
					margin-top: 20px;
				}
			}
			a {
				font-size: 14px;
				line-height: 18px;
				color: #5EBFD9;
			}
		}

		.my-container {
			width: 1000px;
			margin: 24px auto 0;
			text-align: left;
			background: #FFFFFF;
			padding: 56px 120px 70px 80px;
			position: relative;
			box-shadow: 0 2px 3px rgba(0,0,0,0.1);
			.my-err-msg {
				padding: 15px;
				margin-bottom: 20px;
				border-radius: 4px;
				margin-left: 110px;
				text-align: center;
			}
			.my-group{
				display: flex;
				align-items: center;
				justify-content: space-between;
				margin-bottom: 60px;
				&:nth-child(2) {
					margin-bottom: 0;
				}
				&.my-check {
					margin-bottom: 24px;
					margin-left: 112px;
					justify-content: normal;
					> p {
						margin: 12px 0 0 8px;
						width: auto;
						font-size: 13px;
						line-height: 24px;
					}
					> input {
						margin: 14px 0 0;
						height: auto;
					}
				}

				&.my-mobile {
					/*这是脚本添加的类名*/
					margin-bottom: 12px;
					.intl-tel-input {
						display: flex;
						justify-content: space-between;
						width: 700px;
						input {
							padding-left: 60px;
						}
					}
				}
				&.my-recaptcha {
					margin-bottom: 30px;
					margin-left: 110px;
				}
				&.my-verify {
					.input-btn {
						display: flex;
						justify-content: space-between;
						width: 700px;
						> input {
							width: 500px;
						}
						> button {
							width: 196px;
							border-width: 0;
							background-color: #F6F7F8;
							&:hover {
								border-width: 1px;
							}
						}
					}

				}
				.my-input {
					width: 690px;
					background-color: #F6F7F8;
					outline: none;
					border: 0;
					padding: 6px 12px;
					color: #2c3e50;
				}
				p {
					margin-bottom: 0;
					width: 120px;
					text-align: right;
					margin-left: -20px;
					font-size: 14px;
					line-height: 18px;
					font-weight: bold;
				}
				input {
					height: 54px;
				}
			}

			> button {
				width: 700px;
				height: 54px;
				margin-left: 110px;
				background-color: rgba(44,46,49,0.5);
				border: none;
				color: #FFFFFF;
				font-weight: bold;
				box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
				&.active {
					background-color: rgba(44,46,49,1);
				}
			}
		}
	}
	@media screen and (max-width:480px) {
		.applyIost-box {
			padding-bottom: 24px;
			.applyIost-header {
				.my-header-container {
					width: 100%;
					padding: 0 25px;
				}
			}
			.my-tips-container {
				width: 100%;
				padding: 20px 25px;
			}

			.my-container {
				width: 100%;
				margin: 24px auto 0;
				text-align: left;
				padding: 30px 25px;
				.my-err-msg {
					margin-left: 80px;
				}
				.my-group{
					margin-bottom: 24px;
					&.my-check {
						margin-left: 80px;
					}

					&.my-mobile {
						/*这是脚本添加的类名*/
						.intl-tel-input {
							width: 285px;
							input {
								padding-left: 60px;
							}
						}
					}
					&.my-recaptcha {
						margin-left: 78px;
					}
					&.my-verify {
						> p {
							width: 78px;
						}
						.input-btn {
							width: 285px;
							> input {
								width: 185px;
							}
							> button {
								width: 90px;
							}
						}

					}
					.my-input {
						width: 350px;
					}
					p {
						text-align: left;
						margin-left: 0;
					}
				}
				> button {
					width: 285px;
					margin-left: 78px;
				}
			}
		}
	}

	@media screen and (max-width:375px) {
		.applyIost-box {
			.my-container {
				.my-err-msg {
					margin-left: 70px;
				}
				.my-group{
					&.my-mobile {
						/*这是脚本添加的类名*/
						.intl-tel-input {
							width: 254px;
						}
					}
					&.my-recaptcha {
						margin-left: 0;
					}
					&.my-verify {
						.input-btn {
							width: 250px;
							> input {
								width: 170px;
							}
							> button {
								width: 70px;
							}
						}
					}
				}
				> button {
					width: 254px;
					margin-left: 72px;
				}
			}
		}
	}
</style>