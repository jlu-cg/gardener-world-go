package service

import "fmt"

//QuestionSolution 问题解决方法
type QuestionSolution struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	SolutionType int    `json:"solutionType"`
}

const (
	queryQuestionSolutionsSQL = "select id, content, solution_type from question_solution "
	addQuestionSolutionSQL    = "insert into question_solution(content, solution_type)values($1, $2)"
	updateQuestionSolutionSQL = "update question_solution set %s where id=$1"
	deleteQuestionSolutionSQL = "delete from question_solution "
)

func queryQuestionSolutions(solution QuestionSolution, lastID int) []QuestionSolution {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if solution.SolutionType > 0 {
		whereSQL += " and solution_type=" + intToSafeString(solution.SolutionType)
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
		rows.Scan(&temp.ID, &temp.Content, &temp.SolutionType)
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
	_, err = stmt.Exec(solution.Content, solution.SolutionType)
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

	if solution.SolutionType > 0 {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " solution_type=" + intToSafeString(solution.SolutionType)
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

	stmt, err := connection.Prepare(deleteQuestionSolutionSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
