<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
            searchUsername: null,
            profile: null,
            posts:[],
            follower:[],
			banned:[],
            token:null,
            
            CommentBoxShowing:false,
			selectedPost:null,
			commentContent:null,

		}
	},
	methods: {
        usernameInFollowerList(_followerUsername){
			if (!Array.isArray(this.follower)) {
				this.follower = [];
				return false;
			}
			for (let i = 0; i < this.follower.length; i++) {
				if (this.follower[i].Username == _followerUsername){
					return true;
				}
				
			}
			
			return false;
		},
		
		async unFollowUser(){
			this.loading = true;
			this.errormsg = null;


            const token = localStorage.getItem('token');
            
            
            
            const headers = {
                'Authorization': 'Bearer ' + token
            };
           
            var url = '/follow/' + this.profile.Id;
			try {
				let response = await this.$axios.delete(url,  {headers: headers});
				
                if (this.usernameInFollowerList(this.username)){
					for (let i = 0; i < this.follower.length; i++) {
						if (this.follower[i].Id == token){
							this.follower.splice(i, 1);
						}
						
					}
			   	}
			    
				console.log(this.follower);
                
			} catch (e) {
				this.errormsg = e;
			}
			
			
			
			this.notfound = false;
			this.loading = false;
		},
        async followUser(){
            this.loading = true;
			this.errormsg = null;


            const token = localStorage.getItem('token');
            
            
            
            const headers = {
                'Authorization': 'Bearer ' + token
            };
           
            var url = '/follow/' + this.profile.Id;
			try {
				let response = await this.$axios.put(url, null, {headers: headers});
				
                if (!this.usernameInFollowerList(this.username)){
					this.follower.push({'Id':token, 'Username':this.username});
			   	}
			    
				console.log(this.follower);
                
			} catch (e) {
				this.errormsg = e;
			}
			
			
			
			this.notfound = false;
			this.loading = false;
        },
        getUserInfo(){
            this.token = localStorage.getItem('token');
            this.username = sessionStorage.getItem(this.token);
            
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
				
				for (let i = 0; i < this.posts?.length; i++) {
					if (!this.posts[i].Comments) {
						this.posts[i].Comments = [];
					}
					if (this.posts[i].PostId == this.selectedPost.PostId){
						this.posts[i].Comments.push({'CommentId': response.data.message.CommentId,'Username': this.username, 'Message':this.commentContent});
					}
				}
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			
			this.loading = false;
			this.commentContent = null;
			
			
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
							
							if (this.posts[i].Likes.includes(this.username)) {
								const index = this.posts[i].Likes.indexOf(this.username);
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
						if (!this.posts[i].Likes.includes(this.username)) {
							this.posts[i].Likes.push(this.username);	
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
        formatDate(time){
			const date = new Date(time);
			const day = String(date.getDate()).padStart(2, '0');
			const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
			const year = date.getFullYear();
			const hours = String(date.getHours()).padStart(2, '0');
            const minutes = String(date.getMinutes()).padStart(2, '0');

			return `${day}/${month}/${year} ${hours}:${minutes}`;
		},
        getImagePath(photoPath) {
			
			const basePath = "images/"+ photoPath + ".jpg";

			return basePath
		},
		async searchProfile(){
            this.loading = true;
			this.errormsg = null;


            const token = localStorage.getItem('token');
            
            
            const queryParams = {
                'username': this.searchUsername
            }
            const headers = {
                'Authorization': 'Bearer ' + token
            };
            var url = '/users/' + token + '/profile';
			try {
				let response = await this.$axios.get(url,  {params: queryParams, headers: headers});
				this.profile = response.data.message;
				this.posts = this.profile.Post;
				this.follower = this.profile.Follower;
				
                
			} catch (e) {
				this.errormsg = "User not found";
				this.profile = '';
			}
			
			
         
			this.notfound = false;
			this.loading = false;
			
		},
        toggleCommentBox(postId){
			
			for (let i = 0; i < this.posts?.length; i++) {
				if(this.posts[i].PostId == postId){
					this.selectedPost = this.posts[i];
				}
			}
			
			this.CommentBoxShowing = true;
		},
		
        
	},
    mounted(){
		const logged = sessionStorage.getItem('logged');
		
		if (!logged) {
			console.log("log in before accessing Search");
			this.$router.push('/');

		}
		else {
			this.getUserInfo()
		}
        
        
    },  
	
}
</script>

<template>
    <div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search page</h1>
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
        <form @submit.prevent="searchProfile">
            <input type="text" v-model="this.searchUsername" placeholder="Search Username" required />
            <button type="submit">Search</button>
        </form>
    </div>
    <div v-if="this.profile">
        <h4>Profile Information</h4>
        <p>Username: {{ this.profile.Username }}</p>
        <p>Number of posts: {{this.posts?.length || 0}}</p>
    </div>
	<div v-if="this.profile && this.profile.Username != this.username && !usernameInFollowerList(this.username)">
		<button  @click="followUser()">follow</button> <br>
	</div>
	<div v-if="this.profile && this.profile.Username != this.username && usernameInFollowerList(this.username)">
		<button  @click="unFollowUser()">unfollow</button> <br>
	</div>

	
    
    <div class="scroll-container" v-if="this.profile"> 
        <ul>
            <li v-for="(image, index) in this.posts" :key="image.PostId"> 
                <img :id="image.PostId" :src="getImagePath(image.PostId)" style="width: 200px; height: auto;">
                <button class="buttonLike-scroll" @click="likePhoto(image.PostId)">Like</button> 
                <button class="buttonUnlike-scroll" @click="unlikePhoto(image.PostId)">unlike</button> 
                <button class="comments-scroll" @click="toggleCommentBox(image.PostId)">commments</button> 
                <p style="color:white;">Number of likes: {{this.posts[index].Likes?.length || 0}}</p>
                
                <p style="color:white;">Upload: {{formatDate(image.Time)}}</p>
            </li>
        </ul>
        <div class="comment-section" v-if="this.CommentBoxShowing"> 
            <button class="exit-button" @click="this.CommentBoxShowing=false">X</button> 
            <ul class="comment-container" >
                <li class="comment"  v-for="comment in this.selectedPost?.Comments" :key="comment.CommentId"> 
                    <p>{{ comment.Username }} : {{ comment.Message }}</p> 
                    <button class="deleteComment-button" @click="deleteComment(comment.CommentId)" v-if="comment.Username==this.username">delete</button>
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
</template>

<style>
.comment {
	margin-top: 10px;
	
	top:30px;
	left:10px;
	color:aliceblue;
}
.comment-form {
	position: absolute;  
	bottom:160px;
	left:638px;
}
.deletePost-button {
	
	position: relative; 
	bottom:120px; 
	
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
.scroll-container {
    overflow-x: hidden;
    overflow-y: scroll;
    background-color: #333;
    
    white-space: nowrap;
    padding: 10px;
    height: 500px;
    width: 400px;
}

.buttonLike-scroll {
position:relative;
bottom:0px;
right: -5px;
	
}
.buttonUnlike-scroll {
	position:relative;
	bottom:0px;
	right: -5px;
}
</style>
