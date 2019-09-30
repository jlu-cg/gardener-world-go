package service

//QueryQuestionSolutionRelationDetails 查询关联答案
func QueryQuestionSolutionRelationDetails(detail QuestionSolutionRelationDetail) []QuestionSolutionRelationDetail {
	return queryQuestionSolutionRelationDetails(detail)
}

//SaveQuestionSolutionRelation 保存问题和答案的关系
func SaveQuestionSolutionRelation(relation QuestionSolutionRelation) int {
	return addQuestionSolutionRelation(relation)
}

//DeleteQuestionSolutionRelations 删除问题和答案的关系
func DeleteQuestionSolutionRelations(relation QuestionSolutionRelation) int {
	return deleteQuestionSolutionRelations(relation)
}
