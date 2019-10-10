package service

//QueryFragmentIntroductionRelationWithIntroductions 查询碎片关联详细介绍
func QueryFragmentIntroductionRelationWithIntroductions(relation FragmentIntroductionRelationWithIntroduction, lastID int) []FragmentIntroductionRelationWithIntroduction {
	return queryFragmentIntroductionRelationWithIntroductions(relation, lastID)
}

//SaveFragmentIntroductionRelation 保存碎片详情关联
func SaveFragmentIntroductionRelation(relation FragmentIntroductionRelation) int {
	return addFragmentIntroductionRelation(relation)
}

//DeleteFragmentIntroductionRelations 删除碎片详情关联
func DeleteFragmentIntroductionRelations(relation FragmentIntroductionRelation) int {
	return deleteFragmentIntroductionRelations(relation)
}

//DeleteFragmentIntroductionRelationByID 删除碎片详情关联
func DeleteFragmentIntroductionRelationByID(id int) int {
	var deleteRelation FragmentIntroductionRelation
	deleteRelation.ID = id
	return deleteFragmentIntroductionRelations(deleteRelation)
}
