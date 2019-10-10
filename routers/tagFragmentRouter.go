package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initTagFragment(app *iris.Application, crs context.Handler) {
	tagFragmentV1 := app.Party("/tag/fragment/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagFragmentV1.Post("/list", func(ctx iris.Context) {
			var queryTagFragment service.TagFragment
			ctx.ReadJSON(&queryTagFragment)
			lastID := postIntVal("lastId", 0, ctx)

			tagFragments := service.QueryTagFragments(queryTagFragment, lastID)
			ctx.JSON(tagFragments)
		})

		tagFragmentV1.Post("/save", func(ctx iris.Context) {
			var addTagFragment service.TagFragment
			ctx.ReadJSON(&addTagFragment)
			code := service.SaveTagFragment(addTagFragment)
			ctx.JSON(code)
		})

		tagFragmentV1.Get("/detail", func(ctx iris.Context) {
			tagFragmentID := getIntVal("tagFragmentId", 0, ctx)
			result := service.QueryTagFragmentByID(tagFragmentID)
			ctx.JSON(result)
		})

		tagFragmentV1.Get("/delete", func(ctx iris.Context) {
			tagFragmentID := getIntVal("tagFragmentId", 0, ctx)
			result := service.DeleteTagFragmentByID(tagFragmentID)
			ctx.JSON(result)
		})
	}
}
