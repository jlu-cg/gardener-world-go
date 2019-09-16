package service

import (
	"database/sql"

	"github.com/gardener/gardener-world-go/config"

	_ "github.com/lib/pq"
)

const maxConnection int = 2

var (
	connectionPool = make(chan *sql.DB, maxConnection)
)

//InitPool 初始化连接池
func InitPool(config *config.WorldConfig) {
	index := 0
	for {
		db, err := sql.Open("postgres", config.PgConfig.URL)
		if err != nil {
			panic(err)
		}
		connectionPool <- db
		index++
		if index >= maxConnection {
			break
		}
	}
}

//Connect 获取连接
func connect() *sql.DB {
	return <-connectionPool
}

//Release 释放连接
func release(db *sql.DB) {
	connectionPool <- db
}

//Close 关闭连接池
func close() {
	for {
		db, _ := <-connectionPool
		defer db.Close()
	}
}
