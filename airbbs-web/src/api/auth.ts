// 校验模块 api

import axios from "@/api/axios";

// 这里的后缀必须加 / 否则重定向后，响应头会丢失
const BASE_URL = 'auths'


// POST /auths/register
export const sendRegisterAuthCode = (mail: string) => {
	const formData = new FormData();
	formData.append('mail', mail);
	return axios.post(`${ BASE_URL }/register`, formData);
};

// POST /auths/reset
export const sendResetAuthCode = (account: string) => {
	const formData = new FormData();
	formData.append('account', account);
	return axios.post(`${ BASE_URL }/reset`, formData);
};

// POST /auths/reset/match
export const matchResetAuthCode = (account: string, code: string) => {
	const formData = new FormData();
	formData.append('account', account);
	formData.append('code', code);
	return axios.post(`${ BASE_URL }/reset/match`, formData);
};