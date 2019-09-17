package service

//FragmentTagRelation 碎片标签
type FragmentTagRelation struct {
	ID         int `json:"id"`
	FragmentID int `json:"fragmentId"`
	TagID      int `json:"tagId"`
}

const (
	queryFragmentTagRelationsSQL = "select id, fragment_id, tag_id from fragment_tag_relation "
	addFragmentTagRelationSQL    = "insert into fragment_tag_relation(fragment_id, tag_id) values($1, $2)"
	updateFragmentTagRelationSQL = "update fragment_tag_relation fragment_id=$1, tag_id=$2 where id=$3"
	deleteFragmentTagRelationSQL = "delete from fragment_tag_relation where id=$1"
)

func addFragmentTagRelation(relation FragmentTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addFragmentTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.FragmentID, relation.TagID)
	if err != nil {
		return -1

	}
	return 0
}

func updateFragmentTagRelation(relation FragmentTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateFragmentTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.FragmentID, relation.TagID, relation.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteFragmentTagRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteFragmentTagRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
