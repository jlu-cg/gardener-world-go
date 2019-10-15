package service

//QueryDetailIntroductions 查询详细介绍
func QueryDetailIntroductions(introduction DetailIntroduction, lastID int) []DetailIntroduction {
	if lastID < 0 {
		lastID = 0
	}
	return queryDetailIntroductions(introduction, lastID)
}

//QueryDetailIntroductionByID 通过ID查询详细介绍
func QueryDetailIntroductionByID(id int) DetailIntroduction {
	var queryDetailIntroduction DetailIntroduction
	queryDetailIntroduction.ID = id
	var result DetailIntroduction
	introductions := queryDetailIntroductions(queryDetailIntroduction, 0)
	if len(introductions) > 0 {
		result = introductions[0]
	}
	return result
}

//SaveEnvironmentLabel 保存详细介绍
func SaveDetailIntroduction(introduction DetailIntroduction) int {
	if introduction.ID > 0 {
		return updateDetailIntroduction(introduction)
	}
	return addDetailIntroduction(introduction)
}

//DeleteDetailIntroductions 删除详细介绍
func DeleteDetailIntroductions(introduction DetailIntroduction) int {
	code := deleteDetailIntroductions(introduction)

	//删除碎片和详细介绍的对应关系
	var deleteFragmentIntroductionRelation FragmentIntroductionRelation
	deleteFragmentIntroductionRelation.DetailIntroductionID = introduction.ID
	code = DeleteFragmentIntroductionRelations(deleteFragmentIntroductionRelation)

	//删除详细介绍和环境标签对应关系
	var deleteIntroductionEnvironmentRelation IntroductionEnvironmentRelation
	deleteIntroductionEnvironmentRelation.DetailIntroductionID = introduction.ID
	code = DeleteIntroductionEnvironmentRelations(deleteIntroductionEnvironmentRelation)

	return code
}
