package service

import "github.com/gardener/gardener-world-go/config"

//ArticleTagRelation 文章标签关系
type ArticleTagRelation struct {
	ID           int `json:"id"`
	ArticleID    int `json:"articleId"`
	TagArticleID int `json:"tagArticleId"`
}

type ArticleTagRelationWithArticle struct {
	ArticleTagRelation
	Title string `json:"title"`
}

type ArticleTagRelationWithTag struct {
	ArticleTagRelation
	Name string `json:"name"`
	Type int    `json:"type"`
}

const (
	queryArticleTagRelationWithTagsSQL     = "select a.id, a.article_id, a.tag_article_id, b.name, b.type from article_tag_relation a inner join tag_article b on a.tag_article_id=b.id "
	queryArticleTagRelationWithArticlesSQL = "select a.id, a.article_id, a.tag_article_id, b.title from article_tag_relation a inner join article b on a.article_id=b.id "
	addArticleTagRelationSQL               = "insert into article_tag_relation(article_id, tag_article_id) values($1, $2)"
	updateArticleTagRelationSQL            = "update article_tag_relation article_id=$1, tag_article_id=$2 where id=$3"
	deleteArticleTagRelationSQL            = "delete from article_tag_relation "
)

func queryArticleTagRelationWithTags(relation ArticleTagRelationWithTag) []ArticleTagRelationWithTag {

	var articleTagRelationWithTags []ArticleTagRelationWithTag
	whereSQL := " where 1=1 "
	if relation.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(relation.ArticleID)
	} else {
		return articleTagRelationWithTags
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleTagRelationWithTagsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleTagRelationWithTags
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleTagRelationWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.TagArticleID, &temp.Name, &temp.Type)
		articleTagRelationWithTags = append(articleTagRelationWithTags, temp)
	}

	return articleTagRelationWithTags
}

func queryArticleTagRelationWithArticles(relation ArticleTagRelationWithArticle) []ArticleTagRelationWithArticle {
	var articleTagRelationWithArticles []ArticleTagRelationWithArticle
	whereSQL := " where 1=1 "
	if relation.TagArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(relation.TagArticleID)
	} else {
		return articleTagRelationWithArticles
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleTagRelationWithArticlesSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleTagRelationWithArticles
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleTagRelationWithArticle
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.TagArticleID, &temp.Title)
		articleTagRelationWithArticles = append(articleTagRelationWithArticles, temp)
	}

	return articleTagRelationWithArticles
}

func addArticleTagRelation(relation ArticleTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.ArticleID, relation.TagArticleID)
	if err != nil {
		return config.DBErrorExecution

	}
	return config.DBSuccess
}

func updateArticleTagRelation(relation ArticleTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.ArticleID, relation.TagArticleID, relation.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteArticleTagRelations(relation ArticleTagRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "

	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}
	if relation.ArticleID > 0 {
		whereSQL += " and article_id=" + intToSafeString(relation.ArticleID)
		hasCondition = true
	}

	if relation.TagArticleID > 0 {
		whereSQL += " and tag_article_id=" + intToSafeString(relation.TagArticleID)
		hasCondition = true
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleTagRelationSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}
