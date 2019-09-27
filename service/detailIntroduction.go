package service

import "fmt"

type DetailIntroduction struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

const (
	queryDetailIntroductionsSQL = "select id, summary, content from detail_introduction "
	addDetailIntroductionSQL    = "insert into detail_introduction(summary, content)values($1, $2)"
	updateDetailIntroductionSQL = "update detail_introduction set %s where id=$1"
	deleteDetailIntroductionSQL = "delete from detail_introduction "
)

func queryDetailIntroductions(introduction DetailIntroduction, lastID int) []DetailIntroduction {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if introduction.ID > 0 {
		whereSQL += " and id=" + intToSafeString(introduction.ID)
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryDetailIntroductionsSQL + whereSQL)
	defer rows.Close()
	var introductions []DetailIntroduction
	if rows == nil {
		return introductions
	}
	if err != nil {
		panic(err)
	}

	var temp DetailIntroduction
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Summary, &temp.Content)
		introductions = append(introductions, temp)
	}

	return introductions
}

func addDetailIntroduction(introduction DetailIntroduction) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addDetailIntroductionSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(introduction.Summary, introduction.Content)
	if err != nil {
		return -1
	}
	return 0
}

func updateDetailIntroduction(introduction DetailIntroduction) int {

	hasUpdate := false

	updateFieldSQL := ""

	if introduction.Summary != "" {
		hasUpdate = true
		updateFieldSQL += " summary='" + strToSafeString(introduction.Summary) + "' "
	}

	if introduction.Content != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " content='" + strToSafeString(introduction.Content) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateDetailIntroductionSQL, updateFieldSQL))
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(introduction.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteDetailIntroductions(introduction DetailIntroduction) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if introduction.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(introduction.ID)
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteDetailIntroductionSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
