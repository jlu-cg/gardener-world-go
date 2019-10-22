package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryTagArticle struct {
	service.TagArticle
	LastID int `json:"lastId"`
}

func initTagArticle(app *iris.Application, crs context.Handler) {
	tagArticleV1 := app.Party("/tag/article/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagArticleV1.Post("/list", func(ctx iris.Context) {
			var queryTagArticle queryTagArticle
			ctx.ReadJSON(&queryTagArticle)

			tagArticles := service.QueryTagArticles(queryTagArticle.TagArticle, queryTagArticle.LastID)
			ctx.JSON(tagArticles)
		})

		tagArticleV1.Post("/save", func(ctx iris.Context) {
			var addTagArticle service.TagArticle
			ctx.ReadJSON(&addTagArticle)
			code := service.SaveTagArticle(addTagArticle)
			ctx.JSON(code)
		})

		tagArticleV1.Get("/detail", func(ctx iris.Context) {
			tagArticleID := getIntVal("tagArticleId", 0, ctx)
			result := service.QueryTagArticleByID(tagArticleID)
			ctx.JSON(result)
		})

		tagArticleV1.Get("/delete", func(ctx iris.Context) {
			tagArticleID := getIntVal("tagArticleId", 0, ctx)
			result := service.DeleteTagArticleByID(tagArticleID)
			ctx.JSON(result)
		})
	}
}
