package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

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
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(introduction.Summary, introduction.Content)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateDetailIntroduction(introduction DetailIntroduction) int {

	hasCondition := false

	updateFieldSQL := ""

	if introduction.Summary != "" {
		hasCondition = true
		updateFieldSQL += " summary='" + strToSafeString(introduction.Summary) + "' "
	}

	if introduction.Content != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " content='" + strToSafeString(introduction.Content) + "' "
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateDetailIntroductionSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(introduction.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteDetailIntroductions(introduction DetailIntroduction) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if introduction.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(introduction.ID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteDetailIntroductionSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
