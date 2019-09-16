package config

//PgConfig postgresql配置
type PgConfig struct {
	MinConnection int64
	MaxConnection int64
	URL           string
}

//NewPgConfig 创建配置
func NewPgConfig() PgConfig {
	return PgConfig{
		MinConnection: 5,
		MaxConnection: 5,
		URL:           "",
	}
}
