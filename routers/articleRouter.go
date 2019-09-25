package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initArticle(app *iris.Application, crs context.Handler) {
	articleV1 := app.Party("/article/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleV1.Post("/list", func(ctx iris.Context) {
			var queryArticle service.Article
			ctx.ReadJSON(&queryArticle)
			lastID := postIntVal("lastId", 0, ctx)
			articles := service.QueryArticles(queryArticle, lastID)
			ctx.JSON(articles)
		})

		articleV1.Post("/save", func(ctx iris.Context) {
			var addArticle service.Article
			ctx.ReadJSON(&addArticle)
			code := service.SaveArticle(addArticle)
			ctx.JSON(code)
		})

		articleV1.Get("/detail", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			tagID := getIntVal("tagId", 0, ctx)
			result := service.GetArticleByID(articleID, tagID)
			ctx.JSON(result)
		})

		articleV1.Get("/delete", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			result := service.DeleteArticleByID(articleID)
			ctx.JSON(result)
		})

		articleV1.Get("/document/generate", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			result := service.CouchdbArticleGenerateDocument(articleID)
			ctx.JSON(result)
		})

		articleV1.Get("/document/cancel", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			result := service.CouchdbArticleCancelDocument(articleID)
			ctx.JSON(result)
		})

		articleV1.Get("/document/detail", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			result := service.CouchdbGetArticleDocumentByArticleID(articleID)
			ctx.JSON(result)
		})
	}

}
