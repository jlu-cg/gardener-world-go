package service

import "github.com/gardener/gardener-world-go/config"

//TagFragment 标签
type TagFragment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

const (
	queryTagFragmentsSQL     = "select id, name, type from tag_fragment "
	addTagFragmentSQL        = "insert into tag_fragment(name, type)values($1, $2)"
	updateTagFragmentSQL     = "update tag_fragment set name=$1, type=$2 where id=$3"
	deleteTagFragmentByIDSQL = "delete from tag_fragment where id=$1"
)

func queryTagFragments(tagFragment TagFragment, lastID int) []TagFragment {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if tagFragment.Name != "" {
		whereSQL += " and name like '" + strToSafeString(tagFragment.Name) + "%' "
	}
	if tagFragment.ID > 0 {
		whereSQL += " and id=" + intToSafeString(tagFragment.ID)
	}
	if tagFragment.Type > 0 {
		whereSQL += " and type=" + intToSafeString(tagFragment.Type)
	}
	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryTagFragmentsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var tagFragments []TagFragment
	if rows == nil {
		return tagFragments
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp TagFragment
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.Type)
		tagFragments = append(tagFragments, temp)
	}

	return tagFragments
}

//保存标签
func addTagFragment(tagFragment TagFragment) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagFragmentSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(tagFragment.Name, tagFragment.Type)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

//更新标签
func updateTagFragment(tagFragment TagFragment) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagFragmentSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(tagFragment.Name, tagFragment.Type)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteTagFragmentByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagFragmentByIDSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
