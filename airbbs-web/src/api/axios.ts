import axios from 'axios';
import { BackendUrl, HeaderAuthTokenName } from "@/config";
import { useUserStore } from '@/stores/userStore';
import pinia from "@/stores/pinia";
const userStore = useUserStore(pinia);

const instance = axios.create({
	baseURL: BackendUrl, // 你的后端 API 地址
	timeout: 10000, // 请求超时时间
	validateStatus: (status) => {
		return status >= 200 && status < 300;
	}
});

// 设置动态请求头
instance.interceptors.request.use(
	config => {
		// 每次请求时动态设置 Authorization 头部
		const token = userStore.JWTToken; // 获取最新的 Token

		if (token) {
			config.headers['Authorization'] = `Bearer ${token}`;
		}

		return config;
	},
	error => {
		return Promise.reject(error);
	}
);


instance.interceptors.response.use(
	response => {
		if (response.status != 200) {
			// 在这里处理非 200 错误响应，不打印控制台错误信息
			console.warn('业务异常:', response.data.message);
		}
		// 处理成功响应
		return response;
	},
	error => {
		// 检查响应是否存在
		if (error.response) {
			return Promise.resolve(error.response);
		} else {
			// 处理网络错误
			// 返回一个假响应
			return Promise.resolve({data: null});
		}
	}
);


export default instance;
