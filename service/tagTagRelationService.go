package service

//AddTagTagRelation 添加标签间关联
func AddTagTagRelation(relation TagTagRelation) int {
	return addTagTagRelation(relation)
}

//UpdateTagTagRelation 更新标签间关联
func UpdateTagTagRelation(relation TagTagRelation) int {
	return updateTagTagRelation(relation)
}

//DeleteTagTagRelationByID 删除标签间关联
func DeleteTagTagRelationByID(id int) int {
	return deleteTagTagRelationByID(id)
}
