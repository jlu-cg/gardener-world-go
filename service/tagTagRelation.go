package service

import "github.com/gardener/gardener-world-go/config"

//TagTagRelation 标签间关系
type TagTagRelation struct {
	ID          int    `json:"id"`
	TagID       int    `json:"tagId"`
	RelateTagID int    `json:"relateTagId"`
	RelateType  int    `json:"relateType"`
	TagName     string `json:"tagName"`
}

const (
	queryTagTagRelationsSQL = "select a.id, a.tag_id, a.relate_tag_id, a.relate_type, b.name from tag_tag_relation a inner join tag b on a.relate_tag_id=b.id "
	addTagTagRelationSQL    = "insert into tag_tag_relation(tag_id, relate_tag_id, relate_type) values($1, $2, $3)"
	updateTagTagRelationSQL = "update tag_tag_relation set tag_id=$1, relate_tag_id=$2, relate_type=$3 where id=$4"
	deleteTagTagRelationSQL = "delete from tag_tag_relation"
)

func queryTagTagRelations(relation TagTagRelation) []TagTagRelation {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "
	if relation.TagID > 0 {
		whereSQL += " and a.tag_id=" + intToSafeString(relation.TagID)
	}
	if relation.RelateType > 0 {
		whereSQL += " and a.relate_type=" + intToSafeString(relation.RelateType)
	}

	rows, err := connection.Query(queryTagTagRelationsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var tagTagRelations []TagTagRelation
	if rows == nil {
		return tagTagRelations
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp TagTagRelation
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.TagID, &temp.RelateTagID, &temp.RelateType, &temp.TagName)
		tagTagRelations = append(tagTagRelations, temp)
	}

	return tagTagRelations
}

func addTagTagRelation(relation TagTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.TagID, relation.RelateTagID, relation.RelateType)
	if err != nil {
		return config.DBErrorExecution

	}
	return config.DBSuccess
}

func updateTagTagRelation(relation TagTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.TagID, relation.RelateTagID, relation.RelateType, relation.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteTagTagRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagTagRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteTagTagRelations(relation TagTagRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}

	if relation.TagID > 0 {
		whereSQL += " and tag_id=" + intToSafeString(relation.TagID)
		hasCondition = true
	}

	if relation.RelateTagID > 0 {
		whereSQL += " and relate_tag_id=" + intToSafeString(relation.RelateTagID)
		hasCondition = true
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagTagRelationSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}
