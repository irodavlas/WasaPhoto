<script>

export default {
	data: function() {
		return {
			username: sessionStorage.getItem('username'),
			userId: sessionStorage.getItem('userId'),
			logged: sessionStorage.getItem('logged'),
			deleteMode: false,
			profile:  [],
			errormsg: null,
			positiveBanner: null,
			profileRetrieved: null,
			following: 0,
			followers: 0,

		}
	},
	mounted() {
		this.getYourProfile(sessionStorage.getItem('username'))
	},
	methods: {
		async getYourProfile(username){
			if (!sessionStorage.getItem('logged')){
				return
			}
			this.loading = true;
			this.errormsg = null;
			this.profileRetrieved = false;
			console.log(username)
			const config = {
				headers: {
				'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			};
			try {
				let response = await this.$axios.get("/profiles/"+ username + "?userId=" + sessionStorage.getItem('userId'), config);
	
				this.profile = response.data.photos
				this.followers = response.data.followers
				this.following = response.data.following

				for (let i = 0; i < this.profile.length; i++) {
					this.profile[i].image = 'data:image/*;base64,' + this.profile[i].image
				}
	
				console.log(this.profile)

			} catch (e) {
				this.errormsg = e.toString();
			}
			if (this.profile.length){
				const deleteButton = document.getElementById("delete-button")
				deleteButton.disabled = false
			}
			this.loading = false;
		},

		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			console.log("refreshing")
		},

		async uploadPhoto() {
		console.log('userId ' + sessionStorage.getItem('userId'))
  		const fileInput = document.getElementById('photoInput');
 		const file = fileInput.files[0];
		if (!file) {
			return;
		}

		// const formData = new FormData();
		// formData.append("image", file)
		try {
			const url = "/photos/?userId=" + sessionStorage.getItem('userId');
			const config = {
				headers: {
				'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			};
			console.log("file", file)
			const response = await this.$axios.post(url, file, config);
			if (response.status == 201){
				const newPhoto = {
				photoId: response.data
			}
			this.positiveBanner = 'Photo succesfully uploaded'
				setTimeout(()=>{
					this.positiveBanner = null
				}, 3000)
			this.profile.push(newPhoto)
			if (this.profile.length){
				const deleteButton = document.getElementById("delete-button")
				deleteButton.disabled = false
			}
			window.location.reload()
			}

    } catch (error) {
      console.error("Error occurred:", error);
    }
},

	async toggleDeleteMode() {
	if (this.deleteMode == false){
		var deleteButton = document.getElementById('delete-button');
		var photos = document.getElementsByClassName("profile-image")
		var buttons = document.querySelectorAll('button');
		var selector = document.getElementById('photoInput');
		selector.disabled = true

		buttons.forEach(function(button) {
    		button.disabled = true;
		});
		deleteButton.disabled = false
		this.deleteMode = true;
		deleteButton.textContent = 'Cancel';


		for (var i = 0; i < photos.length; i++) {
			photos[i].style.borderColor = "#EE5535";
			photos[i].style.borderWidth = '3px';
			photos[i].style.borderStyle = 'solid';
			// photos[i].style.backgroundColor = "#EE5535";
			photos[i].disabled = false;
	}}
	else{
		var deleteButton = document.getElementById('delete-button');
		this.deleteMode = false;
		deleteButton.textContent = 'Delete Photo';
		var photos = document.getElementsByClassName("profile-image")
		for (var i = 0; i < photos.length; i++) {
			// photos[i].style.backgroundColor = "transparent";
			photos[i].style.border = "none";
			photos[i].disabled = true;
			}
		var buttons = document.querySelectorAll('button');
		var selector = document.getElementById('photoInput');
		selector.disabled = false

		buttons.forEach(function(button) {
    	button.disabled = false;
		if (photos.length == 0){
			deleteButton.disabled = true
		}
		});
}      
},	
async deletePhoto(photoId){
		this.loading = true
		if (this.deleteMode){
		var photo = document.getElementById('ForProfile' + photoId)
		this.toggleDeleteMode()
		try{
			const response = this.$axios.delete("/photos/" + photoId + "?userId=" + sessionStorage.getItem('userId'), {
				headers:{
					'Authorization': 'Bearer ' + sessionStorage.getItem('userId')
				}
			})
			
			for (var i = 0; i < this.profile.length; i++) {
				if (this.profile[i].photoId == photoId){
					this.profile.splice(i,1);
					break
				}
			}
			console.log(this.profile)
			window.location.reload()
			const deleteButton = document.getElementById("delete-button")
		if (this.profile.length){
			deleteButton.disabled = false
		} else{
			deleteButton.disabled = true
		}
		this.loading = false
		} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		}
		
	},
async likePhotoProfile(photoId, i){
		try {
			let response = await this.$axios.post("/likes/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId'),{},{headers:{
				"Authorization": "Bearer " + sessionStorage.getItem('userId')
			}});
			if (response.status<400) {
				console.log("photo liked")
				const button = document.getElementById('Lo'+photoId)
				button.id = response.data
				button.className= 'like-button-active'
				this.profile[i].likeId = response.data
				this.profile[i].likes++;
			}

		} catch (e) {
			this.errormsg = e.toString();
			console.log(e)
		}
	},

	async unlikePhotoProfile(photoId, id, i){
		try {
			let response = await this.$axios.delete("/likes/" + id + "?userId=" + sessionStorage.getItem('userId'), {
				headers: {
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}
			});

			if (response.status<400) {
				console.log("Photo disliked");
				const button = document.getElementById(id)
				button.id = 'Lo' + photoId
				button.className = 'like-button'
				this.profile[i].likes--;
			}

		} catch (e) {
			this.errormsg = e.toString();
			console.error(e);
		}
	},

	async likeEventHandler(photoId, event, i) {
		this.loading = true;
		this.errormsg = null;
		const id = event.target.id
		console.log(event.target.id);

		if (id == 'Lo'+photoId) {
			this.likePhotoProfile(photoId, i)
		} else {
			this.unlikePhotoProfile(photoId, id, i)
		}
		},
	async commentPhotoProfile(photoId, i) {
				this.loading = true;
				this.errormsg = null;
				const commentInput = document.getElementById('CI'+ photoId)
				const comment = commentInput.value
				if (!comment){
					return
				}
				const formData = new FormData();
				formData.append("content", comment)
				console.log(comment)
				commentInput.value = ""
				const config = {headers:{
						"Authorization": "Bearer " + sessionStorage.getItem('userId')
					}}
				const url = "/comments/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId')
				try {
					let response = await this.$axios.post(url, formData, config);
					if (response.status == 201) {
						console.log(response.data)
					}
				this.profile[i].comments.push({
					content: comment,
					commentId: response.data,
					username: this.username
				})
				} catch (e) {
					this.errormsg = e.toString();
					console.log(e)
				}
				this.loading = false;
	},
	async uncommentPhotoProfile(commentId){
	try {
			let response = await this.$axios.delete("/comments/" + commentId + "?userId=" + sessionStorage.getItem('userId'), {
				headers: {
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}
			});

			if (true) {
				console.log("comment deleted");
				const content = document.getElementById('C'+commentId)
				const div = document.getElementById('D'+commentId)
				div.remove()
				content.remove()
			}

		} catch (e) {
			console.log("/likes/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId'))
			this.errormsg = e.toString();
			console.error(e);
		}
	}
}}

</script>

<template>
	<p v-if="!this.logged"  style="margin-top: 40px;">Please log-in first in the home page</p>
	<div v-if="this.logged">
		<div>
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">Hi {{ this.username }}</h1>
			</div>
			<p v-if="positiveBanner" class="positive-banner">{{ this.positiveBanner }}</p>
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div>
				<h1 class="h4">Followers: {{ this.followers }}</h1>
				<h1 class="h4" style="margin-bottom: 20px;">Following: {{ this.following }}</h1>
				<div style="display: flex; flex-direction: row; margin-bottom: 20px;">
					<input type="file" class="btn btn-outline-secondary" id="photoInput" ref="photoInput" accept="image/*">
					<button class="btn btn-outline-secondary" id="post-button" type="button" style="margin-left: 10px" @click="uploadPhoto">Post Photo</button>
				</div>
				<p v-if="deleteMode" class="h3" style="margin-top: 20px;">Click on the image you want to delete</p>
			</div>
			<div>
				<div v-for="(photo, i) in profile" class="post-container" v-bind:key="'forProfile'+photo.photoId" @click="deletePhoto(photo.photoId)">
					<div class="image-container">
						<p class="username-title"> {{ photo.username }}</p>
						<img :src="photo.image" :id="photo.photoId" class="profile-image"> 
					</div>
					<div class="buttons-container">
						<button :class="photo.likeId === '' ? 'like-button' : 'like-button-active'" :id="photo.likeId !== '' ? photo.likeId : 'Lo'+photo.photoId " @click="likeEventHandler(photo.photoId, $event, i)">{{photo.likes}} <svg class="feather" @click.stop><use href="/feather-sprite-v4.29.0.svg#heart"/></svg></button>
						<div>
							<input type="text" placeholder="Leave a Comment..." class="comment-input" :id="'CI'+photo.photoId">
							<button class="comment-button" :id="'C'+photo.photoId" @click="commentPhotoProfile(photo.photoId, i)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg> Comment</button>
						</div>
					</div>
					<div v-for="comment in photo.comments" class="comment-box" v-bind:key="'forProfile2' + comment.commentId">
						<div style="flex-direction: row; display: flex; justify-content: space-between;" :id="'D'+comment.commentId">
							<p v-if="comment.username == this.username" class="comment-title">You</p>
							<p v-if="comment.username != this.username" class="comment-title">{{ comment.username }}</p>
							<button v-if="comment.username == this.username" :id="comment.commentId" class="delete-comment-button" @click="uncommentPhotoProfile(comment.commentId)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg></button>
			</div>
			<p class="comment-content" :id="'C'+comment.commentId">{{ comment.content }}</p>
		</div>
	  </div>
			</div>
			<button class="btn btn-outline-secondary" id="delete-button" type="button" @click="toggleDeleteMode" disabled>Delete Photo</button>	
		</div>
	</div>
</template>

<style>

.profile-image {
    background-color: transparent;
	border-radius: 10px;
	width: 90%;
	height: auto;
	
}	

</style>


