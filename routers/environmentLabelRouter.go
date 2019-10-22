package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryEnvironmentLabel struct {
	service.EnvironmentLabel
	LastID int `json:"lastId"`
}

func initEnvironmentLabel(app *iris.Application, crs context.Handler) {
	environmentLabelV1 := app.Party("/environment/label/v1", crs).AllowMethods(iris.MethodOptions)
	{
		environmentLabelV1.Post("/list", func(ctx iris.Context) {
			var queryEnvironmentLabel queryEnvironmentLabel
			ctx.ReadJSON(&queryEnvironmentLabel)
			environmentLabels := service.QueryEnvironmentLabels(queryEnvironmentLabel.EnvironmentLabel, queryEnvironmentLabel.LastID)
			ctx.JSON(environmentLabels)
		})

		environmentLabelV1.Post("/save", func(ctx iris.Context) {
			var saveEnvironmentLabel service.EnvironmentLabel
			ctx.ReadJSON(&saveEnvironmentLabel)
			code := service.SaveEnvironmentLabel(saveEnvironmentLabel)
			ctx.JSON(code)
		})

		environmentLabelV1.Get("/detail", func(ctx iris.Context) {
			environmentLabelID := getIntVal("environmentLabelId", 0, ctx)
			result := service.QueryEnvironmentLabelByID(environmentLabelID)
			ctx.JSON(result)
		})

		environmentLabelV1.Get("/delete", func(ctx iris.Context) {
			environmentLabelID := getIntVal("environmentLabelId", 0, ctx)
			var deleteEnvironmentLabel service.EnvironmentLabel
			deleteEnvironmentLabel.ID = environmentLabelID
			result := service.DeleteEnvironmentLabels(deleteEnvironmentLabel)
			ctx.JSON(result)
		})
	}
}
