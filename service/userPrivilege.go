package service

import "fmt"

//UserPrivilege 标签
type UserPrivilege struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

const (
	queryUserPrivilegesSQL  = "select id, name, status from user_privilege "
	addUserPrivilegeSQL     = "insert into user_privilege(name, status)values($1, $2)"
	updateUserPrivilegeSQL  = "update user_privilege set %s where id=$1"
	deleteUserPrivilegesSQL = "delete from user_privilege "
)

func queryUserPrivileges(userPrivilege UserPrivilege, lastID int) []UserPrivilege {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if userPrivilege.ID > 0 {
		whereSQL += " and id=" + intToSafeString(userPrivilege.ID)
	}

	if userPrivilege.Name != "" {
		whereSQL += " and name like '" + strToSafeString(userPrivilege.Name) + "%' "
	}

	if userPrivilege.Status > 0 {
		whereSQL += " and status=" + intToSafeString(userPrivilege.Status)
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryUserPrivilegesSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var userPrivileges []UserPrivilege
	if rows == nil {
		return userPrivileges
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp UserPrivilege
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.Status)
		userPrivileges = append(userPrivileges, temp)
	}

	return userPrivileges
}

func addUserPrivilege(userPrivilege UserPrivilege) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserInfoSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(userPrivilege.Name, userPrivilege.Status)
	if err != nil {
		return -1
	}
	return 0
}

func updateUserPrivilege(userPrivilege UserPrivilege) int {

	hasUpdate := false

	updateFieldSQL := ""

	if userPrivilege.Name != "" {
		hasUpdate = true
		updateFieldSQL += " name='" + strToSafeString(userPrivilege.Name) + "' "
	}

	if userPrivilege.Status > 0 {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " status=" + intToSafeString(userPrivilege.Status)
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
	_, err = stmt.Exec(userPrivilege.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteUserPrivileges(userPrivilege UserPrivilege) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if userPrivilege.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(userPrivilege.ID)
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
