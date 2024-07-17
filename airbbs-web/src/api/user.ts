// 用户模块 api

import axios from "@/api/axios";
import type { UserLoginInfo, UserQueryInfo, UserRegisterInfo, UserResetPasswordInfo } from "@/models/user";

// 这里的后缀必须加 / 否则重定向后，响应头会丢失
const BASE_URL = 'users'

// GET /users/:uid
export const getUserById = (userId: string) => {
	return axios.get(`${ BASE_URL }/${ userId }`);
};

// GET /users/
export const getUsers = () => {
	return axios.get(`${ BASE_URL }/`);
}

// POST 用户注册 /users/register
export const registerUser = (user: UserRegisterInfo) => {
	return axios.post(`${ BASE_URL }/register`, user)
}

// POST 用户登录 /users/login
export const loginUser = (user: UserLoginInfo) => {
	return axios.post(`${ BASE_URL }/login`, user)
}

// PUT 更新用户基本信息 /users/ 根据传进来的 userId 来更新
export const updateUserInfo = (user: UserQueryInfo) => {
	return axios.put(`${ BASE_URL }/`, user)
}

// POST 重置用户密码 /users/password/reset
export const resetUserPassword = (userResetPasswordInfo: UserResetPasswordInfo) => {
	return axios.post(`${ BASE_URL }/password/reset`, userResetPasswordInfo)
}


// POST 上传头像 /users/photo
export const uploadUserPhoto = (file: File) => {
	const formData = new FormData();
	formData.append('file', file)
	return axios.post(`${ BASE_URL }/photo`, formData)
}