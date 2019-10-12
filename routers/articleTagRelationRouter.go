package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticleTagRelation(app *iris.Application, crs context.Handler) {
	articleTagRelationV1 := app.Party("/article/tag/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleTagRelationV1.Post("/list/article", func(ctx iris.Context) {
			var queryArticleTagRelationWithArticle service.ArticleTagRelationWithArticle
			ctx.ReadJSON(&queryArticleTagRelationWithArticle)
			articleTagRelationWithArticles := service.QueryArticleTagRelationWithArticles(queryArticleTagRelationWithArticle)
			ctx.JSON(articleTagRelationWithArticles)
		})

		articleTagRelationV1.Post("/list/tag", func(ctx iris.Context) {
			var queryArticleTagRelationWithTag service.ArticleTagRelationWithTag
			ctx.ReadJSON(&queryArticleTagRelationWithTag)
			articleTagRelationWithTags := service.QueryArticleTagRelationWithTags(queryArticleTagRelationWithTag)
			ctx.JSON(articleTagRelationWithTags)
		})

		articleTagRelationV1.Post("/save", func(ctx iris.Context) {
			var addArticleTagRelation service.ArticleTagRelation
			ctx.ReadJSON(&addArticleTagRelation)
			code := service.SaveArticleTagRelation(addArticleTagRelation)
			ctx.JSON(code)
		})

		articleTagRelationV1.Get("/delete", func(ctx iris.Context) {
			articleTagRelationID := getIntVal("articleTagRelationId", 0, ctx)
			result := service.DeleteArticleTagRelationByID(articleTagRelationID)
			ctx.JSON(result)
		})
	}
}
