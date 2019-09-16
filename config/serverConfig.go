package config

//ServerConfig 服务器配置
type ServerConfig struct {
	Port int64
}

//NewServerConfig 创建配置
func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port: 8081,
	}
}
