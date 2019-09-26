package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initIntroductionEnvironmentRelation(app *iris.Application, crs context.Handler) {
	introductionEnvironmentRelationV1 := app.Party("/introduction/environment/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		introductionEnvironmentRelationV1.Post("/list", func(ctx iris.Context) {
			var queryIntroductionEnvironmentRelationDetail service.IntroductionEnvironmentRelationDetail
			ctx.ReadJSON(&queryIntroductionEnvironmentRelationDetail)
			details := service.QueryIntroductionEnvironmentRelations(queryIntroductionEnvironmentRelationDetail)
			ctx.JSON(details)
		})

		introductionEnvironmentRelationV1.Post("/save", func(ctx iris.Context) {
			var addIntroductionEnvironmentRelation service.IntroductionEnvironmentRelation
			ctx.ReadJSON(&addIntroductionEnvironmentRelation)
			code := service.SaveIntroductionEnvironmentRelation(addIntroductionEnvironmentRelation)
			ctx.JSON(code)
		})

		introductionEnvironmentRelationV1.Get("/delete", func(ctx iris.Context) {
			introductionEnvironmentRelationID := getIntVal("introductionEnvironmentRelationId", 0, ctx)
			var relation service.IntroductionEnvironmentRelation
			relation.ID = introductionEnvironmentRelationID
			result := service.DeleteIntroductionEnvironmentRelations(relation)
			ctx.JSON(result)
		})
	}
}
