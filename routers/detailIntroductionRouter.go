package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initDetailIntroduction(app *iris.Application, crs context.Handler) {
	detailIntroductionV1 := app.Party("/detail/introduction/v1", crs).AllowMethods(iris.MethodOptions)
	{
		detailIntroductionV1.Post("/list", func(ctx iris.Context) {
			var queryDetailIntroduction service.DetailIntroduction
			ctx.ReadJSON(&queryDetailIntroduction)
			lastID := postIntVal("lastId", 0, ctx)
			environmentLabels := service.QueryDetailIntroductions(queryDetailIntroduction, lastID)
			ctx.JSON(environmentLabels)
		})

		detailIntroductionV1.Post("/save", func(ctx iris.Context) {
			var saveDetailIntroduction service.DetailIntroduction
			ctx.ReadJSON(&saveDetailIntroduction)
			code := service.SaveDetailIntroduction(saveDetailIntroduction)
			ctx.JSON(code)
		})

		detailIntroductionV1.Get("/detail", func(ctx iris.Context) {
			detailIntroductionID := getIntVal("detailIntroductionId", 0, ctx)
			result := service.QueryDetailIntroductionByID(detailIntroductionID)
			ctx.JSON(result)
		})

		detailIntroductionV1.Get("/delete", func(ctx iris.Context) {
			detailIntroductionID := getIntVal("detailIntroductionId", 0, ctx)
			var deleteDetailIntroduction service.DetailIntroduction
			deleteDetailIntroduction.ID = detailIntroductionID
			result := service.DeleteDetailIntroductions(deleteDetailIntroduction)
			ctx.JSON(result)
		})
	}
}
