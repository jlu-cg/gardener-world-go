package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

//Tag 用户信息
type UserInfo struct {
	ID          int    `json:"id"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobilePhone"`
	Status      int    `json:"status"`
}

const (
	queryUserInfosSQL  = "select id, user_name, password, nick_name, email, mobile_phone, status from user_info "
	addUserInfoSQL     = "insert into user_info(user_name, password, nick_name, email, mobile_phone, status)values($1, $2, $3, $4, $5, $6)"
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
		whereSQL += " and email='" + strToSafeString(userInfo.Email) + "' "
	}

	if userInfo.MobilePhone != "" {
		whereSQL += " and mobile_phone like '" + strToSafeString(userInfo.MobilePhone) + "%' "
	}

	if userInfo.Status > 0 {
		whereSQL += " and status=" + intToSafeString(userInfo.Status)
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
		rows.Scan(&temp.ID, &temp.UserName, &temp.Password, &temp.NickName, &temp.Email, &temp.MobilePhone, &temp.Status)
		userInfos = append(userInfos, temp)
	}

	return userInfos
}

func addUserInfo(userInfo UserInfo) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserInfoSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userInfo.UserName, userInfo.Password, userInfo.NickName, userInfo.Email, userInfo.MobilePhone, userInfo.Status)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateUserInfo(userInfo UserInfo) int {

	hasCondition := false

	updateFieldSQL := ""

	if userInfo.UserName != "" {
		hasCondition = true
		updateFieldSQL += " user_name='" + strToSafeString(userInfo.UserName) + "' "
	}

	if userInfo.Password != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " password='" + strToSafeString(userInfo.Password) + "' "
	}

	if userInfo.NickName != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " nick_name='" + strToSafeString(userInfo.NickName) + "' "
	}

	if userInfo.Email != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " email='" + strToSafeString(userInfo.Email) + "' "
	}

	if userInfo.MobilePhone != "" {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " mobile_phone='" + strToSafeString(userInfo.MobilePhone) + "' "
	}

	if userInfo.Status > 0 {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " status=" + intToSafeString(userInfo.Status)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateUserInfoSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userInfo.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteUserInfos(userInfo UserInfo) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if userInfo.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(userInfo.ID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteUserInfosSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
