package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initFragment(app *iris.Application, crs context.Handler) {
	fragmentV1 := app.Party("/fragment/v1", crs).AllowMethods(iris.MethodOptions)
	{
		fragmentV1.Post("/list", func(ctx iris.Context) {
			var queryFragmentWithTag service.FragmentWithTag
			ctx.ReadJSON(&queryFragmentWithTag)
			lastID := postIntVal("lastId", 0, ctx)
			fragmentWithTags := service.QueryFragments(queryFragmentWithTag, lastID)
			ctx.JSON(fragmentWithTags)
		})

		fragmentV1.Post("/save", func(ctx iris.Context) {
			var addFragment service.Fragment
			ctx.ReadJSON(&addFragment)
			code := service.SaveFragment(addFragment)
			ctx.JSON(code)
		})

		fragmentV1.Get("/detail", func(ctx iris.Context) {
			fragmentID := getIntVal("fragmentId", 0, ctx)
			result := service.QueryFragmentByID(fragmentID)
			ctx.JSON(result)
		})

		fragmentV1.Get("/delete", func(ctx iris.Context) {
			fragmentID := getIntVal("fragmentId", 0, ctx)
			result := service.DeleteFragmentByID(fragmentID)
			ctx.JSON(result)
		})
	}
}
