package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

//Fragment 碎片
type Fragment struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type FragmentWithTag struct {
	Fragment
	TagFragmentID int `json:"tagFragmentId"`
}

const (
	queryFragmentWithTagsSQL = "select a.id, a.title, a.content from fragment a inner join fragment_tag_relation b where a.id=b.fragment_id "
	queryFragmentsSQL        = "select id, title, content from fragment "
	addFragmentSQL           = "insert into fragment(title, content)values($1, $2)"
	updateFragmentSQL        = "update fragment set %s where id=$1"
	deleteFragmentByIDSQL    = "delete from fragment where id=$1"
)

func queryFragmentWithTags(fragment FragmentWithTag, lastID int) []FragmentWithTag {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if fragment.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(fragment.ID)
	}

	if fragment.TagFragmentID > 0 {
		whereSQL += " and b.tag_fragment_id=" + intToSafeString(fragment.TagFragmentID)
	}

	if lastID >= 0 {
		whereSQL += " and a.id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryFragmentWithTagsSQL + whereSQL)
	defer rows.Close()
	var fragmentWithTags []FragmentWithTag
	if rows == nil {
		return fragmentWithTags
	}
	if err != nil {
		panic(err)
	}

	var temp FragmentWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Content)
		fragmentWithTags = append(fragmentWithTags, temp)
	}

	return fragmentWithTags
}

func queryFragments(fragment Fragment, lastID int) []FragmentWithTag {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if fragment.ID > 0 {
		whereSQL += " and id=" + intToSafeString(fragment.ID)
	}

	if fragment.Title != "" {
		whereSQL += " and title like '" + strToSafeString(fragment.Title) + "%' "
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryFragmentsSQL + whereSQL)
	var fragments []FragmentWithTag
	if rows == nil {
		return fragments
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp FragmentWithTag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Content)
		fragments = append(fragments, temp)
	}

	return fragments
}

func addFragment(fragment Fragment) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addFragmentSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(fragment.Title, fragment.Content)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateFragment(fragment Fragment) int {
	hasCondition := false

	updateFieldSQL := ""

	if fragment.Title != "" {
		hasCondition = true
		updateFieldSQL += " title='" + strToSafeString(fragment.Title) + "' "
	}

	if fragment.Content != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " content='" + strToSafeString(fragment.Content) + "' "
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateFragmentSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(fragment.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteFragmentByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteFragmentByIDSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
