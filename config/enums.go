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
