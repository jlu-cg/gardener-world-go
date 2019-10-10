package service

//QueryTagArticleFragmentRelationDetails 查询标签关联关系
func QueryTagArticleFragmentRelationDetails(relation TagArticleFragmentRelationDetail) []TagArticleFragmentRelationDetail {
	return queryTagArticleFragmentRelationDetails(relation)
}

//SaveTagArticleFragmentRelation 添加标签间关联
func SaveTagArticleFragmentRelation(relation TagArticleFragmentRelation) int {
	return addTagArticleFragmentRelation(relation)
}

//DeleteTagArticleFragmentRelationByID 删除标签间关联
func DeleteTagArticleFragmentRelationByID(id int) int {
	var relation TagArticleFragmentRelation
	relation.ID = id
	return deleteTagArticleFragmentRelations(relation)
}

//DeleteTagArticleFragmentRelations 删除标签间关联
func DeleteTagArticleFragmentRelations(relation TagArticleFragmentRelation) int {
	return deleteTagArticleFragmentRelations(relation)
}
