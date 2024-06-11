<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: null,
			profile : null,
			follower: null,
			following: null,
			banned:[],
			posts: [],
			userToBan: null,
			newUsername: null,
			selectedFile: null,
			loading: false,
			isLoggedIn:false,
			

			CommentBoxShowing:false,
			selectedPost:null,
			commentContent:null,
		}
	},
	methods: {
		async unbanUser(username){
			this.loading = true;
			this.errormsg = false;
			

			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			const url = '/ban/' + username;
			
			try {
				let response = await this.$axios.delete(url, {headers : headers});
				
				if (response.data.code == "200"){
					const index = this.banned.indexOf(username);
					if (index !== -1) { // -1 makes check makes sure the username exists in the list before attempting to remove it 
						
						this.banned.splice(index, 1);
					}
				}
				
			}catch (e) {
				this.errormsg = e.toString();
				
			}
			
			this.loading = false;
			this.selectedFile = null;
		},
		async banUser(){
			this.loading = true;
			this.errormsg = null;


            const token = localStorage.getItem('token');
            
            
            
            const headers = {
                'Authorization': 'Bearer ' + token
            };
           
            var url = '/ban/' + this.userToBan;
			try {
				let response = await this.$axios.put(url, null, {headers: headers});
				if (response.data.code == "200"){
					if (!this.banned.includes(this.userToBan)) {
						this.banned.push(this.userToBan);
					}
				}
                console.log(this.banned);
                
			} catch (e) {
				this.errormsg = e;
			}
			this.searchUser()
			
			this.userToBan = "";
			this.notfound = false;
			this.loading = false;
		},
		handleFileUpload(){
			 this.selectedFile = event.target.files[0];
		},
		async uploadPost(){
			this.loading = true;
			this.errormsg = false;
			let formData = new FormData();
			
			if (this.selectedFile){
				
				formData.append('image', this.selectedFile);
			} 

			const token = localStorage.getItem('token');
			const headers = {
				'Content-Type': 'multipart/form-data',
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/';
			try {
				let response = await this.$axios.post(url, formData, {headers : headers});
				const post = response.data.message;
				console.log(response);

				if (response.data.code == "200"){
					
					if (!this.posts){
						this.posts = [];
					}
					this.posts.push({'PostId': post.PostId, 'OwnerId':post.OwnerId, 'OwnerUsername':post.OwnerUsername, 'Photo': post.Photo, 'PostId':post.PostId, 'Time':post.Time});
					this.posts.sort((a, b) => new Date(b.Time) - new Date(a.Time));
				}
				
			}catch (e) {
				if (e.response.data.code == "400"){
					this.errormsg = e.response.data.message;
				}
				
				
			}
			
			this.loading = false;
			
		},
		async deletePost(postId){
			this.loading = true;
			this.errormsg = false;
			

			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + postId;
			
			try {
				let response = await this.$axios.delete(url, {headers : headers});
				
				for (let i = 0; i < this.posts.length; i++) {
					if (this.posts[i].PostId == postId){
						this.posts.splice(i, 1);
					}
					
				}
				
			}catch (e) {
				this.errormsg = e.toString();
				
			}
			
			this.loading = false;
			this.selectedFile = null;
		},
		formatDate(time){
			const date = new Date(time);
			const day = String(date.getDate()).padStart(2, '0');
			const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
			const year = date.getFullYear();
			const hours = String(date.getHours()).padStart(2, '0');
            const minutes = String(date.getMinutes()).padStart(2, '0');

			return `${day}/${month}/${year} ${hours}:${minutes}`;
		},
		async commentPost(){
			
			this.loading = true;
			this.errormsg = false;
			
			const commentMessage = this.commentContent;
			
			const token = localStorage.getItem('token');
			const headers = {
				'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + this.selectedPost.PostId + '/comment';
			try {
				let response = await this.$axios.post(url, commentMessage, {headers:headers});
				console.log(response.data.message);
				for (let i = 0; i < this.posts?.length; i++) {
					if (!this.posts[i].Comments) {
						this.posts[i].Comments = [];
					}
					if (this.posts[i].PostId == this.selectedPost.PostId){
						this.posts[i].Comments.push({'CommentId': response.data.message.CommentId,'Username': this.profile.Username, 'Message':this.commentContent});
					}
				}
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			
			this.loading = false;
			this.commentContent = null;
			
			
		},
		async deleteComment(_commentId){
			this.loading = true;
			this.errormsg = false;
			
			
			
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + this.selectedPost.PostId + '/comments/' + _commentId;
			try {
				let response = await this.$axios.delete(url, {headers:headers});
				console.log(response.data.message);

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			//update comments on post so no need to refresh information will be the same on the database side 
			for (let i = 0; i < this.posts?.length; i++) {
				if (this.posts[i].PostId == this.selectedPost.PostId){
					if (this.posts[i].Comments && this.posts[i].Comments.length > 0) {
						for (let j = 0; j < this.posts[i].Comments.length; j++) {
							
							if (this.posts[i].Comments[j].CommentId === _commentId) {
								this.posts[i].Comments.splice(j, 1); 
								break; 
							}
						}
					}
				}
			}

			this.loading = false;
			this.follow = '';
			this.commentContent = '';
			
			
		
		},
		toggleCommentBox(postId){
			
			for (let i = 0; i < this.posts?.length; i++) {
				if(this.posts[i].PostId == postId){
					this.selectedPost = this.posts[i];
				}
			}
			
			this.CommentBoxShowing = true;
		},
		async ChangeUsername(){
			this.loading = true;
			this.errormsg = null;

			const url = '/settings/username';
			const token = localStorage.getItem('token');
            const headers = {
                'Authorization': 'Bearer ' + token
            };
			const queryParams = {
                username: this.newUsername
            };
			try {
				let response = await this.$axios.put(url,  null, {params:queryParams, headers: headers});
				
				
				
				if (response.data.code == "200"){
					sessionStorage.clear();
					sessionStorage.setItem(this.token, response.data.message);
					this.searchUser();
				}
				
				
                
			} catch (e) {
				if (e.response.data.code == "400"){
					this.errormsg = e.response.data.message;
					console.log(this.errormsg);
				}
				
			}
			
			
			
			this.notfound = false;
			this.loading = false;
			this.newUsername = '';
			
			
			
		},

		async unlikePhoto(postId){
			this.loading = true;
			this.errormsg = false;

			const token = localStorage.getItem('token');
			const config = {
				headers: {
					'Authorization': 'Bearer ' + token
				}
			};
			const url = '/posts/' + postId + '/like' ;
			try {
				let response = await this.$axios.delete(url, config);
				for (let i = 0; i < this.posts?.length; i++) {
					if (this.posts[i].PostId == postId){
						if (this.posts[i].Likes && this.posts[i].Likes.length > 0) {
							
							if (this.posts[i].Likes.includes(this.profile.Username)) {
								const index = this.posts[i].Likes.indexOf(this.profile.Username);
								this.posts[i].Likes.splice(index, 1);	
							}
						}
					}
				
				}
				
				
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.loading = false;
			this.follow = '';
			
		},
		async likePhoto(postId){
			this.loading = true;
			this.errormsg = false;

			
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + postId + '/like';
			try {
				let response = await this.$axios.put(url, null, {headers:headers});
				
				for (let i = 0; i < this.posts?.length; i++) {
					if (this.posts[i].PostId == postId){
						if (!this.posts[i].Likes) {
							this.posts[i].Likes = [];
						}
						if (!this.posts[i].Likes.includes(this.profile.Username)) {
							this.posts[i].Likes.push(this.profile.Username);	
						}
					}
				
				}
				
				
			console.log(this.posts);
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.loading = false;
			this.follow = '';
			
		},
		getImagePath(photoPath) {
			
			const basePath = "images/"+ photoPath + ".jpg";

			return basePath;
		},
		async searchUser() {
			this.loading = true;
			this.errormsg = null;


            const token = localStorage.getItem('token');
           
            
            
            const headers = {
                'Authorization': 'Bearer ' + token
            };
            var url = '/users/' + token + '/profile';
			try {
				let response = await this.$axios.get(url,  {headers: headers});
				this.profile = response.data.message;
				this.banned = this.profile.Banned;
				this.follower = this.profile.Follower;
				this.following = this.profile.Following;
				this.posts = this.profile.Post;
				
                
			} catch (e) {
				this.errormsg = "User not found";
				this.profile = '';
			}
			
			
			
			this.notfound = false;
			this.loading = false;
			
		},
	},
	mounted() {
		const logged = sessionStorage.getItem('logged');
		
		if (!logged) {
			console.log("log in before accessing Profile");
			this.$router.push('/');

		}
		else{
			this.searchUser();
			
		}
		
		
	},
	
};
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profile page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
    <div> 
		

		<div v-if="this.profile" > <br>
			<p> Profile owner: {{ this.profile?.Username }} </p>
			<p> Number of posts: {{ this.posts?.length || 0 }} </p>
			<div class="scroll-container"> 
				<ul>
					<li v-for="(image, index) in this.posts" :key="image.PostId"> 
						<img :id="image.PostId" :src="getImagePath(image.PostId)" style="width: 200px; height: auto;">
						<button class="deletePost-button" @click="deletePost(image.PostId)">X</button>
						<button class="buttonLike-scroll" @click="likePhoto(image.PostId)">Like</button> 
						<button class="buttonUnlike-scroll" @click="unlikePhoto(image.PostId)">unlike</button> 
						<button class="comments-scroll" @click="toggleCommentBox(image.PostId)">commments</button> 
						<p style="color:white;">Number of likes: {{this.posts[index].Likes?.length || 0}}</p>
						
						<p style="color:white;">Upload: {{formatDate(image.Time)}}</p>
					</li>
				</ul>
				<div class="comment-section" v-if="this.CommentBoxShowing"> 
					<button class="exit-button" @click="this.CommentBoxShowing=false">X</button> 
					<ul class="comment-container">
						<li class="comment"  v-for="comment in this.selectedPost?.Comments" :key="comment.CommentId"> 
							<p>{{ comment.Username }} : {{ comment.Message }}</p> 
							<button class="deleteComment-button" @click="deleteComment(comment.CommentId)" v-if="comment.Username==this.profile.Username">delete</button>
						</li>
					</ul>
				</div>
				<div v-if="this.CommentBoxShowing">
					<form class="comment-form" @submit.prevent="commentPost" > 
						<input type="text" id="commentPost" v-model="commentContent" required> <br>
						<input type="submit" value="Submit">
					</form>
				</div>
			</div>
			<div class="settings-form">
				<input v-model="newUsername" type="text" placeholder="Enter New Username">
        		<button @click="ChangeUsername">Submit</button>
			</div>
			<div class="ban-form">
				<input v-model="userToBan" type="text" placeholder="Ban user">
        		<button @click="banUser">Submit</button>
			</div>
			<div class="following-form">
				<ul>
					<p>following:</p>
					<li v-for="(following, index) in this.following" :key="index"> 
						<p>{{following.Username}}</p>
					</li>
				</ul>
			</div>
			<div class="follower-form">
				<ul>
					<p>followers:</p>
					<li v-for="(follower, index) in this.follower" :key="index"> 
						<p>{{follower.Username}}</p>
					</li>
				</ul>
			</div>
			<div class="banned-form">
				<ul>
					<p>banned:</p>
					<li v-for="(banned, index) in banned" :key="index" style="display: flex; align-items: center;">
						<p style="margin-right: 10px;">{{ banned }}</p>
						<button @click="unbanUser(banned)">Unban</button>
					</li>
				</ul>
			</div>
			<div class="uploadPost-form">
				<form @submit.prevent="uploadPost">
					<input type="file" @change="handleFileUpload" required>
					<button type="submit">Upload</button>
				</form>
			</div>

		</div>
		
        
	</div>
</template>

<style>
.deletePost-button {
	
	position: relative; 
	bottom:120px; 
	
}
.uploadPost-form {
	position: absolute;  
	top:152px;
	right:200px;
}
.ban-form {
	position: absolute;  
	top:202px;
	right:600px;
}
.banned-form {
	position: absolute;  
	top:200px;
	right:200px;
}
.following-form {
	position: absolute;  
	top:200px;
	right:500px;
}
.follower-form {
	position: absolute;  
	top:200px;
	right:350px;
}
.settings-form {
	position: absolute;  
	top:152px;
	right:600px;
}
.comment-form {
	position: absolute;  
	bottom:160px;
	left:638px;
}

.comment {
	margin-top: 10px;
	
	top:30px;
	left:10px;
	color:aliceblue;
}
.exit-button {
	
	left:210px;
	bottom: 270px;
}
.comment-section {
    position: absolute;  
	top:450px;
	left:630px;
	width: 250px;
	height: 300px;
	background-color: #3f3f3f;
	overflow-y: scroll;
	overflow-x: scroll;
}
 
.comments-scroll {
	position:relative;
	top: 70px;
	right: 90px;
}
.buttonUnlike-scroll {
	position:relative;
	bottom:0px;
	right: -5px;
}
.buttonDelete-scroll-container {
	position:relative;
	bottom:120px;
	color:red;
	right: -10px;
}
.buttonLike-scroll {
	position:relative;
	bottom:0px;
	right: -5px;
	
}
.scroll-container {
  overflow-x: hidden;
  overflow-y: scroll;
  background-color: #333;
  
  white-space: nowrap;
  padding: 10px;
  height: 550px;
  width: 400px;
}
</style>
