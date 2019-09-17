package service

//QueryTagAttributes 查询属性
func QueryTagAttributes(attribute TagAttribute) []TagAttribute {

	return queryTagAttributes(attribute)
}

//AddTagAttribute 添加属性
func AddTagAttribute(attribute TagAttribute) int {
	return addTagAttribute(attribute)
}

//UpdateTagAttribute 更新属性
func UpdateTagAttribute(attribute TagAttribute) int {
	return updateTagAttribute(attribute)
}

//DeleteTagAttributeByID 删除属性
func DeleteTagAttributeByID(id int) int {
	return deleteTagAttributeByID(id)
}
