<script>


export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
			token:null,
		}
	},
	methods: {
		async login() {
			this.loading = true;
			this.errormsg = null;

			const queryParams = {
				username: this.username	
			};
			
			try {
				
				let response = await this.$axios.post("/session",null,{params: queryParams});
				if (response.status == "200") {
					this.token = response.data.message;
					localStorage.setItem('token', this.token);
					sessionStorage.setItem(this.token, this.username);
					sessionStorage.setItem('logged', true);
					this.$router.push('users/profile');
				}
				
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.loading = false;
			this.isLoggedIn = true;
		},
		
	},
	mounted(){
		sessionStorage.clear();
	},
	
	
	
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login</h1>
			
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
	<div> 
		<form @submit.prevent="login"> 
			<input type="text" id="username" v-model="username" required> <br>
		</form>
		<button @click="login">Confirm</button>

	</div>
</template>

<style>
</style>