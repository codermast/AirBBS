import axios from "@/api/axios";
import type { BlinkCreateRequest } from "@/models/ro/blink";

const BASE_URL: string = "blink"

export const createBlink = (blinkCreateRequest: BlinkCreateRequest) => {
	return axios.post(`${ BASE_URL }/`, blinkCreateRequest);
}

// 动态查看
export const getBlinkList = () => {
	return axios.get(`${ BASE_URL }/list`);
}