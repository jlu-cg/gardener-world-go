package service

//ArticleFragmentAttributeRelation 文章碎片对应的可对比属性
type ArticleFragmentAttributeRelation struct {
	ID             int `json:"id"`
	FragmentID     int `json:"fragmentId"`
	ArticleID      int `json:"articleId"`
	TagAttributeID int `json:"tagAttributeId"`
}
