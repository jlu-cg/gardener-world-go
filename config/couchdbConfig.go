package config

//CouchdbConfig 配置
type CouchdbConfig struct {
	URL      string
	UserName string
	Password string
}

//NewCouchdbConfig 创建配置
func NewCouchdbConfig() CouchdbConfig {
	return CouchdbConfig{
		URL:      "",
		UserName: "",
		Password: "",
	}
}
