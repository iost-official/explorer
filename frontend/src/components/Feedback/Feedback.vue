<template>
  <div>
  	<div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Feedback</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Feedback</li>
          </ol>
        </div>
      </div>
    </div>
  <div class="container" id="feedback">
  	<div class="row">
  		<div class="col-md-1"></div>
  		<div class="col-md-10">
  			<div class="" role="alert" id="errFeedbackAlert" style="text-align: left;">{{errMsg}}</div>
  			<form class="form-horizontal">
			  <div class="form-group">
			    <label for="inputEmail3" class="col-sm-2 control-label">Email</label>
			    <div class="col-sm-10">
			      <input type="email" class="form-control" placeholder="Email" v-model.trim="email">
			    </div>
			  </div>
			  <div class="form-group">
			    <label for="inputPassword3" class="col-sm-2 control-label" v-model.trim="email">Feedback</label>
			    <div class="col-sm-10">
			      <textarea class="form-control" rows="10" v-model.trim="content"></textarea>
			    </div>
			  </div>
			  <div class="form-group">
			    <div class="col-sm-offset-2 col-sm-10" style="text-align: left;">
			      <button type="button" class="btn btn-primary" @click="sendFeedback">Send</button>
			    </div>
			  </div>
			</form>
  		</div>
  		<div class="col-md-1"></div>
  	</div>
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

			axios.post('https://explorer.iost.io/api/feedback', params).then((response) => {
				swal(
				  'Feedback success!',
				  '',
				  'success'
				).then(function(router) {
					window.location.href="/#/"
				})
			})
		}
	}
}
</script>