package service

//GetArticleRelationDetails 查询关联碎片
func GetArticleRelationDetails(detail ArticleFragmentRelationDetail) []ArticleFragmentRelationDetail {
	return queryArticleRelationDetails(detail)
}

//AddArticleRelation 添加关联碎片
func AddArticleRelation(detail ArticleFragmentRelationDetail) int {
	return addArticleRelation(detail)
}

//UpdateArticleRelations 修改顺序
func UpdateArticleRelations(details []ArticleFragmentRelationDetail) int {
	return updateArticleRelations(details)
}

//DeleteArticleRelationByID 删除关联关系
func DeleteArticleRelationByID(id int) int {
	return deleteArticleRelationByID(id)
}
