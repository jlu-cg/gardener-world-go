package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initTagTagRelation(app *iris.Application, crs context.Handler) {
	tagTagRelationV1 := app.Party("/tag/tag/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagTagRelationV1.Post("/list", func(ctx iris.Context) {
			var queryTagTagRelation service.TagTagRelation
			ctx.ReadJSON(&queryTagTagRelation)
			tagTagRelations := service.GetTagTagRelations(queryTagTagRelation)
			ctx.JSON(tagTagRelations)
		})

		tagTagRelationV1.Post("/save", func(ctx iris.Context) {
			var addTagTagRelation service.TagTagRelation
			ctx.ReadJSON(&addTagTagRelation)
			code := service.AddTagTagRelation(addTagTagRelation)
			ctx.JSON(code)
		})

		tagTagRelationV1.Get("/delete", func(ctx iris.Context) {
			tagTagRelationID := getIntVal("tagTagRelationId", 0, ctx)
			result := service.DeleteTagTagRelationByID(tagTagRelationID)
			ctx.JSON(result)
		})
	}
}
