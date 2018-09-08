<template>
  <div>
  	<div class="explorer-header">
      <div class="container">
        <div class="row">
          <h1 class="pull-left" style="display: block">Request</h1>
          <ol class="breadcrumb pull-right">
            <li><a href="/#/">Home</a></li>
            <li class="active">Request</li>
          </ol>
        </div>
      </div>
    </div>

    <div class="container">
      <div class="row" id="banner-luckybet">
        <i class="far fa-calendar-alt"></i> Latest Activity: <a href="/luckyBet" target="_blank">Play Lucky Bet !</a>
      </div>
    </div>

    <div class="container" style="margin-top: 30px; text-align: left;">
      <div class="row">
      	<table class="table apply-iost">
      	  <thead>
      	  	<tr>
		  	  <th colspan="2">Request successfully processed</th>
      	  	</tr>
      	  </thead>
		  <tbody>
		  	<tr>
		  		<td>TxHash:</td>
		  		<td><a :href="'/#/tx/' + txHash" target="_blank">{{txHash}}</a></td>
		  	</tr>
		  	<tr>
		  	  <td>Address:</td>
		  	  <td>{{address}}</td>
		  	</tr>
		  	<tr>
		  	  <td>Amount:</td>
		  	  <td>10.1 IOST</td>
		  	</tr>
		  	<tr>
		  	  <td>Email:</td>
		  	  <td>{{email}}</td>
		  	</tr>
		  	<tr>
		  	  <td>Mobile:</td>
		  	  <td>{{mobile}}</td>
		  	</tr>
		  </tbody>
		</table>
      </div>

      <div class="row applyAddressPrivateKey" v-if="privKey != ''">
      	<p>Save your Account Info</p>
      	<div><p class="bg-success center-block"><span>Address:</span> {{address}}</p></div>
      	<div><p class="bg-primary center-block margin-top:-10px;"><span>Private Key:</span> {{privKey}}</p></div>
      	<p>**Do not lose it!** It cannot be recovered if you lose it.</p>
      	<p>**Do not share it!** Your funds will be stolen if you use this private key on a malicious/phishing site.</p>
      	<p></p>
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
			email: this.$route.params.email,
			mobile: this.$route.params.mobile.replace(' ', '+'),
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