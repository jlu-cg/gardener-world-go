package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initQuestionSolutionRelation(app *iris.Application, crs context.Handler) {
	questionSolutionRelationV1 := app.Party("/question/solution/relation/v1", crs).AllowMethods(iris.MethodOptions)
	{
		questionSolutionRelationV1.Post("/list", func(ctx iris.Context) {
			var queryQuestionSolutionRelationDetail service.QuestionSolutionRelationDetail
			ctx.ReadJSON(&queryQuestionSolutionRelationDetail)
			questionSolutionRelationDetails := service.QueryQuestionSolutionRelationDetails(queryQuestionSolutionRelationDetail)
			ctx.JSON(questionSolutionRelationDetails)
		})

		questionSolutionRelationV1.Post("/save", func(ctx iris.Context) {
			var addQuestionSolutionRelation service.QuestionSolutionRelation
			ctx.ReadJSON(&addQuestionSolutionRelation)
			code := service.SaveQuestionSolutionRelation(addQuestionSolutionRelation)
			ctx.JSON(code)
		})

		questionSolutionRelationV1.Get("/delete", func(ctx iris.Context) {
			questionSolutionRelationID := getIntVal("questionSolutionRelationId", 0, ctx)
			var deleteQuestionSolutionRelation service.QuestionSolutionRelation
			deleteQuestionSolutionRelation.ID = questionSolutionRelationID
			result := service.DeleteQuestionSolutionRelations(deleteQuestionSolutionRelation)
			ctx.JSON(result)
		})
	}
}
