package service

import (
	"github.com/rapde/rap/lib/service/vars"
	"strings"

	"github.com/rapde/rap/lib/config"
)

// IMgr 服务管理器接口定义
type IMgr interface {
	// Register 注册服务
	Register(serv IService)
	// SupportedServices 获取支持的服务列表
	SupportedServices() []config.ServiceKey
	// GenDockerCompose 根据服务配置，生成 DockerCompose 配置文件
	GenDockerComposeConfig(configs []*config.Service) string
}

// Mgr 获取 Mgr 单例
func Mgr() IMgr {
	if mgr == nil {
		mgr = &_Mgr{
			servMap: make(map[config.ServiceKey]IService),
		}

		for _, v := range vars.ServiceMap {
			mgr.Register(NewService(v))
		}
	}
	return mgr
}

// ------- IMPL -------
var mgr *_Mgr

type _Mgr struct {
	servMap map[config.ServiceKey]IService
}

func (m *_Mgr) Register(serv IService) {
	mgr.servMap[serv.ServiceKey()] = serv
}

func (m *_Mgr) SupportedServices() []config.ServiceKey {
	ret := make([]config.ServiceKey, len(m.servMap))

	for key := range m.servMap {
		ret = append(ret, key)
	}

	return ret
}

func (m *_Mgr) GenDockerComposeConfig(configs []*config.Service) string {
	var b strings.Builder

	b.WriteString("# rap generated docker-compose config\n\nversion: \"2.0\"\n\nservices:")

	for _, c := range configs {
		if serv, ok := m.servMap[c.Key()]; ok {
			b.WriteString(serv.GenDockerComposeConfig(c))
		}
	}

	return b.String()
}
