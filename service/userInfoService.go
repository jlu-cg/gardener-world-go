package service

//QueryUserInfos 查询用户信息
func QueryUserInfos(userInfo UserInfo, lastID int) []UserInfo {
	return queryUserInfos(userInfo, lastID)
}

//QueryUserInfoByID 通过ID查询用户信息
func QueryUserInfoByID(id int) UserInfo {
	var queryUserInfo UserInfo
	queryUserInfo.ID = id
	var result UserInfo
	userInfos := queryUserInfos(queryUserInfo, 0)
	if len(userInfos) > 0 {
		result = userInfos[0]
	}
	return result
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
