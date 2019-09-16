package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initTag(app *iris.Application, crs context.Handler) {
	tagV1 := app.Party("/tag/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagV1.Post("/list", func(ctx iris.Context) {
			var queryTag service.Tag
			ctx.ReadJSON(&queryTag)
			lastID := postIntVal("lastId", 0, ctx)
			tags := service.GetTags(queryTag, lastID)
			ctx.JSON(tags)
		})

		tagV1.Post("/save", func(ctx iris.Context) {
			var addTag service.Tag
			ctx.ReadJSON(&addTag)
			code := service.SaveTag(addTag)
			ctx.JSON(code)
		})

		tagV1.Get("/detail", func(ctx iris.Context) {
			tagID := getIntVal("tagId", 0, ctx)
			result := service.GetTagByID(tagID)
			ctx.JSON(result)
		})
	}
}
