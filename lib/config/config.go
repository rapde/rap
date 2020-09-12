package config

// Config rap 服务配置，解析自 rap.yaml
type Config struct {
	Services []*Service // 服务列表
}

// Parse 从字符串解析配置信息
func Parse(configStr string) (config *Config, err error) {
	// TODO
	return
}
