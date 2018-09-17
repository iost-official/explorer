<template>
  <div class="feedBack-box">
		<LuckyBet/>

  	<div class="feedBack-header">
      <div class="my-header-container">
				<h1>Feedback</h1>
      </div>
    </div>
  	<div class="my-container">
			<div class="err-msg" role="alert" id="errFeedbackAlert" style="text-align: center;">{{errMsg}}</div>
			<div class="my-group">
				<p>Email：</p>
				<input type="email" class="my-input" placeholder="Email" v-model.trim="email">
			</div>
			<div class="my-group">
				<p>Feedback：</p>
				<textarea class="my-input" v-model.trim="content" style="resize:none"></textarea>
			</div>
			<button type="button" @click="sendFeedback" :class="{active: isOK}">Send</button>
  	</div>
  </div>
</template>

<script>
import swal from 'sweetalert2'
import axios from 'axios';
import LuckyBet from '../../components/LuckyBet'

export default {
	name: "Feedback",
	data() {
		return {
			email: '',
			content: '',
			errMsg: '',
		}
	},

	computed: {
	  isOK () {
	    if (this.email != '' && this.content != '') {
	    	return true
	    } else {
	      return false
			}
		}
	},
	methods: {
		sendFeedback: function() {
			var regExp = /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/
			if (this.email.length == 0 || !regExp.test(this.email)) {
				$('#errFeedbackAlert').addClass('alert alert-danger')
				this.errMsg = 'invalid email address'
				return false
			}

			if (this.content.length == 0 || this.content.length < 100) {
				$('#errFeedbackAlert').addClass('alert alert-danger')
				this.errMsg = 'feedback content length must be greater than 100'
				return false
			}

			var params = new URLSearchParams();
			params.append('email', this.email)
			params.append('content', this.content)

			// axios.post('https://explorer.iost.io/api/feedback', params).then((response) => {
			axios.post('http://47.75.223.44:8080/api/feedback', params).then((response) => {
				swal(
				  'Feedback success!',
				  '',
				  'success'
				).then(function(router) {
					window.location.href="/"
				})
			})
		}
	},

  components: {
    LuckyBet
  }
}
</script>

<style lang="less" rel="stylesheet/less">
	.feedBack-box {
		padding-bottom: 160px;
		background: #F6F7F8;
		.feedBack-header {
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
			.my-group {
				display: flex;
				justify-content: space-between;
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
				textarea {
					height: 420px;
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

	@media screen and (max-width:480px) {
		.feedBack-box {
			.feedBack-header {
				.my-header-container {
					height: auto;
					width: 100%;
					flex-direction: column;
					padding: 0 25px;
					text-align: left;
				}
			}

			.my-container {
				width: 100%;
				padding: 24px 25px;
				.err-msg {
					margin-bottom: 24px;
					margin-left: 76px;
				}
				> div {
					&:nth-child(2) {
						margin-bottom: 24px;
					}
					&:nth-child(3) {
						margin-bottom: 24px;
					}
				}
				.my-group {
					.my-input {
						width: 300px;
					}
					input {
						height: 54px;
					}
					textarea {
						height: 320px;
					}
				}

				> button {
					width: 288px;
					margin-left: 76px;
				}
			}
		}
	}

	@media screen and (max-width:375px) {
		.feedBack-box {
			.my-container {
				.err-msg {
					margin-left: 64px;
				}
				> button {
					width: 260px;
					margin-left: 64px;
				}
			}
		}
	}
</style>