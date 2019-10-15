package service

import "github.com/gardener/gardener-world-go/config"

//FragmentTagRelation 碎片标签
type FragmentTagRelation struct {
	ID            int    `json:"id"`
	FragmentID    int    `json:"fragmentId"`
	TagFragmentID int    `json:"tagFragmentId"`
	Name          string `json:"name"`
	TagType       int    `json:"tagType"`
}

type FragmentTagRelationWithTag struct {
	FragmentTagRelation
	Type int `json:"type"`
}

type FragmentTagRelationWithFragment struct {
	FragmentTagRelation
	Title string `json:"title"`
}

const (
	queryFragmentTagRelationWithTagsSQL      = "select a.id, a.fragment_id, a.tag_fragment_id, b.name, b.type from fragment_tag_relation a inner join tag_fragment b on a.tag_fragment_id=b.id "
	queryFragmentTagRelationWithFragmentsSQL = "select a.id, a.fragment_id, a.tag_fragment_id, b.title from fragment_tag_relation a inner join fragment b on a.fragment_id=b.id "
	addFragmentTagRelationSQL                = "insert into fragment_tag_relation(fragment_id, tag_fragment_id) values($1, $2)"
	updateFragmentTagRelationSQL             = "update fragment_tag_relation set fragment_id=$1, tag_fragment_id=$2 where id=$3"
	deleteFragmentTagRelationSQL             = "delete from fragment_tag_relation "
)

func queryFragmentTagRelationWithTags(relation FragmentTagRelationWithTag) []FragmentTagRelationWithTag {
	var fragmentTagRelationWithTags []FragmentTagRelationWithTag
	whereSQL := " where 1=1 "
	if relation.FragmentID > 0 {
		whereSQL += " and a.fragment_id=" + intToSafeString(relation.FragmentID)
	} else {
		return fragmentTagRelationWithTags
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryFragmentTagRelationWithTagsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return fragmentTagRelationWithTags
	}
	if err != nil {
		panic(err)
	}

	var temp FragmentTagRelationWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.FragmentID, &temp.TagFragmentID, &temp.Name, &temp.Type)
		fragmentTagRelationWithTags = append(fragmentTagRelationWithTags, temp)
	}

	return fragmentTagRelationWithTags
}

func queryFragmentTagRelationWithFragments(relation FragmentTagRelationWithFragment) []FragmentTagRelationWithFragment {
	var fragmentTagRelationWithFragments []FragmentTagRelationWithFragment
	whereSQL := " where 1=1 "
	if relation.TagFragmentID > 0 {
		whereSQL += " and a.tag_fragment_id=" + intToSafeString(relation.TagFragmentID)
	} else {
		return fragmentTagRelationWithFragments
	}

	whereSQL += " order by a.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryFragmentTagRelationWithFragmentsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return fragmentTagRelationWithFragments
	}
	if err != nil {
		panic(err)
	}

	var temp FragmentTagRelationWithFragment
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.FragmentID, &temp.TagFragmentID, &temp.Title)
		fragmentTagRelationWithFragments = append(fragmentTagRelationWithFragments, temp)
	}

	return fragmentTagRelationWithFragments
}

func addFragmentTagRelation(relation FragmentTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addFragmentTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.FragmentID, relation.TagFragmentID)
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
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(relation.FragmentID, relation.TagFragmentID, relation.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteFragmentTagRelations(relation FragmentTagRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}
	if relation.FragmentID > 0 {
		whereSQL += " and fragment_id=" + intToSafeString(relation.FragmentID)
		hasCondition = true
	}

	if relation.TagFragmentID > 0 {
		whereSQL += " and tag_id=" + intToSafeString(relation.TagFragmentID)
		hasCondition = true
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)
	stmt, err := connection.Prepare(deleteFragmentTagRelationSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}
