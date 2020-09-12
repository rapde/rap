package service

import "github.com/rapde/rap/lib/config"

func init() {
	Mgr().Register(&mysql57{})
}

type mysql57 struct{}

// ServiceKey 获取该 service 实现对应的服务标识符
func (m *mysql57) ServiceKey() config.ServiceKey {
	return config.NewServiceKey("mysql", "5.7")
}

// GenDockerComposeConfig 根据服务配置，生成 DockerCompose 文件配置
func (m *mysql57) GenDockerComposeConfig(servConf *config.Service) string {
	panic("not implemented") // TODO: Implement
}
