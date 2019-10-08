package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

//Question 问题
type Question struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
	Detail  string `json:"detail"`
}

const (
	queryQuestionsSQL = "select id, summary, detail from question "
	addQuestionSQL    = "insert into question(summary, detail)values($1, $2)"
	updateQuestionSQL = "update question set %s where id=$1"
	deleteQuestionSQL = "delete from question "
)

func queryQuestions(question Question, lastID int) []Question {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if question.Summary != "" {
		whereSQL += " and summary like '" + strToSafeString(question.Summary) + "%' "
	}
	if question.ID > 0 {
		whereSQL += " and id=" + intToSafeString(question.ID)
	}
	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryQuestionsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var questions []Question
	if rows == nil {
		return questions
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp Question
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Summary, &temp.Detail)
		questions = append(questions, temp)
	}

	return questions
}

func addQuestion(question Question) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addQuestionSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(question.Summary, question.Detail)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateQuestion(question Question) int {

	hasCondition := false

	updateFieldSQL := ""

	if question.Summary != "" {
		hasCondition = true
		updateFieldSQL += " summary='" + strToSafeString(question.Summary) + "' "
	}

	if question.Detail != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " detail='" + strToSafeString(question.Detail) + "' "
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateQuestionSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(question.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteQuestions(question Question) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if question.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(question.ID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteQuestionSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
