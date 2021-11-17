import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/HomeView.vue";


const routes: Array<RouteRecordRaw> = [
	{
		path: "/",
		name: "Home",
		component: Home,
	},
	{
		path: "/d/:id",
		name: "devotional",
		component: () => import(/* webpackChunkName: "devotional" */ "../views/DevotionalView.vue"),
	},
	{
		path: "/:pathMatch(.*)*", 
		name: "NotFound",
		component: () => import("../views/404View.vue"),
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});


export default router;
