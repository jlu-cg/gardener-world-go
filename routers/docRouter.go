package routers

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//InitDocRouter 初始化文档
func InitDocRouter(app *iris.Application, crs context.Handler) {

	initTag(app, crs)

	initTagTagRelation(app, crs)

	initFragment(app, crs)

	initFragmentTagRelation(app, crs)

	initArticle(app, crs)

	initArticleTagRelation(app, crs)

	initArticleFragmentRelation(app, crs)

	initArticleArticleRelation(app, crs)
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
