package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryUserInfo struct {
	service.UserInfo
	LastID int `json:"lastId"`
}

func initUserInfo(app *iris.Application, crs context.Handler) {
	userInfoV1 := app.Party("/user/info/v1", crs).AllowMethods(iris.MethodOptions)
	{
		userInfoV1.Post("/list", func(ctx iris.Context) {
			var queryUserInfo queryUserInfo
			ctx.ReadJSON(&queryUserInfo)
			userInfos := service.QueryUserInfos(queryUserInfo.UserInfo, queryUserInfo.LastID)
			ctx.JSON(userInfos)
		})

		userInfoV1.Post("/save", func(ctx iris.Context) {
			var addUserInfo service.UserInfo
			ctx.ReadJSON(&addUserInfo)
			code := service.SaveUserInfo(addUserInfo)
			ctx.JSON(code)
		})

		userInfoV1.Get("/detail", func(ctx iris.Context) {
			userInfoID := getIntVal("userInfoId", 0, ctx)
			result := service.QueryUserInfoByID(userInfoID)
			ctx.JSON(result)
		})

		userInfoV1.Get("/delete", func(ctx iris.Context) {
			userInfoID := getIntVal("userInfoId", 0, ctx)
			var deleteUserInfo service.UserInfo
			deleteUserInfo.ID = userInfoID
			result := service.DeleteUserInfos(deleteUserInfo)
			ctx.JSON(result)
		})
	}
}
