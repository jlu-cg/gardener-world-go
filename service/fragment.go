package service

//Fragment 碎片
type Fragment struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	TagID   int    `json:"tagId"`
}

const (
	queryFragmentsSQL     = "select id, title, summary, content, tag_id from fragment "
	addFragmentSQL        = "insert into fragment(title, summary, content, tag_id)values($1, $2, $3, $4)"
	updateFragmentSQL     = "update fragment title=$1, summary=$2, content=$3, tag_id=$4 where id=$5"
	deleteFragmentByIDSQL = "delete from fragment where id=$1"
)

func queryFragments(fragment Fragment, lastID int) []Fragment {
	connection := connect()
	defer release(connection)

	whereSQL := " where "

	if fragment.ID > 0 {
		whereSQL += " id=" + intToSafeString(fragment.ID) + " and "
	}

	if fragment.TagID > 0 {
		whereSQL += " tag_id=" + intToSafeString(fragment.TagID) + " and "
	}
	whereSQL += " id>" + intToSafeString(lastID) + " limit 20 "
	rows, err := connection.Query(queryFragmentsSQL + whereSQL)
	var fragments []Fragment
	if rows == nil {
		return fragments
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp Fragment
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Summary, &temp.Content, &temp.TagID)
		fragments = append(fragments, temp)
	}

	return fragments
}

func addFragment(fragment Fragment) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addFragmentSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(fragment.Title, fragment.Summary, fragment.Content, fragment.TagID)
	if err != nil {
		return -1
	}
	return 0
}

func updateFragment(fragment Fragment) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateFragmentSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(fragment.Title, fragment.Summary, fragment.Content, fragment.TagID, fragment.ID)
	if err != nil {
		return -1
	}
	return 0
}

func queryFragmentByID(fragmentID int) Fragment {
	var fragment Fragment
	fragment.ID = fragmentID
	fragments := queryFragments(fragment, 0)
	return fragments[0]
}

func deleteFragmentByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteFragmentByIDSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
