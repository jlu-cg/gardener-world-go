package service

import "github.com/gardener/gardener-world-go/config"

//TagArticleFragmentRelation 标签间关系
type TagArticleFragmentRelation struct {
	ID            int `json:"id"`
	TagArticleID  int `json:"tagArticleID"`
	TagFragmentID int `json:"tagFragmentID"`
}

type TagArticleFragmentRelationDetail struct {
	TagArticleFragmentRelation
	TagName string `json:"tagName"`
	Type    int    `json:"type"`
}

const (
	queryTagArticleFragmentRelationDetailsSQL = "select a.id, a.tag_article_id, a.tag_fragment_id, b.name, b.type from tag_article_fragment_relation a inner join tag_fragment b on a.tag_fragment_id=b.id "
	addTagArticleFragmentRelationSQL          = "insert into tag_article_fragment_relation(tag_article_id, tag_fragment_id) values($1, $2)"
	deleteTagArticleFragmentRelatiosSQL       = "delete from tag_article_fragment_relation"
)

func queryTagArticleFragmentRelationDetails(relation TagArticleFragmentRelationDetail) []TagArticleFragmentRelationDetail {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(relation.ID)
	}
	if relation.TagArticleID > 0 {
		whereSQL += " and a.tag_article_id=" + intToSafeString(relation.TagArticleID)
	}
	if relation.TagFragmentID > 0 {
		whereSQL += " and a.tag_fragment_id=" + intToSafeString(relation.TagFragmentID)
	}

	rows, err := connection.Query(queryTagArticleFragmentRelationDetailsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var tagArticleFragmentRelationDetails []TagArticleFragmentRelationDetail
	if rows == nil {
		return tagArticleFragmentRelationDetails
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp TagArticleFragmentRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.TagArticleID, &temp.TagFragmentID, &temp.TagName, &temp.Type)
		tagArticleFragmentRelationDetails = append(tagArticleFragmentRelationDetails, temp)
	}

	return tagArticleFragmentRelationDetails
}

func addTagArticleFragmentRelation(relation TagArticleFragmentRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagArticleFragmentRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.TagArticleID, relation.TagFragmentID)
	if err != nil {
		return config.DBErrorExecution

	}
	return config.DBSuccess
}

func deleteTagArticleFragmentRelations(relation TagArticleFragmentRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}

	if relation.TagArticleID > 0 {
		whereSQL += " and tag_article_id=" + intToSafeString(relation.TagArticleID)
		hasCondition = true
	}

	if relation.TagFragmentID > 0 {
		whereSQL += " and tag_fragment_id=" + intToSafeString(relation.TagFragmentID)
		hasCondition = true
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagArticleFragmentRelatiosSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}
