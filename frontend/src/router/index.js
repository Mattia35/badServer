import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/loginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
	]
})

export default router
