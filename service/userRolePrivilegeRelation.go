package service

import (
	"github.com/gardener/gardener-world-go/config"
)

//UserRolePrivilegeRelation 用户角色权限关系
type UserRolePrivilegeRelation struct {
	ID          int `json:"id"`
	RoleID      int `json:"roleId"`
	PrivilegeID int `json:"privilegeId"`
}

type UserRolePrivilegeRelationWithPrivilege struct {
	UserRolePrivilegeRelation
	PrivilegeName string `json:"privilegeName"`
	Status        int    `json:"status"`
}

const (
	queryUserRolePrivilegeRelationWithPrivilegesSQL = "select a.id, a.role_id, a.privilege_id, b.name, b.status from user_role_privilege_relation a inner join user_privilege b on a.privilege_id=b.id "
	addUserRolePrivilegeRelationSQL                 = "insert into user_role_privilege_relation(user_id, role_id)values($1, $2)"
	deleteUserRolePrivilegeRelationsSQL             = "delete from user_role_privilege_relation "
)

func queryUserRolePrivilegeRelationWithPrivileges(userRolePrivilegeRelationWithPrivilege UserRolePrivilegeRelationWithPrivilege, lastID int) []UserRolePrivilegeRelationWithPrivilege {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if userRolePrivilegeRelationWithPrivilege.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(userRolePrivilegeRelationWithPrivilege.ID)
	}

	if userRolePrivilegeRelationWithPrivilege.RoleID > 0 {
		whereSQL += " and a.role_id=" + intToSafeString(userRolePrivilegeRelationWithPrivilege.RoleID)
	}

	if userRolePrivilegeRelationWithPrivilege.PrivilegeName != "" {
		whereSQL += " and b.name like '" + strToSafeString(userRolePrivilegeRelationWithPrivilege.PrivilegeName) + "%' "
	}

	if userRolePrivilegeRelationWithPrivilege.Status > 0 {
		whereSQL += " and b.status=" + intToSafeString(userRolePrivilegeRelationWithPrivilege.Status)
	}

	if lastID >= 0 {
		whereSQL += " and a.id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryUserRolePrivilegeRelationWithPrivilegesSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var userRolePrivilegeRelationWithPrivileges []UserRolePrivilegeRelationWithPrivilege
	if rows == nil {
		return userRolePrivilegeRelationWithPrivileges
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp UserRolePrivilegeRelationWithPrivilege
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.RoleID, &temp.PrivilegeID, &temp.PrivilegeName, &temp.Status)
		userRolePrivilegeRelationWithPrivileges = append(userRolePrivilegeRelationWithPrivileges, temp)
	}

	return userRolePrivilegeRelationWithPrivileges
}

func addUserRolePrivilegeRelation(userRolePrivilegeRelation UserRolePrivilegeRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addUserRolePrivilegeRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(userRolePrivilegeRelation.RoleID, userRolePrivilegeRelation.PrivilegeID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteUserRolePrivilegeRelations(userRolePrivilegeRelation UserRolePrivilegeRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if userRolePrivilegeRelation.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(userRolePrivilegeRelation.ID)
	}
	if userRolePrivilegeRelation.RoleID > 0 {
		hasCondition = true
		whereSQL += " and role_id=" + intToSafeString(userRolePrivilegeRelation.RoleID)
	}
	if userRolePrivilegeRelation.PrivilegeID > 0 {
		hasCondition = true
		whereSQL += " and privilege_id=" + intToSafeString(userRolePrivilegeRelation.PrivilegeID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteUserRolePrivilegeRelationsSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
