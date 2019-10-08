package service

//QueryUserRoles 查询用户角色列表
func QueryUserRoles(userRole UserRole, lastID int) []UserRole {
	return queryUserRoles(userRole, lastID)
}

//SaveUserRole 保存用户角色
func SaveUserRole(userRole UserRole) int {
	if userRole.ID > 0 {
		return updateUserRole(userRole)
	}

	return addUserRole(userRole)
}

//DeleteUserRoles 删除用户角色
func DeleteUserRoles(userRole UserRole) int {
	return deleteUserRoles(userRole)
}
