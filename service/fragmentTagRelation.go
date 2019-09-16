package service

//FragmentTagRelation 碎片标签
type FragmentTagRelation struct {
	ID         int `json:"id"`
	FragmentID int `json:"fragmentId"`
	TagID      int `json:"tagId"`
}
