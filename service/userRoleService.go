package service

//QueryUserRoles 查询用户角色列表
func QueryUserRoles(userRole UserRole, lastID int) []UserRole {
	return queryUserRoles(userRole, lastID)
}

//QueryUserRoleByID 通过ID查询用户角色
func QueryUserRoleByID(id int) UserRole {
	var queryUserRole UserRole
	queryUserRole.ID = id
	var result UserRole
	userRoles := queryUserRoles(queryUserRole, 0)
	if len(userRoles) > 0 {
		result = userRoles[0]
	}
	return result
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
