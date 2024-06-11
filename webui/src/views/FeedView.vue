<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			isLoggedIn:false,
			feed:null,
			commentContentFeed: null,
			username: null

			
		}
	},
	mounted(){
		const logged = sessionStorage.getItem('logged');
		
		if (!logged) {
			console.log("log in before accessing Feed");
			this.$router.push('/');

		}
		else{
			
			this.getFeed();
			
		}
		

		
	},
	methods: {
		
		formatDate(time){
			const date = new Date(time);
			
			return date.toLocaleDateString('en-US', { weekday: 'long' }) + " | " + date.toLocaleTimeString('en-US', { timeStyle: 'short' })
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
				for (let i = 0; i < this.feed?.length; i++) {
					if (this.feed[i].PostId == postId){
						
						if (this.feed[i].Likes && this.feed[i].Likes.length > 0) {
							
							if (this.feed[i].Likes.includes(this.username)) {
								
								const index = this.feed[i].Likes.indexOf(this.username);
							
								this.feed[i].Likes.splice(index, 1);	
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
		async deleteComment(_postId, _commentId){
			this.loading = true;
			this.errormsg = false;
			
			
			
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + _postId + '/comments/' + _commentId;
			try {
				let response = await this.$axios.delete(url, {headers:headers});
				console.log(response.data.message);

				if (response.data.code == "200") {
					for (let i = 0; i < this.feed?.length; i++) {
						if (this.feed[i].PostId == _postId){
							if (this.feed[i].Comments && this.feed[i].Comments.length > 0) {
								for (let j = 0; j < this.feed[i].Comments.length; j++) {
									
									if (this.feed[i].Comments[j].CommentId === _commentId) {
										this.feed[i].Comments.splice(j, 1); 
										break; 
									}
								}
							}
						}
					}
				}

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			
			

			this.loading = false;
			this.follow = '';
			this.commentContent = '';
			
			
		
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
				
				for (let i = 0; i < this.feed?.length; i++) {
					if (this.feed[i].PostId == postId){
						if (!this.feed[i].Likes) {
							this.feed[i].Likes = [];
						}
						if (!this.feed[i].Likes.includes(this.username)) {
							this.feed[i].Likes.push(this.username);	
						}
					}
				
				}
				
				
				
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
		async commentPost(_postId){
			this.loading = true;
			this.errormsg = false;
			
			const commentMessage = this.commentContentFeed;
			
			const token = localStorage.getItem('token');
			const headers = {
				'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token
            };
			const url = '/posts/' + _postId + '/comment';
			try {
				let response = await this.$axios.post(url, commentMessage, {headers:headers});
				
				for (let i = 0; i < this.feed?.length; i++) {
					if (!this.feed[i].Comments) {
						this.feed[i].Comments = [];
					}
					if (this.feed[i].PostId == _postId){
						this.feed[i].Comments.push({'CommentId': response.data.message.CommentId,'Username': this.username, 'Message':commentMessage});
					}
				}
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			
			
			this.loading = false;
			this.commentContent = null;
			
			
		},
		async getFeed(){

			this.loading = true;
			this.errormsg = false;
			
			
			
			const token = localStorage.getItem('token');
			
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.get("/profile/feed",  {headers});
				if (response.data.code == "200") {
					this.feed = response.data.message;
					this.username = sessionStorage.getItem(token);
					
                	console.log(this.feed);
				}
				

			} catch (e) {
				this.errormsg = e.toString()
			}
			
			
        },
    }
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			
			<h1> Your Feed </h1>
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
	<div v-if="!this.feed"> 
		<h6>Empty feed</h6>
	</div>
	<div v-if="this.feed"> 
        <ul>
            <li class="post-container" v-for="post in this.feed" :key="post.PostId"> <br>
                <h6>Post owner: {{ post.OwnerUsername }}</h6>
				<p>Number of comments: {{ post && post.Comments ? post.Comments.length : 0 }}</p>
				<p>Number of likes: {{ post && post.Likes ? post.Likes.length : 0 }}</p>
				
                <img :id="post.PostId" :src="getImagePath(post.PostId)" style="width: 200px; height: auto;">
				<p>Upload: {{formatDate(post.Time)}}</p>
				
				<div class="like-container"> 
					<button  @click="likePhoto(post.PostId)">Like</button> 
					<button  @click="unlikePhoto(post.PostId)">unlike</button> 

					<form @submit.prevent="commentPost(post.PostId)" > 
						<input type="text" id="commentPost" v-model="commentContentFeed" required> <br>
						<input type="submit" value="Submit">
				</form>
				</div>
				
				
                

                
				<ul class="comment-container-feed">
					<li class="comment" v-for="comment in post?.Comments" :key="comment['CommentId']"> 
						<p>{{ comment["Username"] }} : {{ comment["Message"] }} </p> 
						<button  @click="deleteComment(post.PostId, comment['CommentId'])" v-if="comment['Username'] ==this.username">delete</button>
					</li>
				</ul>
			
				
			

            </li>
        </ul>
		
    </div>
	
	
</template>

<style>
.post-container {
	
	height: 500px;
}
.like-container {
	position: relative;
	
}

.comment-container-feed{
	position: relative;
	bottom:385px;
	left:250px;
	width: 250px;
	height: 280px;
	background-color: #3f3f3f;
	overflow-y: scroll;
	overflow-x: scroll;
}
.comment {
	margin-top: 10px;
	
	top:30px;
	left:10px;
	color:aliceblue;
}
</style>