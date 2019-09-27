package service

import "fmt"

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
		return -1
	}
	_, err = stmt.Exec(question.Summary, question.Detail)
	if err != nil {
		return -1
	}
	return 0
}

func updateQuestion(question Question) int {

	hasUpdate := false

	updateFieldSQL := ""

	if question.Summary != "" {
		hasUpdate = true
		updateFieldSQL += " summary='" + strToSafeString(question.Summary) + "' "
	}

	if question.Detail != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " detail='" + strToSafeString(question.Detail) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateQuestionSQL, updateFieldSQL))
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(question.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteQuestions(question Question) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if question.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(question.ID)
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteQuestionSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
