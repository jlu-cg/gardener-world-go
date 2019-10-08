package service

import "fmt"

//Tag 用户信息
type UserInfo struct {
	ID          int    `json:"id"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilePhone"`
}

const (
	queryUserInfosSQL  = "select id, user_name, password, nick_name, email, mobile_phone from user_info "
	addUserInfoSQL     = "insert into user_info(user_name, password, nick_name, email, mobile_phone)values($1, $2, $3, $4, $5)"
	updateUserInfoSQL  = "update user_info set %s where id=$1"
	deleteUserInfosSQL = "delete from user_info "
)

func queryUserInfos(userInfo UserInfo, lastID int) []UserInfo {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if userInfo.ID > 0 {
		whereSQL += " and id=" + intToSafeString(userInfo.ID)
	}

	if userInfo.UserName != "" {
		whereSQL += " and user_name like '" + strToSafeString(userInfo.UserName) + "%' "
	}

	if userInfo.Email != "" {
		whereSQL += " and email= '" + strToSafeString(userInfo.Email) + "' "
	}

	if userInfo.MobilePhone != "" {
		whereSQL += " and mobile_phone like '" + strToSafeString(userInfo.MobilePhone) + "%' "
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryUserInfosSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var userInfos []UserInfo
	if rows == nil {
		return userInfos
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp UserInfo
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.UserName, &temp.Password, &temp.NickName, &temp.Email, &temp.MobilePhone)
		userInfos = append(userInfos, temp)
	}

	return userInfos
}

func addUserInfo(userInfo UserInfo) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserInfoSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(userInfo.UserName, userInfo.Password, userInfo.NickName, userInfo.Email, userInfo.MobilePhone)
	if err != nil {
		return -1
	}
	return 0
}

func updateUserInfo(userInfo UserInfo) int {

	hasUpdate := false

	updateFieldSQL := ""

	if userInfo.UserName != "" {
		hasUpdate = true
		updateFieldSQL += " user_name='" + strToSafeString(userInfo.UserName) + "' "
	}

	if userInfo.Password != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " password='" + strToSafeString(userInfo.Password) + "' "
	}

	if userInfo.NickName != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " nick_name='" + strToSafeString(userInfo.NickName) + "' "
	}

	if userInfo.Email != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " email='" + strToSafeString(userInfo.Email) + "' "
	}

	if userInfo.MobilePhone != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " mobile_phone='" + strToSafeString(userInfo.MobilePhone) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateUserInfoSQL, updateFieldSQL))
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(userInfo.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteUserInfos(userInfo UserInfo) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if userInfo.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(userInfo.ID)
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteUserInfosSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
