package service

//QueryTagTagRelations 查询标签关联关系
func QueryTagTagRelations(relation TagTagRelation) []TagTagRelation {
	return queryTagTagRelations(relation)
}

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
	var relation TagTagRelation
	relation.ID = id
	return deleteTagTagRelations(relation)
}

//DeleteTagTagRelations 删除标签间关联
func DeleteTagTagRelations(relation TagTagRelation) int {
	return deleteTagTagRelations(relation)
}
