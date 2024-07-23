import axios from "@/api/axios";
import type { Article, ArticlePageRequest } from "@/models/article";

const BASE_URL: string = "article"

// GET /articles/all 获取所有文章
export const getArticleAllList = () => {
	return axios.get(`${ BASE_URL }/all`);
};

// GET /articles/publish 获取公开文章
export const getArticlePublicList = () => {
	return axios.get(`${ BASE_URL }/publish`);
};

// GET /articles/:aid
export const getArticleById = (articleId: string) => {
	return axios.get(`${ BASE_URL }/${ articleId }`);
};

// POST /articles/
export const createArticle = (article: Article) => {
	return axios.post(`${ BASE_URL }/`, article);
}

// DELETE /articles/:id
export const deleteArticle = (articleId: string) => {
	return axios.delete(`${ BASE_URL }/${ articleId }`)
}

// PUT /articles/:id
export const updateArticle = (article: Article) => {
	return axios.put(`${ BASE_URL }/${ article.id }`, article)
}

// PUT /articles/
export const updateArticleListStatus = (articleIdList: string[], status: number) => {
	return axios.put(`${ BASE_URL }/`, {
		ids: articleIdList,
		status: status
	})
}

// GET /articles/page
export const getArticleListPage = (articlePageRequest: ArticlePageRequest) => {
	return axios.get(`${ BASE_URL }/page`, { params : articlePageRequest })
}

// GET /articles/author/:id
export const getAuthorInfo = (authorId : string) => {
	return axios.get(`${ BASE_URL }/author/${ authorId }`)
}