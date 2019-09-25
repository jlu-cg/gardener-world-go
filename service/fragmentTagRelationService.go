package service

//QueryFragmentTagRelations 查询碎片标签关系
func QueryFragmentTagRelations(relation FragmentTagRelation) []FragmentTagRelation {
	return queryFragmentTagRelations(relation)
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
