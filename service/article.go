package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

//Article 文章
type Article struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}
type ArticleWithTag struct {
	Article
	TagArticleID int `json:"tagArticleId"`
}

const (
	queryArticleWithTagsSQL = "select a.id, a.title, a.status from article a inner join article_tag_relation b on a.id=b.article_id "
	queryArticlesSQL        = "select id, title, status from article "
	addArticleSQL           = "insert into article(title)values($1)"
	updateArticleSQL        = "update article set %s where id=$1"
	deleteArticleByIDSQL    = "delete from article where id=$1"
)

func queryArticleWithTags(article ArticleWithTag, lastID int) []ArticleWithTag {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if article.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(article.ID)
	}

	if article.TagArticleID > 0 {
		whereSQL += " and b.tag_article_id=" + intToSafeString(article.TagArticleID)
	}

	if lastID >= 0 {
		whereSQL += " and a.id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryArticleWithTagsSQL + whereSQL)
	defer rows.Close()
	var articleWithTags []ArticleWithTag
	if rows == nil {
		return articleWithTags
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Status)
		articleWithTags = append(articleWithTags, temp)
	}

	return articleWithTags
}

func queryArticles(article Article, lastID int) []ArticleWithTag {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if article.ID > 0 {
		whereSQL += " and id=" + intToSafeString(article.ID)
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryArticlesSQL + whereSQL)
	defer rows.Close()
	var articles []ArticleWithTag
	if rows == nil {
		return articles
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Status)
		articles = append(articles, temp)
	}

	return articles
}

func addArticle(article Article) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(article.Title)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateArticle(article Article) int {

	hasCondition := false

	updateFieldSQL := ""

	if article.Title != "" {
		hasCondition = true
		updateFieldSQL += " title='" + strToSafeString(article.Title) + "' "
	}

	if article.Status > 0 {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " status=" + intToSafeString(article.Status)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateArticleSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(article.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteArticleByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleByIDSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
