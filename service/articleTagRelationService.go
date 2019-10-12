package service

//QueryArticleTagRelationWithTags 查询文章对应的标签
func QueryArticleTagRelationWithTags(relation ArticleTagRelationWithTag) []ArticleTagRelationWithTag {
	return queryArticleTagRelationWithTags(relation)
}

//QueryArticleTagRelationWithArticles 查询标签对应的文章
func QueryArticleTagRelationWithArticles(relation ArticleTagRelationWithArticle) []ArticleTagRelationWithArticle {
	return queryArticleTagRelationWithArticles(relation)
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
	var relation ArticleTagRelation
	relation.ID = id
	return deleteArticleTagRelations(relation)
}

//DeleteArticleTagRelations 删除文章标签关系
func DeleteArticleTagRelations(relation ArticleTagRelation) int {
	return deleteArticleTagRelations(relation)
}
