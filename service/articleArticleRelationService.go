package service

//GetArticleArticleRelationDetails 查询依赖文章
func GetArticleArticleRelationDetails(detail ArticleArticleRelationDetail) []ArticleArticleRelationDetail {
	return queryArticleArticleRelationDetails(detail)
}

//AddArticleArticleRelation 添加依赖文章
func AddArticleArticleRelation(detail ArticleArticleRelationDetail) int {
	return addArticleArticleRelation(detail)
}

//UpdateArticleArticleRelations 修改顺序
func UpdateArticleArticleRelations(details []ArticleArticleRelationDetail) int {
	return updateArticleArticleRelations(details)
}

//DeleteArticleArticleRelationByID 删除依赖文章
func DeleteArticleArticleRelationByID(id int) int {
	return deleteArticleArticleRelationByID(id)
}
