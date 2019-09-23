package service

//QueryArticleTagRelations 查询标签对应的文章
func QueryArticleTagRelations(relation ArticleTagRelation) []ArticleTagRelation {
	if relation.ArticleID > 0 {
		return queryTagArticleTagRelations(relation)
	}
	if relation.TagID > 0 {
		return queryArticleTagRelations(relation)
	}
	var articleTagRelations []ArticleTagRelation
	return articleTagRelations
}

//SaveArticleTagRelation 保存文章标签关系
func SaveArticleTagRelation(relation ArticleTagRelation) int {
	if relation.ID > 0 {
		return updateArticleTagRelation(relation)
	}
	return addArticleTagRelation(relation)
}

//DeleteArticleTagRelationByID 删除文章标签关系
func DeleteArticleTagRelationByID(id int) int {
	return deleteArticleTagRelationByID(id)
}

//DeleteArticleTagRelations 删除文章标签关系
func DeleteArticleTagRelations(relation ArticleTagRelation) int {
	return deleteArticleTagRelations(relation)
}
