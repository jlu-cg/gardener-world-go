package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryArticleWithTag struct {
	service.ArticleWithTag
	LastID int `json:"lastId"`
}

func initArticle(app *iris.Application, crs context.Handler) {
	articleV1 := app.Party("/article/v1", crs).AllowMethods(iris.MethodOptions)
	{
		articleV1.Post("/list", func(ctx iris.Context) {
			var queryArticleWithTag queryArticleWithTag
			ctx.ReadJSON(&queryArticleWithTag)
			articleWithTags := service.QueryArticles(queryArticleWithTag.ArticleWithTag, queryArticleWithTag.LastID)
			ctx.JSON(articleWithTags)
		})

		articleV1.Post("/save", func(ctx iris.Context) {
			var addArticle service.Article
			ctx.ReadJSON(&addArticle)
			code := service.SaveArticle(addArticle)
			ctx.JSON(code)
		})

		articleV1.Get("/detail", func(ctx iris.Context) {
			articleID := getIntVal("articleId", 0, ctx)
			result := service.QueryArticleByID(articleID)
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
