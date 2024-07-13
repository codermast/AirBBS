// 状态

// 引入 defineStore 用于创建 store
import { defineStore } from 'pinia'
import { ref,watch} from "vue"
// 定义并暴露一个 store
export const useStatusStore = defineStore('status', () => {
	// 用户登录状态
	let userLoginStatus = ref(localStorage.getItem('userLoginStatus') === 'true')

	// 监听 userLoginStatus 的变化，更新到 LocalStorage 中
	watch(userLoginStatus, (newValue) => {
		localStorage.setItem('userLoginStatus', newValue.toString());
	});
	
	// 当前登录用户的 ID
	let userLoginId  = ref(localStorage.getItem("userId") || "27");

	// 监听 userLoginId 的变化，更新到 LocalStorage 中
	watch(userLoginId, (newValue) => {
		localStorage.setItem('userId', newValue.toString());
	});
	

	// 需要手动进行返回暴露，这和 Vue3 中的 setup 有些差别
	return {userLoginStatus,userLoginId};
})