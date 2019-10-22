package routers

import (
	"github.com/gardener/gardener-world-go/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type queryQuestion struct {
	service.Question
	LastID int `json:"lastId"`
}

func initQuestion(app *iris.Application, crs context.Handler) {
	questionV1 := app.Party("/question/v1", crs).AllowMethods(iris.MethodOptions)
	{
		questionV1.Post("/list", func(ctx iris.Context) {
			var queryQuestion queryQuestion
			ctx.ReadJSON(&queryQuestion)
			questions := service.QueryQuestions(queryQuestion.Question, queryQuestion.LastID)
			ctx.JSON(questions)
		})

		questionV1.Post("/save", func(ctx iris.Context) {
			var saveQuestion service.Question
			ctx.ReadJSON(&saveQuestion)
			code := service.SaveQuestion(saveQuestion)
			ctx.JSON(code)
		})

		questionV1.Get("/detail", func(ctx iris.Context) {
			questionID := getIntVal("questionId", 0, ctx)
			result := service.QueryQuestionByID(questionID)
			ctx.JSON(result)
		})

		questionV1.Get("/delete", func(ctx iris.Context) {
			questionID := getIntVal("questionId", 0, ctx)
			var deleteQuestion service.Question
			deleteQuestion.ID = questionID
			result := service.DeleteQuestions(deleteQuestion)
			ctx.JSON(result)
		})
	}
}
