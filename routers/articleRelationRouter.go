package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticleRelation(app *iris.Application, crs context.Handler) {
	articleRelationV1 := app.Party("/article/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleRelationV1.Post("/list", func(ctx iris.Context) {
			var queryArticleRelationDetail service.ArticleRelationDetail
			ctx.ReadJSON(&queryArticleRelationDetail)
			details := service.GetArticleRelationDetails(queryArticleRelationDetail)
			ctx.JSON(details)
		})

		articleRelationV1.Post("/save", func(ctx iris.Context) {
			var addArticleRelation service.ArticleRelationDetail
			ctx.ReadJSON(&addArticleRelation)
			code := service.AddArticleRelation(addArticleRelation)
			ctx.JSON(code)
		})

		articleRelationV1.Post("/saveOrder", func(ctx iris.Context) {
			var addArticleRelations []service.ArticleRelationDetail
			ctx.ReadJSON(&addArticleRelations)
			code := service.UpdateArticleRelations(addArticleRelations)
			ctx.JSON(code)
		})

		articleRelationV1.Get("/delete", func(ctx iris.Context) {
			articleRelationID := getIntVal("articleRelationId", 0, ctx)
			result := service.DeleteArticleRelationByID(articleRelationID)
			ctx.JSON(result)
		})
	}
}
