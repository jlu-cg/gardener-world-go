package service

//QueryQuestions 查询问题
func QueryQuestions(question Question, lastID int) []Question {
	if lastID < 0 {
		lastID = 0
	}

	return queryQuestions(question, lastID)
}

//QueryQuestionByID 通过ID查询问题
func QueryQuestionByID(id int) Question {
	var queryQuestion Question
	queryQuestion.ID = id
	var result Question
	questions := queryQuestions(queryQuestion, 0)
	if len(questions) > 0 {
		result = questions[0]
	}
	return result
}

//SaveQuestion 保存问题
func SaveQuestion(question Question) int {
	if question.ID > 0 {
		return updateQuestion(question)
	}
	return addQuestion(question)
}

//DeleteQuestions 删除问题
func DeleteQuestions(question Question) int {
	code := deleteQuestions(question)

	var deleteQuestionSolutionRelation QuestionSolutionRelation
	deleteQuestionSolutionRelation.QuestionID = question.ID
	code = DeleteQuestionSolutionRelations(deleteQuestionSolutionRelation)

	return code
}
