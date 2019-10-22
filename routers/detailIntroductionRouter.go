package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryDetailIntroduction struct {
	service.DetailIntroduction
	LastID int `json:"lastId"`
}

func initDetailIntroduction(app *iris.Application, crs context.Handler) {
	detailIntroductionV1 := app.Party("/detail/introduction/v1", crs).AllowMethods(iris.MethodOptions)
	{
		detailIntroductionV1.Post("/list", func(ctx iris.Context) {
			var queryDetailIntroduction queryDetailIntroduction
			ctx.ReadJSON(&queryDetailIntroduction)
			detailIntroductions := service.QueryDetailIntroductions(queryDetailIntroduction.DetailIntroduction, queryDetailIntroduction.LastID)
			ctx.JSON(detailIntroductions)
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
