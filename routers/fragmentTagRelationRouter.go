package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initFragmentTagRelation(app *iris.Application, crs context.Handler) {
	fragmentTagRelationV1 := app.Party("/fragment/tag/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		fragmentTagRelationV1.Post("/list", func(ctx iris.Context) {
			var queryFragmentTagRelation service.FragmentTagRelation
			ctx.ReadJSON(&queryFragmentTagRelation)
			fragmentTagRelations := service.QueryFragmentTagRelations(queryFragmentTagRelation)
			ctx.JSON(fragmentTagRelations)
		})

		fragmentTagRelationV1.Post("/save", func(ctx iris.Context) {
			var addFragmentTagRelation service.FragmentTagRelation
			ctx.ReadJSON(&addFragmentTagRelation)
			code := service.SaveFragmentTagRelation(addFragmentTagRelation)
			ctx.JSON(code)
		})

		fragmentTagRelationV1.Get("/delete", func(ctx iris.Context) {
			fragmentTagRelationID := getIntVal("fragmentTagRelationId", 0, ctx)
			result := service.DeleteFragmentTagRelationByID(fragmentTagRelationID)
			ctx.JSON(result)
		})
	}
}
