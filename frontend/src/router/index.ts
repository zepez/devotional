import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/HomeView.vue";


const routes: Array<RouteRecordRaw> = [
	{
		path: "/",
		component: Home,
	},
	{
		path: "/devotional",
		component: () => import(/* webpackChunkName: "devotional" */ "../views/DevotionalView.vue"),
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});


export default router;
