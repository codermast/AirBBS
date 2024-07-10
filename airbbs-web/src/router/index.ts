import { createRouter, createWebHistory } from 'vue-router';
import Home from "@/pages/Home.vue"
import ArticleList from "@/components/ArticleList.vue";
import ArticleInfo from "@/components/ArticleInfo.vue";

const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home,
		children : [
			{
				path : "/",
				name : "Index",
				component : ArticleList
			},
			{
				path : "/articles/:articleID",
				name : "ArticleInfo",
				component : ArticleInfo
			}
		]
	}

];

const router = createRouter({
	history: createWebHistory(),
	routes
});

export default router;