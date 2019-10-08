package service

import (
	"fmt"

	"github.com/gardener/gardener-world-go/config"
)

//UserRole 用户角色
type UserRole struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

const (
	queryUserRolesSQL  = "select id, name, status from user_role "
	addUserRoleSQL     = "insert into user_role(name, status)values($1, $2)"
	updateUserRoleSQL  = "update user_role set %s where id=$1"
	deleteUserRolesSQL = "delete from user_role "
)

func queryUserRoles(userRole UserRole, lastID int) []UserRole {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if userRole.ID > 0 {
		whereSQL += " and id=" + intToSafeString(userRole.ID)
	}

	if userRole.Name != "" {
		whereSQL += " and name like '" + strToSafeString(userRole.Name) + "%' "
	}

	if userRole.Status > 0 {
		whereSQL += " and status=" + intToSafeString(userRole.Status)
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryUserRolesSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var userRoles []UserRole
	if rows == nil {
		return userRoles
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp UserRole
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.Status)
		userRoles = append(userRoles, temp)
	}

	return userRoles
}

func addUserRole(userRole UserRole) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserRoleSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userRole.Name, userRole.Status)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func updateUserRole(userRole UserRole) int {

	hasCondition := false

	updateFieldSQL := ""

	if userRole.Name != "" {
		hasCondition = true
		updateFieldSQL += " name='" + strToSafeString(userRole.Name) + "' "
	}

	if userRole.Status > 0 {
		hasCondition = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " status=" + intToSafeString(userRole.Status)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateUserRoleSQL, updateFieldSQL))
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userRole.ID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteUserRoles(userRole UserRole) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if userRole.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(userRole.ID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteUserRolesSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
