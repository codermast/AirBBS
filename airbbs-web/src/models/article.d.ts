export type Article = {
    id: string
    title: string
    content: string
    publish_time: string
    update_time: string
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

export  type AuthorInfo = {
    id: string
    nickname: string
    username: string
    photo : string
    introduce: string
    articleTotal: string
    fansTotal: string
    viewTotal: string
}
