package service

//GetArticleDependenceDetails 查询依赖文章
func GetArticleDependenceDetails(detail ArticleArticleRelationDetail) []ArticleArticleRelationDetail {
	return queryArticleDependenceDetails(detail)
}

//AddArticleDependence 添加依赖文章
func AddArticleDependence(detail ArticleArticleRelationDetail) int {
	return addArticleDependence(detail)
}

//UpdateArticleDependences 修改顺序
func UpdateArticleDependences(details []ArticleArticleRelationDetail) int {
	return updateArticleDependences(details)
}

//DeleteArticleDependenceByID 删除依赖文章
func DeleteArticleDependenceByID(id int) int {
	return deleteArticleDependenceByID(id)
}
