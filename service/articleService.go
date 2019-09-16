package service

//GetArticles 查询碎片
func GetArticles(article Article, lastID int) []Article {
	if lastID < 0 {
		lastID = 0
	}
	articles := queryArticles(article, lastID)
	return articles
}

//SaveArticle 保存碎片
func SaveArticle(article Article) int {
	if article.ID > 0 {
		updateArticle(article)
	} else {
		addArticle(article)
	}
	return 0
}

//GetArticleByID 查询碎片
func GetArticleByID(articleID int, tagID int) Article {
	if articleID > 0 {
		return queryArticleByID(articleID)
	}
	var article Article
	article.TagID = tagID
	return article
}

//DeleteArticleByID 删除碎片
func DeleteArticleByID(articleID int) int {
	if articleID > 0 {
		return deleteArticleByID(articleID)
	}
	return 0
}
