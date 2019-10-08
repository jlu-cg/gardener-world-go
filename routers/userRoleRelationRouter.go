package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initUserRoleRelation(app *iris.Application, crs context.Handler) {
	userRoleRelationV1 := app.Party("/user/role/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		userRoleRelationV1.Post("/list", func(ctx iris.Context) {
			var queryUserRoleRelationWithRole service.UserRoleRelationWithRole
			ctx.ReadJSON(&queryUserRoleRelationWithRole)
			lastID := postIntVal("lastId", 0, ctx)
			userInfos := service.QueryUserRoleRelationWithRoles(queryUserRoleRelationWithRole, lastID)
			ctx.JSON(userInfos)
		})

		userRoleRelationV1.Post("/save", func(ctx iris.Context) {
			var addUserRoleRelation service.UserRoleRelation
			ctx.ReadJSON(&addUserRoleRelation)
			code := service.SaveUserRoleRelation(addUserRoleRelation)
			ctx.JSON(code)
		})

		userRoleRelationV1.Get("/delete", func(ctx iris.Context) {
			userRoleRelationID := getIntVal("userRoleRelationId", 0, ctx)
			var deleteUserRoleRelation service.UserRoleRelation
			deleteUserRoleRelation.ID = userRoleRelationID
			result := service.DeleteUserRoleRelations(deleteUserRoleRelation)
			ctx.JSON(result)
		})
	}
}
