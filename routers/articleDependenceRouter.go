package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticleDependence(app *iris.Application, crs context.Handler) {
	articleDependenceV1 := app.Party("/article/dependence/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleDependenceV1.Post("/list", func(ctx iris.Context) {
			var queryArticleDependenceDetail service.ArticleDependenceDetail
			ctx.ReadJSON(&queryArticleDependenceDetail)
			details := service.GetArticleDependenceDetails(queryArticleDependenceDetail)
			ctx.JSON(details)
		})

		articleDependenceV1.Post("/save", func(ctx iris.Context) {
			var addArticleDependence service.ArticleDependenceDetail
			ctx.ReadJSON(&addArticleDependence)
			code := service.AddArticleDependence(addArticleDependence)
			ctx.JSON(code)
		})

		articleDependenceV1.Post("/saveOrder", func(ctx iris.Context) {
			var addArticleDependences []service.ArticleDependenceDetail
			ctx.ReadJSON(&addArticleDependences)
			code := service.UpdateArticleDependences(addArticleDependences)
			ctx.JSON(code)
		})

		articleDependenceV1.Get("/delete", func(ctx iris.Context) {
			articleDependenceID := getIntVal("articleDependenceId", 0, ctx)
			result := service.DeleteArticleDependenceByID(articleDependenceID)
			ctx.JSON(result)
		})
	}
}
