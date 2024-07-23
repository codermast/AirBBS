package vo

type AuthorVo struct {
	ID           string `json:"id"`
	Username     string `json:"username"`     // 用户名
	Nickname     string `json:"nickname"`     // 昵称
	Photo        string `json:"photo"`        // 头像
	Introduce    string `json:"introduce"`    // 个人介绍
	ArticleTotal int    `json:"articleTotal"` // 文章总数
	ViewTotal    int    `json:"viewTotal"`    // 总访问量
	FansTotal    int    `json:"fansTotal"`    // 粉丝量
}
