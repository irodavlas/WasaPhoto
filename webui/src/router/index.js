import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import FollowPage from '../views/FollowPage.vue'
import Profile from '../views/Profile.vue'
import Settings from '../views/Settings.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/home', component: HomeView},
		{path: '/profile', component: Profile},
		{path: '/network', component: FollowPage},
		{path: '/settings', component: Settings},

	]
})


export default router
