package service

//AddFragmentTagRelation 添加碎片和标签的关系
func AddFragmentTagRelation(relation FragmentTagRelation) int {
	return addFragmentTagRelation(relation)
}

//UpdateFragmentTagRelation 更新碎片和标签的关系
func UpdateFragmentTagRelation(relation FragmentTagRelation) int {
	return updateFragmentTagRelation(relation)
}

//DeleteFragmentTagRelationByID 删除碎片和标签的关系
func DeleteFragmentTagRelationByID(id int) int {
	return deleteFragmentTagRelationByID(id)
}
