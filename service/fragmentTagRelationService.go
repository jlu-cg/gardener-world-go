package service

//QueryFragmentTagRelationWithTags 查询碎片标签关系和标签
func QueryFragmentTagRelationWithTags(relation FragmentTagRelationWithTag) []FragmentTagRelationWithTag {
	return queryFragmentTagRelationWithTags(relation)
}

//QueryFragmentTagRelationWithFragments 查询碎片标签关系和碎片
func QueryFragmentTagRelationWithFragments(relation FragmentTagRelationWithFragment) []FragmentTagRelationWithFragment {
	return queryFragmentTagRelationWithFragments(relation)
}

//SaveFragmentTagRelation 保存碎片和标签的关系
func SaveFragmentTagRelation(relation FragmentTagRelation) int {
	if relation.ID > 0 {
		return updateFragmentTagRelation(relation)
	}
	return addFragmentTagRelation(relation)
}

//DeleteFragmentTagRelationByID 删除碎片和标签的关系
func DeleteFragmentTagRelationByID(id int) int {
	var relation FragmentTagRelation
	relation.ID = id
	return deleteFragmentTagRelations(relation)
}

//DeleteFragmentTagRelations 删除碎片和标签的关系
func DeleteFragmentTagRelations(relation FragmentTagRelation) int {
	return deleteFragmentTagRelations(relation)
}
