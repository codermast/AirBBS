import axios from "@/api/axios";
import type { RouteParamValue } from "vue-router";

const BASE_URL : string = "articles"


// GET /article/:aid
export const getArticleById = (articleId: string ) => {
	return axios.get(`${BASE_URL}/${articleId}`);
};
