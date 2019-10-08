package service

//QueryUserRoleRelationWithRoles 查询用户角色关系包含角色信息
func QueryUserRoleRelationWithRoles(userRoleRelationWithRole UserRoleRelationWithRole, lastID int) []UserRoleRelationWithRole {
	return queryUserRoleRelationWithRoles(userRoleRelationWithRole, lastID)
}

//SaveUserRole 保存用户角色关系
func SaveUserRoleRelation(userRoleRelation UserRoleRelation) int {
	return addUserRoleRelation(userRoleRelation)
}

//DeleteUserRoles 删除用户角色关系
func DeleteUserRoleRelations(userRoleRelation UserRoleRelation) int {
	return deleteUserRoleRelations(userRoleRelation)
}
