package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initEnvironmentLabel(app *iris.Application, crs context.Handler) {
	environmentLabelV1 := app.Party("/environment/label/v1", crs).AllowMethods(iris.MethodOptions)
	{
		environmentLabelV1.Post("/list", func(ctx iris.Context) {
			var queryEnvironmentLabel service.EnvironmentLabel
			ctx.ReadJSON(&queryEnvironmentLabel)
			lastID := postIntVal("lastId", 0, ctx)
			environmentLabels := service.QueryEnvironmentLabels(queryEnvironmentLabel, lastID)
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
			var queryEnvironmentLabel service.EnvironmentLabel
			queryEnvironmentLabel.ID = environmentLabelID
			var result service.EnvironmentLabel
			labels := service.QueryEnvironmentLabels(queryEnvironmentLabel, 0)
			if len(labels) > 0 {
				result = labels[0]
			}
			ctx.JSON(result)
		})

		environmentLabelV1.Get("/delete", func(ctx iris.Context) {
			environmentLabelID := getIntVal("environmentLabelId", 0, ctx)
			var deleteEnvironmentLabel service.EnvironmentLabel
			deleteEnvironmentLabel.ID = environmentLabelID
			result := service.DeleteEnvironmentLabel(deleteEnvironmentLabel)
			ctx.JSON(result)
		})
	}
}
