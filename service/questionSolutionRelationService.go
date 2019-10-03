package service

//QueryQuestionSolutionRelationDetails 查询关联答案
func QueryQuestionSolutionRelationDetails(detail QuestionSolutionRelationDetail) []QuestionSolutionRelationDetail {
	return queryQuestionSolutionRelationDetails(detail)
}

//SaveQuestionSolutionRelation 保存问题和答案的关系
func SaveQuestionSolutionRelation(relation QuestionSolutionRelation) int {
	return addQuestionSolutionRelation(relation)
}

func UpdateQuestionSolutionRelations(details []QuestionSolutionRelationDetail) int {
	return updateQuestionSolutionRelations(details)
}

//DeleteQuestionSolutionRelations 删除问题和答案的关系
func DeleteQuestionSolutionRelations(relation QuestionSolutionRelation) int {
	return deleteQuestionSolutionRelations(relation)
}
