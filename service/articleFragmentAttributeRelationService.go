package service

//AddArticleFragmentAttributeRelation 添加文章碎片属性关系
func AddArticleFragmentAttributeRelation(relation ArticleFragmentAttributeRelation) int {
	return addArticleFragmentAttributeRelation(relation)
}

//UpdateArticleFragmentAttributeRelation 更新文章碎片属性关系
func UpdateArticleFragmentAttributeRelation(relation ArticleFragmentAttributeRelation) int {
	return updateArticleFragmentAttributeRelation(relation)
}

//DeleteArticleFragmentAttributeRelationByID 删除文章碎片属性关系
func DeleteArticleFragmentAttributeRelationByID(id int) int {
	return deleteArticleFragmentAttributeRelationByID(id)
}
