// 校验模块 api

import axios from "@/api/axios";

// 这里的后缀必须加 / 否则重定向后，响应头会丢失
const BASE_URL = 'auths'


// POST /auths/
export const sendAuthCode = (mail : string) => {
	const formData = new FormData();
	formData.append('mail', mail);
	return axios.post(`${BASE_URL}/`,formData);
};