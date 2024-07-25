// 关注模块 api
import axios from "@/api/axios";

// 这里的后缀必须加 / 否则重定向后，响应头会丢失
const BASE_URL = 'follow'


// 关注指定用户
export const followUserToId = (userId: string) => {
	return axios.post(`${ BASE_URL }/${ userId }`)
}

// 取消关注指定用户
export const unfollowUserToId = (userId: string) => {
	return axios.delete(`${ BASE_URL }/${ userId }`)
}