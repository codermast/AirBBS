import { createRouter, createWebHistory } from 'vue-router';
import Home from "@/pages/Home.vue"
import ArticleList from "@/components/article/ArticleList.vue";
import ArticleInfo from "@/components/article/ArticleInfo.vue";
import Setting from "@/pages/Setting.vue";
import Login from "@/pages/Login.vue";
import Register from "@/pages/Register.vue";
import UserInfo from "@/components/setting/UserInfo.vue";
import UserAvatar from "@/components/setting/UserAvatar.vue";
import UserSocialBind from "@/components/setting/UserSocialBind.vue";
import ArticleCreate from "@/components/article/ArticleCreate.vue";
import emitter from "@/utils/emitter"

const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home,
		children: [
			{
				path: "/",
				name: "Index",
				component: ArticleList
			},
		]
	},
	{
		path: '/articles',
		name: 'ArticleInfo',
		component: Home,
		children: [
			{
				path: ":articleID",
				name: "ArticleInfo",
				component: ArticleInfo
			},

		]
	},
	{
		path: '/articles/create',
		name: 'ArticleCreate',
		component: ArticleCreate
	},
	{
		path: "/setting",
		name: 'Setting',
		component: Setting,
		redirect: {name: 'SettingUserInfo'},
		children: [
			{
				path: "user",
				name: "SettingUserInfo",
				children: [
					{
						path: "info",
						name: "SettingUserInfo",
						component: UserInfo
					},
					{
						path: "avatar",
						name: "SettingUserAvatar",
						component: UserAvatar
					},
					{
						path: "social",
						name: "SettingUserSocial",
						component: UserSocialBind
					},
				]
			},
		]
	},

	{
		path: "/login",
		name: 'Login',
		component: Login
	},

	{
		path: "/register",
		name: 'Register',
		component: Register
	},

];

const router = createRouter({
	history: createWebHistory(),
	routes
});

router.beforeEach((to, from) => {
	emitter.emit("loadingBarStart")
	return true
})

router.afterEach((to, from) => {
	emitter.emit("loadingBarFinish")

});

export default router;