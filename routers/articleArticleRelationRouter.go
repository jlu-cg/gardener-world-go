package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticleArticleRelation(app *iris.Application, crs context.Handler) {
	articleArticleRelationV1 := app.Party("/article/article/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleArticleRelationV1.Post("/list", func(ctx iris.Context) {
			var queryArticleDependenceDetail service.ArticleArticleRelationDetail
			ctx.ReadJSON(&queryArticleDependenceDetail)
			details := service.GetArticleDependenceDetails(queryArticleDependenceDetail)
			ctx.JSON(details)
		})

		articleArticleRelationV1.Post("/save", func(ctx iris.Context) {
			var addArticleArticleRelationDetail service.ArticleArticleRelationDetail
			ctx.ReadJSON(&addArticleArticleRelationDetail)
			code := service.AddArticleDependence(addArticleArticleRelationDetail)
			ctx.JSON(code)
		})

		articleArticleRelationV1.Post("/saveOrder", func(ctx iris.Context) {
			var articleArticleRelationDetails []service.ArticleArticleRelationDetail
			ctx.ReadJSON(&articleArticleRelationDetails)
			code := service.UpdateArticleDependences(articleArticleRelationDetails)
			ctx.JSON(code)
		})

		articleArticleRelationV1.Get("/delete", func(ctx iris.Context) {
			articleArticleRelationID := getIntVal("articleArticleRelationId", 0, ctx)
			result := service.DeleteArticleDependenceByID(articleArticleRelationID)
			ctx.JSON(result)
		})
	}
}
