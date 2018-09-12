<template>
  <div class="applyIost-box">
		<div class="luckyBet-box">
			<img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
		</div>
  	<div class="applyIost-header">
      <div class="my-header-container">
				<h1>Request</h1>
      </div>
    </div>


		<div class="my-tips-container">
			<div class="" id="luckyBetIntro">
				<p>Lucky Bet is operating on the v0.5 Everest testnet of the IOST Blockchain. It provides a way for users to participate and try out the IOST testnet. Everest is still in its early testing phase. In order to launch a stable Mainnet, IOST engineers will be constantly pressure testing, adjusting, fixing and upgrading the testnet. This may result in instability, periods where the testnet is offline, and other abnormalities.</p>
				<p>If you encounter any issues or bugs, please email team@iost.io or report your issue via the following link:</p>
				<!--<a href="https://explorer.iost.io/#/feedback">https://explorer.iost.io/#/feedback</a>-->
				<router-link to="/feedback">https://explorer.iost.io/#/feedback</router-link>
				<p class="cn">Lucky Bet是基于IOST测试网络开发，目的是为了让大家体验IOST阶段性进展。IOST目前正在早期测试阶段，为了保证将来的主网功能能够如期稳定上线，IOST开发团队随时可能对测试网络进行升级、调整、修复问题和压力测试，可能在某些时刻会造成网络下线、不稳定和其他异常情况发生。</p>
				<p>如果你发现了测试网络的任何Bug，欢迎发邮件给team@iost.io，或在以下网址提交给我们改进：</p>
				<!--<a href="https://explorer.iost.io/#/feedback">https://explorer.iost.io/#/feedback</a>-->
				<router-link to="/feedback">https://explorer.iost.io/#/feedback</router-link>
			</div>
		</div>

		<div class="my-container">
			<div class="my-err-msg" role="alert" id="errAlert">{{errMsg}}</div>
			<div class="my-group">
				<p class="">Address</p>
				<input class="my-input" placeholder="Address" id="applyAddress" v-model.trim()="address">
			</div>
			
			<div class="my-group my-check">
				<input type="checkbox" id="generateAddressCheck" v-model="checked">
				<p>Generate Address</p>
			</div>
			
			<div class="my-group">
				<p class="">Amount</p>
				<input class="my-input" placeholder="10.1 Test IOST" readonly>
			</div>

			<div class="my-group">
				<p>Email</p>
				<input type="email" class="my-input" placeholder="email" v-model.trim()="email">
			</div>

			<div class="my-group my-mobile">
				<p class="">Mobile</p>
				<input type="tel" id="phone" class="my-input" v-model.trim()="mobile">
			</div>

			<div class="my-group my-recaptcha">
				<div class="g-recaptcha" id="recap" data-sitekey="6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"></div>
			</div>

			<div class="my-group my-verify">
				<p class="">Verification Code</p>
				<div class="input-btn">
					<input class="my-input" placeholder="Verification code" v-model.trim()="verify">
					<button type="button" id="sendSmS" class="btn btn-default" @click="sendSmS">{{mobileButtonMsg}}</button>
				</div>
			</div>

			<button type="button" :class="{active: isOK}" id="RequestBuuton" @click="apply">Request</button>
		</div>

		<div class="modal fade bs-example-modal-sm" id="applyIOSTModal" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
	  <div class="modal-dialog modal-sm" role="document">
	    <div class="modal-content">
	      <div class="modal-body">
	      	<p></p>
        	<p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Sending IOST to Your Address<span> {{dott}}</span></p>
        	<p></p>
     	  </div>
	    </div>
	  </div>
	</div>
  </div>
</template>

