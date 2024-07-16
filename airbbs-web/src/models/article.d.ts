export type Article = {
	id: string
	title: string
	content: string
	status: number
	author: string
}

export type ArticlePageRequest = {
	pageNumber: number
	pageSize: number
}


export type ArticleListPage = {
	pageNumber: int
	pageSize: int
	pageCount: int
	totalCount: int
	articles: Article[]
}
