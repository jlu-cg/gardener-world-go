package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initUserPrivilege(app *iris.Application, crs context.Handler) {
	userPrivilegeV1 := app.Party("/user/privilege/v1", crs).AllowMethods(iris.MethodOptions)
	{
		userPrivilegeV1.Post("/list", func(ctx iris.Context) {
			var queryUserPrivilege service.UserPrivilege
			ctx.ReadJSON(&queryUserPrivilege)
			lastID := postIntVal("lastId", 0, ctx)
			userInfos := service.QueryUserPrivileges(queryUserPrivilege, lastID)
			ctx.JSON(userInfos)
		})

		userPrivilegeV1.Post("/save", func(ctx iris.Context) {
			var addUserPrivilege service.UserPrivilege
			ctx.ReadJSON(&addUserPrivilege)
			code := service.SaveUserPrivilege(addUserPrivilege)
			ctx.JSON(code)
		})

		userPrivilegeV1.Get("/detail", func(ctx iris.Context) {
			userPrivilegeID := getIntVal("userPrivilegeId", 0, ctx)
			result := service.QueryUserPrivilegeByID(userPrivilegeID)
			ctx.JSON(result)
		})

		userPrivilegeV1.Get("/delete", func(ctx iris.Context) {
			userPrivilegeID := getIntVal("userPrivilegeId", 0, ctx)
			var deleteUserPrivilege service.UserPrivilege
			deleteUserPrivilege.ID = userPrivilegeID
			result := service.DeleteUserPrivileges(deleteUserPrivilege)
			ctx.JSON(result)
		})
	}
}
