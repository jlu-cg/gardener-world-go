package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryQuestionSolution struct {
	service.QuestionSolution
	LastID int `json:"lastId"`
}

func initQuestionSolution(app *iris.Application, crs context.Handler) {
	questionSolutionV1 := app.Party("/question/solution/v1", crs).AllowMethods(iris.MethodOptions)
	{
		questionSolutionV1.Post("/list", func(ctx iris.Context) {
			var queryQuestionSolution queryQuestionSolution
			ctx.ReadJSON(&queryQuestionSolution)
			questionSolutions := service.QueryQuestionSolutions(queryQuestionSolution.QuestionSolution, queryQuestionSolution.LastID)
			ctx.JSON(questionSolutions)
		})

		questionSolutionV1.Post("/save", func(ctx iris.Context) {
			var saveQuestionSolution service.QuestionSolution
			ctx.ReadJSON(&saveQuestionSolution)
			code := service.SaveQuestionSolution(saveQuestionSolution)
			ctx.JSON(code)
		})

		questionSolutionV1.Get("/detail", func(ctx iris.Context) {
			questionSolutionID := getIntVal("questionSolutionId", 0, ctx)
			result := service.QueryQuestionSolutionByID(questionSolutionID)
			ctx.JSON(result)
		})

		questionSolutionV1.Get("/delete", func(ctx iris.Context) {
			questionSolutionID := getIntVal("questionSolutionId", 0, ctx)
			var deleteQuestionSolution service.QuestionSolution
			deleteQuestionSolution.ID = questionSolutionID
			result := service.DeleteQuestionSolutions(deleteQuestionSolution)
			ctx.JSON(result)
		})
	}
}
