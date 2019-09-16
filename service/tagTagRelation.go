package service

//TagTagRelation 标签间关系
type TagTagRelation struct {
	ID             int `json:"id"`
	TagID          int `json:"tagId"`
	RelateTagID    int `json:"relateTagId"`
	Classification int `json:"classification"`
}
