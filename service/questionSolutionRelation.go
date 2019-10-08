package service

import "github.com/gardener/gardener-world-go/config"

type QuestionSolutionRelation struct {
	ID           int `json:"id"`
	QuestionID   int `json:"questionId"`
	SolutionID   int `json:"solutionId"`
	SolutionType int `json:"solutionType"`
	Position     int `json:"position"`
}

type QuestionSolutionRelationDetail struct {
	QuestionSolutionRelation
	Summary string `json:"summary"`
}

const (
	queryQuestionSolutionRelationDetailsSQL = "select a.id, a.question_id, a.solution_id, a.solution_type, a.position, b.summary from question_solution_relation a inner join question_solution b on a.solution_id=b.id "
	addQuestionSolutionRelationSQL          = "insert into question_solution_relation(question_id, solution_id, solution_type, position)values($1, $2, $3, $4)"
	updateQuestionSolutionRelationPosSQL    = "update question_solution_relation set position=$1 where id=$2"
	deleteQuestionSolutionRelationSQL       = "delete from question_solution_relation "
)

func queryQuestionSolutionRelationDetails(detail QuestionSolutionRelationDetail) []QuestionSolutionRelationDetail {

	hasCondition := false
	whereSQL := " where 1=1 "
	if detail.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(detail.ID)
		hasCondition = true
	}

	if detail.QuestionID > 0 {
		whereSQL += " and a.question_id=" + intToSafeString(detail.QuestionID)
		hasCondition = true
	}

	var details []QuestionSolutionRelationDetail
	if !hasCondition {
		return details
	}

	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryQuestionSolutionRelationDetailsSQL + whereSQL)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	if rows == nil {
		return details
	}

	var temp QuestionSolutionRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.QuestionID, &temp.SolutionID, &temp.SolutionType, &temp.Position, &temp.Summary)
		details = append(details, temp)
	}

	return details
}

func addQuestionSolutionRelation(relation QuestionSolutionRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addQuestionSolutionRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(relation.QuestionID, relation.SolutionID, relation.SolutionType, relation.Position)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateQuestionSolutionRelations(details []QuestionSolutionRelationDetail) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateQuestionSolutionRelationPosSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	for _, detail := range details {
		_, err = stmt.Exec(detail.Position, detail.ID)
		if err != nil {
			return config.DBErrorExecution
		}
	}
	return config.DBSuccess
}

func deleteQuestionSolutionRelations(relation QuestionSolutionRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(relation.ID)
	}
	if relation.QuestionID > 0 {
		hasCondition = true
		whereSQL += " and question_id=" + intToSafeString(relation.QuestionID)
	}
	if relation.SolutionID > 0 {
		hasCondition = true
		whereSQL += " and solution_id=" + intToSafeString(relation.SolutionID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteQuestionSolutionRelationSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
