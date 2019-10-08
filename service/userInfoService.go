package service

//QueryUserInfos 查询用户信息
func QueryUserInfos(userInfo UserInfo, lastID int) []UserInfo {
	return queryUserInfos(userInfo, lastID)
}

//SaveUserInfo 保存用户信息
func SaveUserInfo(userInfo UserInfo) int {
	if userInfo.ID > 0 {
		return updateUserInfo(userInfo)
	}

	return addUserInfo(userInfo)
}

//DeleteUserInfos 删除用户信息
func DeleteUserInfos(userInfo UserInfo) int {
	return deleteUserInfos(userInfo)
}
