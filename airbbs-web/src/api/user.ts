// 用户模块 api

import axios from "@/api/axios";
import type { UserInfo } from "@/models/user";

// 这里的后缀必须加 / 否则重定向后，响应头会丢失
const BASE_URL = 'users'

// GET /users/:uid
export const getUser = (userId : string) => {
	return axios.get(`${BASE_URL}/${userId}`);
};

// GET /users/
export const getUsers = () => {
	return axios.get(`${BASE_URL}/`);
}

// POST /users/
export const registerUser = (user : UserInfo) => {

	return axios.post(`${BASE_URL}/`, user)
}