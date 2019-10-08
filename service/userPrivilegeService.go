package service

//QueryUserPrivileges 查询用户权限列表
func QueryUserPrivileges(userPrivilege UserPrivilege, lastID int) []UserPrivilege {
	return queryUserPrivileges(userPrivilege, lastID)
}

//SaveUserPrivilege 保存用户权限
func SaveUserPrivilege(userPrivilege UserPrivilege) int {
	if userPrivilege.ID > 0 {
		return updateUserPrivilege(userPrivilege)
	}

	return addUserPrivilege(userPrivilege)
}

//DeleteUserPrivileges 删除用户权限
func DeleteUserPrivileges(userPrivilege UserPrivilege) int {
	return deleteUserPrivileges(userPrivilege)
}
