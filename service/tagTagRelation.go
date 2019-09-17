package service

//TagTagRelation 标签间关系
type TagTagRelation struct {
	ID          int `json:"id"`
	TagID       int `json:"tagId"`
	RelateTagID int `json:"relateTagId"`
	RelateType  int `json:"relateType"`
}

const (
	queryTagTagRelationsSQL = "select id, tag_id, relate_tag_id, relate_type from tag_tag_relation "
	addTagTagRelationSQL    = "insert into tag_tag_relation(tag_id, relate_tag_id, relate_type) values($1, $2, $3)"
	updateTagTagRelationSQL = "update tag_tag_relation tag_id=$1, relate_tag_id=$2, relate_type=$3 where id=$4"
	deleteTagTagRelationSQL = "delete from tag_tag_relation where id=$1"
)

func addTagTagRelation(relation TagTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.TagID, relation.RelateTagID, relation.RelateType)
	if err != nil {
		return -1

	}
	return 0
}

func updateTagTagRelation(relation TagTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.TagID, relation.RelateTagID, relation.RelateType, relation.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteTagTagRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagTagRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
