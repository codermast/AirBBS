// 状态

// 引入 defineStore 用于创建 store
import { defineStore } from 'pinia'
import { ref, watch } from "vue"
import { HeaderAuthTokenName } from "@/config";
// 定义并暴露一个 store
export const useStatusStore = defineStore('status', () => {
	// 用户登录状态
	let userLoginStatus = ref(localStorage.getItem('userLoginStatus') === 'true')

	// 监听 userLoginStatus 的变化，更新到 LocalStorage 中
	watch(userLoginStatus, (newValue) => {
		console.log("userLoginStatus 更新", newValue)
		localStorage.setItem('userLoginStatus', newValue.toString());
	});

	// 当前登录用户的 ID
	let userLoginId = ref(localStorage.getItem("userId") || "-1");

	// 监听 userLoginId 的变化，更新到 LocalStorage 中
	watch(userLoginId, (newValue) => {
		console.log("userLoginId 更新", newValue)
		localStorage.setItem('userId', newValue);
	});

	// JWT Token
	let JWTToken = ref(localStorage.getItem(HeaderAuthTokenName) || "");


	// 监听 JWT Token 的变化，更新到 LocalStorage 中
	watch(JWTToken, (newValue) => {
		console.log("JWTToken 更新", newValue)
		localStorage.setItem(HeaderAuthTokenName, newValue);
	});

	// 需要手动进行返回暴露，这和 Vue3 中的 setup 有些差别
	return {userLoginStatus, userLoginId, JWTToken};
})