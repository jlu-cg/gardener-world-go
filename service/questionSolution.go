package service

import "fmt"

//QuestionSolution 问题解决方法
type QuestionSolution struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

const (
	queryQuestionSolutionsSQL = "select id, content, summary from question_solution "
	addQuestionSolutionSQL    = "insert into question_solution(content, summary)values($1, $2)"
	updateQuestionSolutionSQL = "update question_solution set %s where id=$1"
	deleteQuestionSolutionSQL = "delete from question_solution "
)

func queryQuestionSolutions(solution QuestionSolution, lastID int) []QuestionSolution {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if solution.Summary != "" {
		whereSQL += " and summary='" + strToSafeString(solution.Summary) + "' "
	}
	if solution.ID > 0 {
		whereSQL += " and id=" + intToSafeString(solution.ID)
	}
	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryQuestionSolutionsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var solutions []QuestionSolution
	if rows == nil {
		return solutions
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp QuestionSolution
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Content, &temp.Summary)
		solutions = append(solutions, temp)
	}

	return solutions
}

func addQuestionSolution(solution QuestionSolution) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addQuestionSolutionSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(solution.Content, solution.Summary)
	if err != nil {
		return -1
	}
	return 0
}

func updateQuestionSolution(solution QuestionSolution) int {

	hasUpdate := false

	updateFieldSQL := ""

	if solution.Content != "" {
		hasUpdate = true
		updateFieldSQL += " content='" + strToSafeString(solution.Content) + "' "
	}

	if solution.Summary != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " summary='" + strToSafeString(solution.Summary) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateQuestionSolutionSQL, updateFieldSQL))
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(solution.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteQuestionSolutions(solution QuestionSolution) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if solution.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(solution.ID)
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteQuestionSolutionSQL + whereSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
