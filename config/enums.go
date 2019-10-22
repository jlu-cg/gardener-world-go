package config

const (
	DBSuccess = iota
	DBErrorSQLNoCondition
	DBErrorSQLMissKeyCondition
	DBErrorConnection
	DBErrorExecution
	DBErrorCursor
	DBErrorNoData
)

const (
	ArticleStatusNotPublish = iota
	ArticleStatusPublish
)
