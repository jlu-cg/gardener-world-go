package service

//Fragment 碎片
type Fragment struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

const (
	queryFragmentsSQL     = "select id, title, content from fragment "
	addFragmentSQL        = "insert into fragment(title, content)values($1, $2)"
	updateFragmentSQL     = "update fragment set title=$1, content=$2 where id=$3"
	deleteFragmentByIDSQL = "delete from fragment where id=$1"
)

func queryFragments(fragment Fragment, lastID int) []Fragment {
	connection := connect()
	defer release(connection)

	whereSQL := " where "

	if fragment.ID > 0 {
		whereSQL += " id=" + intToSafeString(fragment.ID) + " and "
	}

	if fragment.Title != "" {
		whereSQL += " title like '" + strToSafeString(fragment.Title) + "%' and "
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
		return -1
	}
	_, err = stmt.Exec(fragment.Title, fragment.Content)
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
	_, err = stmt.Exec(fragment.Title, fragment.Content, fragment.ID)
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
