package service

//QueryTagArticles 查询标签列表
func QueryTagArticles(queryTagArticle TagArticle, lastID int) []TagArticle {
	if lastID < 0 {
		lastID = 0
	}
	tagArticles := queryTagArticles(queryTagArticle, lastID)
	return tagArticles
}

//SaveTagArticle 保存类别
func SaveTagArticle(tagArticle TagArticle) int {
	if tagArticle.ID > 0 {
		updateTagArticle(tagArticle)
	} else {
		addTagArticle(tagArticle)
	}
	return 0
}

//QueryTagArticleByID 通过ID查询
func QueryTagArticleByID(tagID int) TagArticle {
	if tagID > 0 {
		var queryTagArticle TagArticle
		queryTagArticle.ID = tagID
		tags := queryTagArticles(queryTagArticle, 0)
		if len(tags) > 0 {
			return tags[0]
		}
	}
	var tagArticle TagArticle
	return tagArticle
}

//DeleteTagArticleByID 通过ID删除标签
func DeleteTagArticleByID(id int) int {
	code := deleteTagArticleByID(id)

	var deleteArticleTagRelation ArticleTagRelation
	deleteArticleTagRelation.TagArticleID = id
	code = DeleteArticleTagRelations(deleteArticleTagRelation)

	var deleteTagArticleFragmentRelation TagArticleFragmentRelation
	deleteTagArticleFragmentRelation.TagArticleID = id
	code = DeleteTagArticleFragmentRelations(deleteTagArticleFragmentRelation)

	return code
}
