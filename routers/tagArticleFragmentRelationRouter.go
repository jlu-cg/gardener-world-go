package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initTagArticleFragmentRelation(app *iris.Application, crs context.Handler) {
	tagArticleFragmentRelationV1 := app.Party("/tag/article/fragment/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagArticleFragmentRelationV1.Post("/list", func(ctx iris.Context) {
			var queryTagArticleFragmentRelationDetail service.TagArticleFragmentRelationDetail
			ctx.ReadJSON(&queryTagArticleFragmentRelationDetail)
			tagArticleFragmentRelations := service.QueryTagArticleFragmentRelationDetails(queryTagArticleFragmentRelationDetail)
			ctx.JSON(tagArticleFragmentRelations)
		})

		tagArticleFragmentRelationV1.Post("/save", func(ctx iris.Context) {
			var addTagArticleFragmentRelation service.TagArticleFragmentRelation
			ctx.ReadJSON(&addTagArticleFragmentRelation)
			code := service.SaveTagArticleFragmentRelation(addTagArticleFragmentRelation)
			ctx.JSON(code)
		})

		tagArticleFragmentRelationV1.Get("/delete", func(ctx iris.Context) {
			tagArticleFragmentRelationID := getIntVal("tagArticleFragmentRelationId", 0, ctx)
			result := service.DeleteTagArticleFragmentRelationByID(tagArticleFragmentRelationID)
			ctx.JSON(result)
		})
	}
}
