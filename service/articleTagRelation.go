package service

import "github.com/gardener/gardener-world-go/config"

//ArticleTagRelation 文章标签关系
type ArticleTagRelation struct {
	ID        int    `json:"id"`
	ArticleID int    `json:"articleId"`
	TagID     int    `json:"tagId"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	TagType   int    `json:"tagType"`
}

const (
	queryArticleTagRelationsSQL    = "select a.id, a.article_id, a.tag_id, b.title from article_tag_relation a inner join article b on a.article_id=b.id "
	queryTagArticleTagRelationsSQL = "select a.id, a.article_id, a.tag_id, b.name, b.tag_type from article_tag_relation a inner join tag b on a.tag_id=b.id "
	addArticleTagRelationSQL       = "insert into article_tag_relation(article_id, tag_id) values($1, $2)"
	updateArticleTagRelationSQL    = "update article_tag_relation article_id=$1, tag_id=$2 where id=$3"
	deleteArticleTagRelationSQL    = "delete from article_tag_relation "
)

func queryArticleTagRelations(relation ArticleTagRelation) []ArticleTagRelation {

	var articleTagRelations []ArticleTagRelation
	whereSQL := " where 1=1 "
	if relation.TagID > 0 {
		whereSQL += " and a.tag_id=" + intToSafeString(relation.TagID)
	} else {
		return articleTagRelations
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleTagRelationsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleTagRelations
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleTagRelation
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.TagID, &temp.Title)
		articleTagRelations = append(articleTagRelations, temp)
	}

	return articleTagRelations
}

func queryTagArticleTagRelations(relation ArticleTagRelation) []ArticleTagRelation {
	var articleTagRelations []ArticleTagRelation
	whereSQL := " where 1=1 "
	if relation.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(relation.ArticleID)
	} else {
		return articleTagRelations
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryTagArticleTagRelationsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleTagRelations
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleTagRelation
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.TagID, &temp.Name, &temp.TagType)
		articleTagRelations = append(articleTagRelations, temp)
	}

	return articleTagRelations
}

func addArticleTagRelation(relation ArticleTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.ArticleID, relation.TagID)
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

	_, err = stmt.Exec(relation.ArticleID, relation.TagID, relation.ID)
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

	if relation.TagID > 0 {
		whereSQL += " and tag_id=" + intToSafeString(relation.TagID)
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
