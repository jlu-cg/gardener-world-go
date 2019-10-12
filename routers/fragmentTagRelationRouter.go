package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initFragmentTagRelation(app *iris.Application, crs context.Handler) {
	fragmentTagRelationV1 := app.Party("/fragment/tag/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		fragmentTagRelationV1.Post("/list/fragment", func(ctx iris.Context) {
			var queryFragmentTagRelationWithFragment service.FragmentTagRelationWithFragment
			ctx.ReadJSON(&queryFragmentTagRelationWithFragment)
			fragmentTagRelationWithFragments := service.QueryFragmentTagRelationWithFragments(queryFragmentTagRelationWithFragment)
			ctx.JSON(fragmentTagRelationWithFragments)
		})

		fragmentTagRelationV1.Post("/list/tag", func(ctx iris.Context) {
			var queryFragmentTagRelationWithTag service.FragmentTagRelationWithTag
			ctx.ReadJSON(&queryFragmentTagRelationWithTag)
			fragmentTagRelationWithTags := service.QueryFragmentTagRelationWithTags(queryFragmentTagRelationWithTag)
			ctx.JSON(fragmentTagRelationWithTags)
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
