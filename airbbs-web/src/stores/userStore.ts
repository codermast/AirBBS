// 状态

// 引入 defineStore 用于创建 store
import { defineStore } from 'pinia'
import { ref, watch } from "vue"
import { HeaderAuthTokenName } from "@/config";
import type { UserQueryInfo } from "@/models/user";
// 定义并暴露一个 store
export const useUserStore = defineStore('status', () => {
	// 用户登录状态
	let userLoginStatus = ref(localStorage.getItem('userLoginStatus') === 'true')

	// 监听 userLoginStatus 的变化，更新到 LocalStorage 中
	watch(userLoginStatus, (newValue) => {
		console.log("userLoginStatus 更新", newValue)
		localStorage.setItem('userLoginStatus', newValue.toString());
	});

	// 当前登录用户的 ID
	let userId = ref(localStorage.getItem("userId") || "-1");

	// 监听 userLoginId 的变化，更新到 LocalStorage 中
	watch(userId, (newValue) => {
		console.log("userLoginId 更新", newValue)
		localStorage.setItem('userId', newValue);
	});

	// JWT Token
	let JWTToken = ref(localStorage.getItem(HeaderAuthTokenName) || "");
	// 持久化 JWT Token
	watch(JWTToken, (newValue) => {
		localStorage.setItem(HeaderAuthTokenName, newValue);
	});

	const storedUserInfo = localStorage.getItem("userInfo");
	let userInfo = ref(storedUserInfo ? JSON.parse(storedUserInfo) : null)

	// 持久化 userInfo
	watch(userInfo, (newValue: any) => {
		localStorage.setItem("userInfo", JSON.stringify(newValue));
	});

	// 退出登录
	function logout() {
		userLoginStatus.value = false;
		userInfo.value = '';
		JWTToken.value = '';
		localStorage.removeItem('userLoginStatus');
		localStorage.removeItem('userId');
		localStorage.removeItem(HeaderAuthTokenName);
	}

	// 需要手动进行返回暴露，这和 Vue3 中的 setup 有些差别
	return {userLoginStatus, userInfo, JWTToken, userId, logout};
})