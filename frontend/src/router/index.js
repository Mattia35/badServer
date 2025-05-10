import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/loginView.vue'
import HomeView from '../views/homeView.vue'
import employeesView from '../views/employeesView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/:profile', component: HomeView},
		{path: '/:profile/employees', component: employeesView},
	]
})

export default router
