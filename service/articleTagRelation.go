package service

//ArticleTagRelation 文章标签关系
type ArticleTagRelation struct {
	ID        int `json:"id"`
	ArticleID int `json:"articleId"`
	TagID     int `json:"tagId"`
}