<script>
import axios from 'axios';
import secp256k1 from 'secp256k1'
import base58 from 'bs58'
import swal from 'sweetalert2'
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
			mobileButtonMsg: 'send',
			ltime: 0,
			errMsg: '',
			checked: false,
			dott: '',
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
		randomBytes: function(size) {
			let rawBytes = new Uint8Array(size)
			let bytes

			do {
				crypto.getRandomValues(rawBytes)
				bytes = Buffer.from(rawBytes.buffer)
			} while (!secp256k1.privateKeyVerify(bytes))

			return bytes
		},
		bytesToHex: function(bytes) {
			return bytes.toString('hex')
		},

		sendSmS: function() {
			if (this.ltime > 0) {
				return false
			}

			var grecap = grecaptcha.getResponse()
			if (grecap.length == 0) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid google reCAPTCHA check'
				return false
			}

			let code = $("#phone").intlTelInput("getSelectedCountryData").dialCode
			var params = new URLSearchParams();
			params.append('mobile', '+' + code + this.mobile)
			params.append('gcaptcha', grecap)

			// axios.post('https://explorer.iost.io/api/sendSMS', params).then((response) => {
			// axios.post('http://47.75.223.44:8080/api/sendSMS', params).then((response) => {
			axios.post(`${apis.sendSMS}`, params).then((response) => {
				var retCode = response.data.code
				if (retCode != 0) {
					$('#errAlert').addClass('alert alert-danger')
					this.errMsg = response.data.message
				}
			})

			this.ltime = 120
			this.mobileButtonMsg = 'retry in ' + this.ltime + 's.'
			$("#sendSmS").attr('disabled', true)
			let interval = window.setInterval(() => {
				if (--this.ltime <=0) {
					this.mobileButtonMsg = 'send'
					window.clearInterval(interval)
					$("#sendSmS").attr('disabled', false)
				} else {
					this.mobileButtonMsg = 'retry in ' + this.ltime + 's.'
				}
			}, 1000)
		},
		getApplyParam(grecap) {
			var params = new URLSearchParams();
			if (this.auto == 1) {
				// params.append('address', this.addressx)
				params.append('address', this.address)
			} else {
				params.append('address', this.address)
			}
			let code = $("#phone").intlTelInput("getSelectedCountryData").dialCode
			params.append('email', this.email)
			params.append('mobile', '+' + code + this.mobile)
			params.append('verify', this.verify)
			params.append('gcaptcha', grecap)
			return params
		},
		apply: function() {
			if (this.addressx.length == 0 && (this.address.length != 44 && this.address.length != 45)) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid IOST address'
				return false
			}

			var regExp = /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/
			if (this.email.length == 0 || !regExp.test(this.email)) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid email address'
				return false
			}

			if (this.mobile.length < 10) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid mobile number'
				return false
			}
			if (this.verify.length != 6) {
				$('#errAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid verification code'
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
			  title: 'Sending IOST to Your Address...',
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

			// axios.post('http://47.75.223.44:8080/api/applyIOST', this.getApplyParam(grecap)).then((response) => {
			axios.post(`${apis.applyIOST}`, this.getApplyParam(grecap)).then((response) => {
				let retCode = response.data.code
				let txHash = response.data.data
				if (retCode != 0) {
					// axios.post('http://47.75.223.44:8080/api/applyIOST', this.getApplyParam()).then((response) => {
					axios.post(`${apis.applyIOST}`, this.getApplyParam()).then((response) => {
						let retCode = response.data.code
						let txHash = response.data.data
						if (retCode != 0) {
							// axios.post('http://47.75.223.44:8080/api/applyIOST', this.getApplyParam()).then((response) => {
							axios.post(`${apis.applyIOST}`, this.getApplyParam()).then((response) => {
								let retCode = response.data.code
								let txHash = response.data.data
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
									pushParams.mobile = this.mobile
									pushParams.privKey = this.privKey
									pushParams.txHash = txHash
									console.log(pushParams)
									this.$router.push({
										name: 'ApplyIOSTSuccess',
										params: pushParams
									})
								}
							})
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
							pushParams.mobile = this.mobile
							pushParams.privKey = this.privKey
							pushParams.txHash = txHash
							this.$router.push({
								name: 'ApplyIOSTSuccess',
								params: pushParams
							})
						}
					})
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
					pushParams.mobile = this.mobile
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
				const pubKey = secp256k1.publicKeyCreate(bytes)

				this.privKey = base58.encode(bytes)
				// this.addressx = base58.encode(pubKey)
				this.address = base58.encode(pubKey)
				this.auto = 1
			}
		}
	},
	mounted: function() {
		$("#phone").intlTelInput()
		
		window.setTimeout(function() {
			if ($('.g-recaptcha').html().length == 0) {
				// swal({
				//   title: 'Failed to load google reCaptcha !',
				//   text: '',
				//   confirmButtonText: 'try again'
				// }).then((result) => {
				// 	location.reload()
				// })
        grecaptcha.render('recap',{"sitekey": "6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"})
			}
		}, 3000)
	},
	// created: function() {
	// 	if (window.location.href.substr(-2) !== '?r') {
   //  		window.location = window.location.href + '?r';
   //  		// console.log(window.location)
   //  		// location.reload()
	// 	}
	// }
}
</script>


<style lang="less" rel="stylesheet/less">
	.applyIost-box {
		padding-top: 90px;
		padding-bottom: 160px;
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
				margin-left: 100px;
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
					margin-left: 100px;
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
					width: 700px;
					background-color: #F6F7F8;
					outline: none;
					border: 0;
					padding: 6px 12px;
					color: #555;
				}
				p {
					margin-bottom: 0;
					width: 100px;
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
				margin-left: 100px;
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
</style>