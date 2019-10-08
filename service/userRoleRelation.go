package service

import (
	"github.com/gardener/gardener-world-go/config"
)

//UserRoleRelation 用户角色关系
type UserRoleRelation struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	RoleID int `json:"roleId"`
}

type UserRoleRelationWithRole struct {
	UserRoleRelation
	RoleName string `json:"roleName"`
	Status   int    `json:"status"`
}

const (
	queryUserRoleRelationWithRolesSQL = "select a.id, a.user_id, a.role_id, b.name, b.status from user_role_relation a inner join user_role b on a.role_id=b.id "
	addUserRoleRelationSQL            = "insert into user_role_relation(user_id, role_id)values($1, $2)"
	deleteUserRoleRelationsSQL        = "delete from user_role_relation "
)

func queryUserRoleRelationWithRoles(userRoleRelationWithRole UserRoleRelationWithRole, lastID int) []UserRoleRelationWithRole {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if userRoleRelationWithRole.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(userRoleRelationWithRole.ID)
	}

	if userRoleRelationWithRole.UserID > 0 {
		whereSQL += " and a.user_id=" + intToSafeString(userRoleRelationWithRole.UserID)
	}

	if userRoleRelationWithRole.RoleName != "" {
		whereSQL += " and b.name like '" + strToSafeString(userRoleRelationWithRole.RoleName) + "%' "
	}

	if userRoleRelationWithRole.Status > 0 {
		whereSQL += " and b.status=" + intToSafeString(userRoleRelationWithRole.Status)
	}

	if lastID >= 0 {
		whereSQL += " and a.id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryUserRoleRelationWithRolesSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var userRoleRelationWithRoles []UserRoleRelationWithRole
	if rows == nil {
		return userRoleRelationWithRoles
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp UserRoleRelationWithRole
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.UserID, &temp.RoleID, &temp.RoleName, &temp.Status)
		userRoleRelationWithRoles = append(userRoleRelationWithRoles, temp)
	}

	return userRoleRelationWithRoles
}

func addUserRoleRelation(userRoleRelation UserRoleRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserRoleRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userRoleRelation.UserID, userRoleRelation.RoleID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteUserRoleRelations(userRoleRelation UserRoleRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if userRoleRelation.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(userRoleRelation.ID)
	}
	if userRoleRelation.UserID > 0 {
		hasCondition = true
		whereSQL += " and user_id=" + intToSafeString(userRoleRelation.UserID)
	}
	if userRoleRelation.RoleID > 0 {
		hasCondition = true
		whereSQL += " and role_id=" + intToSafeString(userRoleRelation.RoleID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteUserRoleRelationsSQL + whereSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
