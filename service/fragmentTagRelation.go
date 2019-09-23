package service

//FragmentTagRelation 碎片标签
type FragmentTagRelation struct {
	ID         int    `json:"id"`
	FragmentID int    `json:"fragmentId"`
	TagID      int    `json:"tagId"`
	Name       string `json:"name"`
	TagType    int    `json:"tagType"`
}

const (
	queryFragmentTagRelationsSQL      = "select a.id, a.fragment_id, a.tag_id, b.name, b.tag_type from fragment_tag_relation a inner join tag b on a.tag_id=b.id "
	addFragmentTagRelationSQL         = "insert into fragment_tag_relation(fragment_id, tag_id) values($1, $2)"
	updateFragmentTagRelationSQL      = "update fragment_tag_relation set fragment_id=$1, tag_id=$2 where id=$3"
	deleteFragmentTagRelationSQL      = "delete from fragment_tag_relation where id=$1"
	deleteFragmentTagRelationOtherSQL = "delete from fragment_tag_relation "
)

func queryFragmentTagRelations(relation FragmentTagRelation) []FragmentTagRelation {
	var fragmentTagRelations []FragmentTagRelation
	whereSQL := " where 1=1 "
	if relation.FragmentID > 0 {
		whereSQL += " and a.fragment_id=" + intToSafeString(relation.FragmentID)
	} else {
		return fragmentTagRelations
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryFragmentTagRelationsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return fragmentTagRelations
	}
	if err != nil {
		panic(err)
	}

	var temp FragmentTagRelation
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.FragmentID, &temp.TagID, &temp.Name, &temp.TagType)
		fragmentTagRelations = append(fragmentTagRelations, temp)
	}

	return fragmentTagRelations
}

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

func deleteFragmentTagRelations(relation FragmentTagRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.FragmentID > 0 {
		whereSQL += " and fragment_id=" + intToSafeString(relation.FragmentID)
		hasCondition = true
	}

	if relation.TagID > 0 {
		whereSQL += " and tag_id=" + intToSafeString(relation.TagID)
		hasCondition = true
	}

	if !hasCondition {
		return -1
	}

	connection := connect()
	defer release(connection)
	stmt, err := connection.Prepare(deleteFragmentTagRelationOtherSQL + whereSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}
	return 0
}
