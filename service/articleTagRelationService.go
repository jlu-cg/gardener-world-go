package service

//AddArticleTagRelation 添加文章标签关系
func AddArticleTagRelation(relation ArticleTagRelation) int {
	return addArticleTagRelation(relation)
}

//UpdateArticleTagRelation 更新文章标签关系
func UpdateArticleTagRelation(relation ArticleTagRelation) int {
	return updateArticleTagRelation(relation)
}

//DeleteArticleTagRelationByID 删除文章标签关系
func DeleteArticleTagRelationByID(id int) int {
	return deleteArticleTagRelationByID(id)
}
