package service

type QuestionSolutionRelation struct {
	ID         int `json:"id"`
	QuestionID int `json:"questionId"`
	SolutionID int `json:"solutionId"`
}

type QuestionSolutionRelationDetail struct {
	QuestionSolutionRelation
	Content      string `json:"content"`
	SolutionType int    `json:"solutionType"`
}

const (
	queryQuestionSolutionRelationDetailsSQL = "select a.id, a.question_id, a.solution_id, b.content, b.solution_type from question_solution_relation a inner join question_solution b on a.solutionId=b.id "
	addQuestionSolutionRelationSQL          = "insert into question_solution_relation(question_id, solution_id)values($1, $2)"
	deleteQuestionSolutionRelationSQL       = "delete from question_solution_relation where id=$1"
)

func queryQuestionSolutionRelationDetails(detail QuestionSolutionRelationDetail) []QuestionSolutionRelationDetail {

	hasCondition := false
	whereSQL := " where 1=1 "
	if detail.ID > 0 {
		whereSQL += " and a.question_id=" + intToSafeString(detail.ID)
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
		rows.Scan(&temp.ID, &temp.QuestionID, &temp.SolutionID, &temp.Content, &temp.SolutionType)
		details = append(details, temp)
	}

	return details
}

func addQuestionSolutionRelation(relation QuestionSolutionRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addQuestionSolutionRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(relation.QuestionID, relation.SolutionID)
	if err != nil {
		return -1
	}
	return 0
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
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteQuestionSolutionRelationSQL + whereSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
