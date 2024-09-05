<script>
export default {
	data: function() {
		return {
			userId: sessionStorage.getItem('userId'),
			username: sessionStorage.getItem('username'),
			logged: sessionStorage.getItem('logged'),
            loading: false,
			positiveBanner: null,
            errormsg: null
		}
	},
	methods: {
        async setMyUserName(){
            this.loading = true;
            this.errormsg = null;
			const newUsername = document.getElementById("username-input").value;
			if (newUsername === ""){
				this.errormsg = 'Username cannot be empty, please type a valid username'
				return
			}
			document.getElementById("username-input").value = ""
            try {
                let response = await this.$axios.put("/users/"+sessionStorage.getItem('userId')+"?username="+ newUsername, {}, {headers:{'Authorization': 'Bearer ' + sessionStorage.getItem('userId')}});
				if (response.status<400){
				sessionStorage.setItem('username', newUsername )
                console.log("Username succesfully changed", sessionStorage.getItem('username'));
				this.positiveBanner = 'Username succesfully changed into "' + newUsername + '"'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
				}
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
	logOut(){
		sessionStorage.clear()
		window.location.reload()
	}
        }
    }

</script>

<template>
	<p  v-if="!this.logged" style="margin-top: 40px;">Please log-in first in the home page</p>
	<div v-if="this.logged">
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Settings</h1>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<p v-if="positiveBanner" class="positive-banner">{{ this.positiveBanner }}</p>
		<div style="border: solid; border-color: white; border-radius: 10px; padding: 20px; margin-bottom: 20px; margin-right: 20vw; ">
			<p class="h4">Change Username?</p>
			<div style="display: flex; flex-direction: row; margin-top: 10px;">
				<input type="text" id="username-input" style="margin-right: 10px;" class="form-control" placeholder="new Username">
        		<button class="btn btn-outline-secondary" type="button" id="username-button" @click="setMyUserName">Submit</button>
			</div>
		</div>
			<div style="margin-left: 10px;">
				<p class="h4">Log out</p>
        		<button class="btn btn-outline-secondary" type="button" id="username-button" @click="logOut">Log Out</button>
			</div>
	</div>

</template>

<style>
</style>

