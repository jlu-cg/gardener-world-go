package service

//QueryArticleFragmentRelationDetails 查询关联碎片
func QueryArticleFragmentRelationDetails(detail ArticleFragmentRelationDetail) []ArticleFragmentRelationDetail {
	return queryArticleFragmentRelationDetails(detail)
}

//AddArticleFragmentRelation 添加关联碎片
func AddArticleFragmentRelation(detail ArticleFragmentRelationDetail) int {
	return addArticleFragmentRelation(detail)
}

//UpdateArticleFragmentRelations 修改顺序
func UpdateArticleFragmentRelations(details []ArticleFragmentRelationDetail) int {
	return updateArticleFragmentRelations(details)
}

//DeleteArticleFragmentRelationByID 删除关联关系
func DeleteArticleFragmentRelationByID(id int) int {
	return deleteArticleFragmentRelationByID(id)
}
