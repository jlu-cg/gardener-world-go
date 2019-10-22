package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryUserRole struct {
	service.UserRole
	LastID int `json:"lastId"`
}

func initUserRole(app *iris.Application, crs context.Handler) {
	userRoleV1 := app.Party("/user/role/v1", crs).AllowMethods(iris.MethodOptions)
	{
		userRoleV1.Post("/list", func(ctx iris.Context) {
			var queryUserRole queryUserRole
			ctx.ReadJSON(&queryUserRole)
			userInfos := service.QueryUserRoles(queryUserRole.UserRole, queryUserRole.LastID)
			ctx.JSON(userInfos)
		})

		userRoleV1.Post("/save", func(ctx iris.Context) {
			var addUserRole service.UserRole
			ctx.ReadJSON(&addUserRole)
			code := service.SaveUserRole(addUserRole)
			ctx.JSON(code)
		})

		userRoleV1.Get("/detail", func(ctx iris.Context) {
			userRoleID := getIntVal("userRoleId", 0, ctx)
			result := service.QueryUserRoleByID(userRoleID)
			ctx.JSON(result)
		})

		userRoleV1.Get("/delete", func(ctx iris.Context) {
			userRoleID := getIntVal("userRoleId", 0, ctx)
			var deleteUserRole service.UserRole
			deleteUserRole.ID = userRoleID
			result := service.DeleteUserRoles(deleteUserRole)
			ctx.JSON(result)
		})
	}
}
