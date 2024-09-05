<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			userId: sessionStorage.getItem('userId'),
			username: sessionStorage.getItem('username'),
			logged: sessionStorage.getItem('logged'),
			stream: [],
			comments: []
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.getMyStream()
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		
		async doLogin() {
            this.loading = true;
            this.errormsg = null;
			sessionStorage.setItem('username', document.getElementById("username-input").value) 
            try {
                console.log(sessionStorage.getItem('userId'))
                let response = await this.$axios.post("/session?username=" + sessionStorage.getItem('username'));
				if (response.status == "200") {
					sessionStorage.setItem('userId', response.data)
					sessionStorage.setItem('logged', true)
					console.log(sessionStorage.getItem('logged'))
					window.location.reload()
				}
            } catch (e) {
                this.errormsg = e.toString();
				console.log(e)
            }
            this.loading = false;
        },

		async getMyStream(refresh=true){
			if (!sessionStorage.getItem('logged')){
				return
			}
			this.loading = true
			if (this.logged){
			try{
				let response = await this.$axios.get("/users/"+ sessionStorage.getItem('userId') + "/stream?refresh=" + refresh,{headers:{
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}
				});
				this.stream = this.stream.concat(response.data)
				if (this.stream.length>0){
				for (let i = 0; i < this.stream.length; i++) {
					this.stream[i].image = 'data:image/*;base64,' + this.stream[i].image
				}}
				this.comments = response.data.comments
				this.loading = false

			} catch(e) {
				this.errormsg = e.toString();
				console.log(e)
			}
		}
		},

		async likePhoto(photoId, i){
			try {
                let response = await this.$axios.post("/likes/?photoId="+ photoId +"&userId=" + sessionStorage.getItem('userId'),{},{headers:{
					"Authorization": "Bearer " + sessionStorage.getItem('userId')
				}});
				if (response.status<400) {
					console.log("photo liked")
					const button = document.getElementById('Lo'+photoId)
					button.id = response.data
					button.className= 'like-button-active'
					this.stream[i].likeId = response.data
					this.stream[i].likes++;
				}

            } catch (e) {
                this.errormsg = e.toString();
				console.log(e)
            }
		},

		async unlikePhoto(photoId, id, i){
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
					this.stream[i].likes--;
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
				this.likePhoto(photoId, i)
			} else {
				this.unlikePhoto(photoId, id, i)
			}
			},
		async commentPhoto(photoId, i) {
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
					this.stream[i].comments.push({
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
		async uncommentPhoto(commentId){
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
		},
	},


	
	mounted() {
		this.getMyStream()
	}
}
</script>

<template>
	<div>
	  <div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Home page</h1>
	  </div>
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  <div v-if="!logged" style="text-align: left; border-radius:10px; padding: 20px;">
	  <p class="h3">Register or log-in</p>
	  
		<div style="border: solid; border-color: white; border-radius: 10px; padding: 20px; margin-right: 50vw; ">
			<div style="display: flex; flex-direction: row;">
				<input type="text" id="username-input" style="margin-right: 10px;" class="form-control" placeholder="Username" v-model="this.username">
        		<button class="btn btn-outline-secondary" type="button" @click="doLogin">Login</button>
			</div>	
		</div>
	</div>
	  <div v-if="logged && !stream.length" style="text-align: center;"> 
		
		<p class="h3">Here you will see all the posts of your friends</p>
	
	  </div>

	  <div v-for="(photo, i) in stream" class="post-container" v-bind:key="'for1'+photo.photoId">
		<div class="image-container">
			<p class="username-title"> {{ photo.username }}</p>
			<div style="object-fit: contain;">
		 		 <img :src="photo.image" class="profile-image" :id="photo.photoId"> 
			</div>
		</div>
		<div class="buttons-container">
			<button :class="photo.likeId === '' ? 'like-button' : 'like-button-active'" :id="photo.likeId !== '' ? photo.likeId : 'Lo'+photo.photoId " @click="likeEventHandler(photo.photoId, $event, i)">{{photo.likes}} <svg class="feather" @click.stop><use href="/feather-sprite-v4.29.0.svg#heart"/></svg></button>
			<div>
		    	<input type="text" placeholder="Leave a Comment..." class="comment-input" :id="'CI'+photo.photoId">
				<button class="comment-button" :id="'C'+photo.photoId" @click="commentPhoto(photo.photoId, i)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg> Comment</button>
			</div>
		</div>
		<div v-for="comment in photo.comments" class="comment-box" v-bind:key="'for2' + comment.commentId">
			<div style="flex-direction: row; display: flex; justify-content: space-between;" :id="'D'+comment.commentId">
				<p v-if="comment.username == this.username" class="comment-title">You</p>
				<p v-if="comment.username != this.username" class="comment-title">{{ comment.username }}</p>
				<button v-if="comment.username == this.username" :id="comment.commentId" class="delete-comment-button" @click="uncommentPhoto(comment.commentId)"><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg></button>
			</div>
			<p class="comment-content" :id="'C'+comment.commentId">{{ comment.content }}</p>
		</div>
	  </div>
	  <button v-if="this.stream.length>14" class="btn btn-outline-secondary" @click="getMyStream(false)" style="margin-bottom: 30px;">Load More</button>
	</div>
  </template>
  
  <style>
  .post-container {
	display: flex;
	flex-direction: column;
	align-items: left;
	width: 50vw;
	margin-bottom: 30px;
	padding-top: 5vh;
	border-top-left-radius: 10px;
	border-bottom-left-radius: 10px;
	border: solid;
	border-color: gray;
	overflow: auto;
	max-height: 100vh;
  }
  
  .image-container {
	width: 100%;
	text-align: center;
	object-fit: contain;
  }
  
  .like-button {
	background-color: white;
	color: red;
	border: solid;
	border-color: red;
	border-radius: 10px;
	padding: 12px 20px;
	cursor: pointer;
	transition: background-color 0.3s, color 0.3s;
	font-family: 'Arial';
	font-weight: bold;
  }

  .like-button-active {
	color: white;
	background-color: red;
	border: solid;
	border-color: red;
	border-radius: 10px;
	padding: 12px 20px;
	cursor: pointer;
	transition: background-color 0.3s, color 0.3s;
	font-family: 'Arial';
	font-weight: bold;
  }
  
  .like-button:hover {
	background-color: red;
	color: white;
  }
  
  .comment-button {
	color: #007bff;
	border: solid;
	border-color: #007bff;
	padding: 12px 20px;
	background-color: #fff;
	border-radius: 10px;
	cursor: pointer;
	transition: background-color 0.3s, color 0.3s;
	font-family: 'Arial';
	font-weight: bold;
	margin-left: 10px;
  }

  .comment-input{

	color: grey;
	border: solid;
	border-color: grey;
	padding: 12px 20px;
	background-color: #fff;
	border-radius: 10px;
	cursor: text;
	transition: background-color 0.3s, color 0.3s;
	font-family: 'Arial';
	font-weight: bold;
	margin-left: 10px;
  }
  
  .comment-button:hover {
	color: #fff;
	background-color: #007bff;
  }
  
  .comment-button::placeholder {
	color: #007bff;
  }
  
  .comment-button:hover::placeholder {
	color: #fff;
  }
  
  .buttons-container {
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	margin-top: 20px;
	margin-bottom: 15px;
	margin-left: 2.5vw;
	margin-right: 2.5vw;
  }
  .username-title{
	font-family: 'Arial'; 
	font-weight: bold; 
	font-size: 20px; 
	text-align: left; 
	margin-left: 3vw;
  }

.comment-title{
	font-family: 'Arial'; 
	font-weight: bold; 
	font-size: 15px; 
	text-align: left; 
	margin-left: 3vw;
	margin-bottom: 2px;
}
.comment-content{
	font-family: 'Arial'; 
	font-size: 13px; 
	text-align: left; 
	margin-left: 3vw;
	margin-right: 5vw;
	border-radius: 10px;
	word-wrap: break-word;
}
.comment-box{
	border-radius: 10px;
	width: 99%
}
.delete-comment-button{
	margin-right: 2vw;
	padding-left: 4px; 
	padding-right: 4px; 
	padding-top: 0px;
	padding-bottom: 0px;
	font-family: 'Arial'; 
	font-weight: bold; 
	margin-left: 10px;
	border-radius: 5px;
	background-color: white;
	border: solid;
	border-color: grey;
	color: grey;
	font-size: 15px;

}
.delete-comment-button:hover{
	background-color: grey;
	color: white;
}
.positive-banner{
	background-color: rgba(26, 255, 0, 0.5);
	border-radius: 10px;
	padding: 20px;
	display: inline-block;
}
  </style>
  
