import axios from "@/api/axios";
import type { Article } from "@/models/article";

const BASE_URL: string = "articles"


// GET /article/:aid
export const getArticleById = (articleId: string) => {
	return axios.get(`${ BASE_URL }/${ articleId }`);
};


// POST /article/
export const createArticle = (article: Article) => {
	return axios.post(`${ BASE_URL }/`, article);
}