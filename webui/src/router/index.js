import {createRouter, createWebHashHistory} from 'vue-router'
import FeedView from '../views/FeedView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/users/profile', component: ProfileView},
		{path: '/profile/search', component: SearchView},
		{path: '/users/profile/feed', component: FeedView},
	]
})
export default router
