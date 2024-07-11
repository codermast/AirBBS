// 用户模块 api

import axios from "@/api/axios";
import type { UserInfo } from "@/models/user";

const BASE_URL = 'users'

// GET /users/:uid
export const getUser = (userId : string) => {
	return axios.get(`${BASE_URL}/${userId}`);
};

// GET /users
export const getUsers = () => {
	return axios.get(BASE_URL);
}

// POST /users/register
export const registerUser = (user : UserInfo) => {
	return axios.post(`${BASE_URL}`, user)
}