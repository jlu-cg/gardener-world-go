package service

//QueryArticles 查询文章
func QueryArticles(article ArticleWithTag, lastID int) []ArticleWithTag {
	if article.TagArticleID > 0 {
		return queryArticleWithTags(article, lastID)
	}
	return queryArticles(article.Article, lastID)
}

//SaveArticle 保存文章
func SaveArticle(article Article) int {
	if article.ID > 0 {
		updateArticle(article)
	} else {
		addArticle(article)
	}
	return 0
}

//GetArticleByID 查询文章
func QueryArticleByID(articleID int) ArticleWithTag {
	if articleID > 0 {
		var queryArticle Article
		queryArticle.ID = articleID
		articles := queryArticles(queryArticle, 0)
		if len(articles) > 0 {
			return articles[0]
		}
	}
	var article ArticleWithTag
	return article
}

//DeleteArticleByID 删除文章
func DeleteArticleByID(articleID int) int {
	if articleID <= 0 {
		return 0
	}

	code := deleteArticleByID(articleID)

	var deleteArticleTagRelation ArticleTagRelation
	deleteArticleTagRelation.ArticleID = articleID
	code = deleteArticleTagRelations(deleteArticleTagRelation)

	return code
}
