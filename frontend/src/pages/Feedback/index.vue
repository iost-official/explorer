<template>
  <div class="feedBack-box">
		<div class="luckyBet-box">
			<img src="../../assets/activity.png" alt="">Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
		</div>
  	<div class="feedBack-header">
      <div class="my-header-container">
				<h1>Feedback</h1>
      </div>
    </div>
  	<div class="my-container">
			<div class="err-msg" role="alert" id="errFeedbackAlert" style="text-align: left;">{{errMsg}}</div>
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
	}
}
</script>

<style lang="less" rel="stylesheet/less">
	.feedBack-box {
		padding-top: 90px;
		margin: 0 auto;
		padding-bottom: 160px;
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
					box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
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
				border: none;
				color: #FFFFFF;
				font-weight: bold;
				&.active {
					background-color: rgba(44,46,49,1);
				}
			}
			.a {
				display: block;
				width: 100%;
				height: 34px;
				padding: 6px 12px;
				font-size: 14px;
				line-height: 1.42857143;
				color: #555;
				background-color: #fff;
				background-image: none;
				border: 1px solid #ccc;
				border-radius: 4px;
				-webkit-box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
				box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
				-webkit-transition: border-color ease-in-out .15s, -webkit-box-shadow ease-in-out .15s;
				-o-transition: border-color ease-in-out .15s, box-shadow ease-in-out .15s;
				transition: border-color ease-in-out .15s, box-shadow ease-in-out .15s;
			}
		}
	}
</style>