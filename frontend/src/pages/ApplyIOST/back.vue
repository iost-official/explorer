<template>
  <div>
  	<div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Apply</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Apply</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container" style="margin-top: 50px; text-align: left;">
      <div class="row">
      	<div class="col-md-2"></div>
      	<div class="col-md-8">
      	  <div class="" role="alert" id="errAlert">{{errMsg}}</div>
      	  <form class="form-horizontal">
			<div class="form-group">
			  <label class="col-sm-3 control-label">Address</label>
			  <div class="col-sm-9">
			    <input class="form-control" placeholder="Address" id="applyAddress" v-model.trim()="address">
			  </div>
			</div>
			<div class="form-group" style="margin-top: -15px;">
			  <label class="col-sm-3 control-label"></label>
			  <div class="col-sm-9">
			    <div class="checkbox">
				  <label>
				    <input type="checkbox" id="generateAddressCheck" v-model="checked">
				    generate address for me
				  </label>
				</div>
			  </div>
			</div>
			<div class="form-group">
			  <label class="col-sm-3 control-label">Amount</label>
			  <div class="col-sm-9">
			    <span style="display: block; padding-top: 7px">10 IOST</span>
			  </div>
			</div>
			<div class="form-group">
			  <label class="col-sm-3 control-label">Email</label>
			  <div class="col-sm-9">
			    <input type="email" class="form-control" placeholder="email" v-model.trim()="email">
			  </div>
			</div>
			<div class="form-group">
			  <label class="col-sm-3 control-label">Mobile</label>
			  <div class="col-sm-9">
			    <input type="mobile" class="form-control" placeholder="mobile" value="+86" v-model.trim()="mobile">
			  </div>
			</div>
			<div class="form-group">
			  <div class="col-sm-offset-3 col-sm-9">
			    <div class="g-recaptcha" data-sitekey="6Lc1vF8UAAAAAMo-EsF4vRt6CWxM8s56lAeyHnBe"></div>
			  </div>
			</div>
			<div class="form-group">
			  <label class="col-sm-3 control-label">Verification Code</label>
			  <div class="col-sm-7">
			    <input type="mobile" class="form-control" placeholder="Verification code" v-model.trim()="verify">
			  </div>
			  <div class="col-sm-2">
			  	<button type="button" id="sendSmS" class="btn btn-default" style="width:100%" @click="sendSmS">{{mobileButtonMsg}}</button>
			  </div>
			</div>
			<div class="form-group">
			  <div class="col-sm-offset-3 col-sm-9">
			    <button type="button" class="btn btn-primary" @click="apply">Apply</button>
			  </div>
			</div>
	      </form>
      	</div>
      	<div class="col-md-2"></div>
      </div>
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
// import swal from 'sweetalert2'

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
			mobile: '+86',
			verify: '',
			mobileButtonMsg: 'send',
			ltime: 0,
			errMsg: '',
			checked: false,
			dott: '',
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

			var params = new URLSearchParams();
			params.append('mobile', this.mobile)
			params.append('gcaptcha', grecap)

			axios.post('https://explorer.iost.io/api/sendSMS', params).then((response) => {
				var retCode = response.data.ret
				if (retCode != 0) {
					$('#errAlert').addClass('alert alert-danger')
					this.errMsg = response.data.msg
				}
			})

			this.ltime = 3
			this.mobileButtonMsg = 'retry in ' + this.ltime + 's...'
			$("#sendSmS").attr('disabled', true)
			let interval = window.setInterval(() => {
				if (--this.ltime <=0) {
					this.mobileButtonMsg = 'send'
					window.clearInterval(interval)
					$("#sendSmS").attr('disabled', false)
				} else {
					this.mobileButtonMsg = 'retry in ' + this.ltime + 's...'
				}
			}, 1000)
		},
		getApplyParam(grecap) {
			var params = new URLSearchParams();
			if (this.auto == 1) {
				params.append('address', this.addressx)
			} else {
				params.append('address', this.address)
			}
			params.append('email', this.email)
			params.append('mobile', this.mobile)
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

			if (this.mobile.length != 14 || this.mobile[0] != '+') {
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

			$('#applyIOSTModal').modal({
				backdrop: 'static',
				keyboard: false
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

			axios.post('https://explorer.iost.io/api/applyIOST', this.getApplyParam(grecap)).then((response) => {
				let retCode = response.data.ret
				let txHash = response.data.msg
				if (retCode != 0) {
					axios.post('https://explorer.iost.io/api/applyIOST', this.getApplyParam()).then((response) => {
						let retCode = response.data.ret
						let txHash = response.data.msg
						if (retCode != 0) {
							axios.post('https://explorer.iost.io/api/applyIOST', this.getApplyParam()).then((response) => {
								let retCode = response.data.ret
								let txHash = response.data.msg
								if (retCode != 0) {
									$('#applyIOSTModal').modal('hide')
									$('#errAlert').addClass('alert alert-danger')
									this.errMsg = response.data.msg
								} else {
									$('#applyIOSTModal').modal('hide')
									let pushParams = {}
									if (this.address.length > 0) {
										pushParams.address = this.address
									} else {
										pushParams.address = this.addressx
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
							$('#applyIOSTModal').modal('hide')
							let pushParams = {}
							if (this.address.length > 0) {
								pushParams.address = this.address
							} else {
								pushParams.address = this.addressx
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
					$('#applyIOSTModal').modal('hide')
					let pushParams = {}
					if (this.address.length > 0) {
						pushParams.address = this.address
					} else {
						pushParams.address = this.addressx
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
				this.addressx = base58.encode(pubKey)
				this.auto = 1
			}
		}
	},
	mounted: function() {
		// window.setTimeout(function() {
		// 	console.log('aa:' + $('.g-recaptcha').html())
		// 	if ($('.g-recaptcha').html().length == 0) {
		// 		swal({
		// 		  title: 'Failed to load google reCaptcha !',
		// 		  text: '',
		// 		  type: 'error',
		// 		  confirmButtonText: 'try again'
		// 		}).then((result) => {
		// 			location.reload()
		// 		})

		// 	}
		// }, 2000)
	},
	created: function() {
		if (window.location.href.substr(-2) !== '?r') {
    		window.location = window.location.href + '?r';
    		location.reload()
		}
	}
}
</script>