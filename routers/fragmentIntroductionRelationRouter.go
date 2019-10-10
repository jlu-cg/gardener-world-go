package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initFragmentIntroductionRelation(app *iris.Application, crs context.Handler) {
	fragmentIntroductionRelationV1 := app.Party("/fragment/introduction/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		fragmentIntroductionRelationV1.Post("/list", func(ctx iris.Context) {
			var queryFragmentIntroductionRelationWithIntroduction service.FragmentIntroductionRelationWithIntroduction
			ctx.ReadJSON(&queryFragmentIntroductionRelationWithIntroduction)
			fragmentIntroductionRelationWithIntroductions := service.QueryFragmentIntroductionRelationWithIntroductions(queryFragmentIntroductionRelationWithIntroduction, 0)
			ctx.JSON(fragmentIntroductionRelationWithIntroductions)
		})

		fragmentIntroductionRelationV1.Post("/save", func(ctx iris.Context) {
			var addFragmentIntroductionRelation service.FragmentIntroductionRelation
			ctx.ReadJSON(&addFragmentIntroductionRelation)
			code := service.SaveFragmentIntroductionRelation(addFragmentIntroductionRelation)
			ctx.JSON(code)
		})

		fragmentIntroductionRelationV1.Get("/delete", func(ctx iris.Context) {
			fragmentIntroductionRelationID := getIntVal("fragmentIntroductionRelationId", 0, ctx)
			result := service.DeleteFragmentIntroductionRelationByID(fragmentIntroductionRelationID)
			ctx.JSON(result)
		})
	}
}
