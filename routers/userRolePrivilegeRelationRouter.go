package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initUserRolePrivilegeRelation(app *iris.Application, crs context.Handler) {
	userRolePrivilegeRelationV1 := app.Party("/user/role/privilege/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		userRolePrivilegeRelationV1.Post("/list", func(ctx iris.Context) {
			var queryUserRolePrivilegeRelationWithPrivilege service.UserRolePrivilegeRelationWithPrivilege
			ctx.ReadJSON(&queryUserRolePrivilegeRelationWithPrivilege)
			lastID := postIntVal("lastId", 0, ctx)
			userInfos := service.QueryUserRolePrivilegeRelationWithPrivileges(queryUserRolePrivilegeRelationWithPrivilege, lastID)
			ctx.JSON(userInfos)
		})

		userRolePrivilegeRelationV1.Post("/save", func(ctx iris.Context) {
			var addUserRolePrivilegeRelation service.UserRolePrivilegeRelation
			ctx.ReadJSON(&addUserRolePrivilegeRelation)
			code := service.SaveUserRolePrivilegeRelation(addUserRolePrivilegeRelation)
			ctx.JSON(code)
		})

		userRolePrivilegeRelationV1.Get("/delete", func(ctx iris.Context) {
			userRolePrivilegeRelationID := getIntVal("userRolePrivilegeRelationId", 0, ctx)
			var deleteUserRolePrivilegeRelation service.UserRolePrivilegeRelation
			deleteUserRolePrivilegeRelation.ID = userRolePrivilegeRelationID
			result := service.DeleteUserRolePrivilegeRelations(deleteUserRolePrivilegeRelation)
			ctx.JSON(result)
		})
	}
}
