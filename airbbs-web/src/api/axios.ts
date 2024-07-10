// 在 src 目录下创建一个 api 文件夹，并创建一个 axios.js 文件
import axios from 'axios';

const instance = axios.create({
	baseURL: 'http://localhost:8080/', // 你的后端 API 地址
	timeout: 10000, // 请求超时时间
	headers: {
		'Content-Type': 'application/json',
	},
});

export default instance;
