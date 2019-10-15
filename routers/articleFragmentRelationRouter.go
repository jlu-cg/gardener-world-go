package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticleFragmentRelation(app *iris.Application, crs context.Handler) {
	articleFragmentRelationV1 := app.Party("/article/fragment/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleFragmentRelationV1.Post("/list", func(ctx iris.Context) {
			var queryArticleFragmentRelationDetail service.ArticleFragmentRelationDetail
			ctx.ReadJSON(&queryArticleFragmentRelationDetail)
			details := service.QueryArticleFragmentRelationDetails(queryArticleFragmentRelationDetail)
			ctx.JSON(details)
		})

		articleFragmentRelationV1.Post("/save", func(ctx iris.Context) {
			var addArticleFragmentRelation service.ArticleFragmentRelationDetail
			ctx.ReadJSON(&addArticleFragmentRelation)
			code := service.AddArticleFragmentRelation(addArticleFragmentRelation)
			ctx.JSON(code)
		})

		articleFragmentRelationV1.Post("/saveOrder", func(ctx iris.Context) {
			var addArticleFragmentRelations []service.ArticleFragmentRelationDetail
			ctx.ReadJSON(&addArticleFragmentRelations)
			code := service.UpdateArticleFragmentRelations(addArticleFragmentRelations)
			ctx.JSON(code)
		})

		articleFragmentRelationV1.Get("/delete", func(ctx iris.Context) {
			articleFragmentRelationID := getIntVal("articleFragmentRelationId", 0, ctx)
			var deleteArticleFragmentRelation service.ArticleFragmentRelationDetail
			deleteArticleFragmentRelation.ID = articleFragmentRelationID
			result := service.DeleteArticleFragmentRelations(deleteArticleFragmentRelation)
			ctx.JSON(result)
		})
	}
}
