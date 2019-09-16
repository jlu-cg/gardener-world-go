package config

import (
	ini "gopkg.in/ini.v1"
)

//WorldConfig 全局配置
type WorldConfig struct {
	PgConfig      PgConfig
	CouchdbConfig CouchdbConfig
	ServerConfig  ServerConfig
}

//New 创建配置
func New() *WorldConfig {
	config := &WorldConfig{
		PgConfig:      NewPgConfig(),
		CouchdbConfig: NewCouchdbConfig(),
		ServerConfig:  NewServerConfig(),
	}
	return config
}

//LoadConfig 加载配置
func (wc *WorldConfig) LoadConfig() {
	cfg, err := ini.Load("/Users/yueleng/dev/go/repository/src/github.com/gardener/gardener-world-go/conf/config.ini")
	if err != nil {
		panic(err)
	}
	wc.PgConfig.MaxConnection = cfg.Section("database").Key("maxConnection").MustInt64(5)
	wc.PgConfig.MinConnection = cfg.Section("database").Key("minConnection").MustInt64(5)
	wc.PgConfig.URL = cfg.Section("database").Key("url").MustString("")

	wc.CouchdbConfig.URL = cfg.Section("couchdb").Key("url").MustString("")
	wc.CouchdbConfig.UserName = cfg.Section("couchdb").Key("userName").MustString("")
	wc.CouchdbConfig.Password = cfg.Section("couchdb").Key("password").MustString("")

	wc.ServerConfig.Port = cfg.Section("Server").Key("port").MustInt64(38080)
}
