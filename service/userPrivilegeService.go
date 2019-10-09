package service

//QueryUserPrivileges 查询用户权限列表
func QueryUserPrivileges(userPrivilege UserPrivilege, lastID int) []UserPrivilege {
	return queryUserPrivileges(userPrivilege, lastID)
}

//QueryUserPrivilegeByID 通过ID查询用户权限
func QueryUserPrivilegeByID(id int) UserPrivilege {
	var queryUserPrivilege UserPrivilege
	queryUserPrivilege.ID = id
	var result UserPrivilege
	userPrivileges := queryUserPrivileges(queryUserPrivilege, 0)
	if len(userPrivileges) > 0 {
		result = userPrivileges[0]
	}
	return result
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
