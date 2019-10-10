package routers

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//InitDocRouter 初始化文档
func InitDocRouter(app *iris.Application, crs context.Handler) {

	initTagArticle(app, crs)
	initTagFragment(app, crs)
	initTagArticleFragmentRelation(app, crs)

	initFragment(app, crs)
	initFragmentTagRelation(app, crs)
	initFragmentIntroductionRelation(app, crs)

	initArticle(app, crs)
	initArticleTagRelation(app, crs)
	initArticleFragmentRelation(app, crs)
	initArticleArticleRelation(app, crs)

	initEnvironmentLabel(app, crs)

	initDetailIntroduction(app, crs)
	initIntroductionEnvironmentRelation(app, crs)

	initQuestion(app, crs)
	initQuestionSolution(app, crs)
	initQuestionSolutionRelation(app, crs)

	initUserInfo(app, crs)
	initUserRole(app, crs)
	initUserPrivilege(app, crs)
	initUserRoleRelation(app, crs)
	initUserRolePrivilegeRelation(app, crs)
}

func getIntVal(key string, defaultVal int, ctx iris.Context) int {
	keyVal := ctx.URLParam(key)
	keyValInt, err := strconv.Atoi(keyVal)
	if err != nil {
		return defaultVal
	}
	return keyValInt
}

func postIntVal(key string, defaultVal int, ctx iris.Context) int {
	keyVal := ctx.FormValue(key)
	keyValInt, err := strconv.Atoi(keyVal)
	if err != nil {
		return defaultVal
	}
	return keyValInt
}
