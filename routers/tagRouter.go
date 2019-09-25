package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/sirupsen/logrus"
)

func initTag(app *iris.Application, crs context.Handler) {
	tagV1 := app.Party("/tag/v1", crs).AllowMethods(iris.MethodOptions)
	{
		tagV1.Post("/list", func(ctx iris.Context) {
			var queryTag service.Tag
			ctx.ReadJSON(&queryTag)
			lastID := postIntVal("lastId", 0, ctx)

			logrus.WithFields(logrus.Fields{
				"method":  "post",
				"lastId":  lastID,
				"id":      queryTag.ID,
				"name":    queryTag.Name,
				"tagType": queryTag.TagType,
			}).Info("/tag/v1/list")

			tags := service.QueryTags(queryTag, lastID)
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
			result := service.QueryTagByID(tagID)
			ctx.JSON(result)
		})

		tagV1.Get("/delete", func(ctx iris.Context) {
			tagID := getIntVal("tagId", 0, ctx)
			result := service.DeleteTagByID(tagID)
			ctx.JSON(result)
		})
	}
}
