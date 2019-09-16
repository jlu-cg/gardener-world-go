package service

//GetArticleDependenceDetails 查询依赖文章
func GetArticleDependenceDetails(detail ArticleDependenceDetail) []ArticleDependenceDetail {
	return queryArticleDependenceDetails(detail)
}

//AddArticleDependence 添加依赖文章
func AddArticleDependence(detail ArticleDependenceDetail) int {
	return addArticleDependence(detail)
}

//UpdateArticleDependences 修改顺序
func UpdateArticleDependences(details []ArticleDependenceDetail) int {
	return updateArticleDependences(details)
}

//DeleteArticleDependenceByID 删除依赖文章
func DeleteArticleDependenceByID(id int) int {
	return deleteArticleDependenceByID(id)
}
