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
			var queryArticleArticleRelationDetail service.ArticleArticleRelationDetail
			ctx.ReadJSON(&queryArticleArticleRelationDetail)
			details := service.QueryArticleArticleRelationDetails(queryArticleArticleRelationDetail)
			ctx.JSON(details)
		})

		articleArticleRelationV1.Post("/save", func(ctx iris.Context) {
			var addArticleArticleRelationDetail service.ArticleArticleRelationDetail
			ctx.ReadJSON(&addArticleArticleRelationDetail)
			code := service.AddArticleArticleRelation(addArticleArticleRelationDetail)
			ctx.JSON(code)
		})

		articleArticleRelationV1.Post("/saveOrder", func(ctx iris.Context) {
			var articleArticleRelationDetails []service.ArticleArticleRelationDetail
			ctx.ReadJSON(&articleArticleRelationDetails)
			code := service.UpdateArticleArticleRelations(articleArticleRelationDetails)
			ctx.JSON(code)
		})

		articleArticleRelationV1.Get("/delete", func(ctx iris.Context) {
			articleArticleRelationID := getIntVal("articleArticleRelationId", 0, ctx)
			result := service.DeleteArticleArticleRelationByID(articleArticleRelationID)
			ctx.JSON(result)
		})
	}
}
