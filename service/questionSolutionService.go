package service

//QueryQuestionSolutions 查询问题解决方式
func QueryQuestionSolutions(solution QuestionSolution, lastID int) []QuestionSolution {
	if lastID < 0 {
		lastID = 0
	}
	return queryQuestionSolutions(solution, lastID)
}

//QueryQuestionSolutionByID 通过ID查询问题解决方式
func QueryQuestionSolutionByID(id int) QuestionSolution {
	var queryQuestionSolution QuestionSolution
	queryQuestionSolution.ID = id
	var result QuestionSolution
	questionSolutions := queryQuestionSolutions(queryQuestionSolution, 0)
	if len(questionSolutions) > 0 {
		result = questionSolutions[0]
	}
	return result
}

//SaveQuestionSolution 保存问题解决方式
func SaveQuestionSolution(solution QuestionSolution) int {
	if solution.ID > 0 {
		return updateQuestionSolution(solution)
	}
	return addQuestionSolution(solution)
}

//DeleteQuestionSolutions 删除问题解决方式
func DeleteQuestionSolutions(solution QuestionSolution) int {
	return deleteQuestionSolutions(solution)
}
