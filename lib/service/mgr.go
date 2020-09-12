package service

import (
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
	}
	return mgr
}

// ------- IMPL -------
var mgr *_Mgr

type _Mgr struct {
	servMap map[config.ServiceKey]IService
}

func (m *_Mgr) Register(serv IService) {
	m.servMap[serv.ServiceKey()] = serv
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

	b.WriteString(`# rap generated docker-compose config
version: "2.0"

services:
`)

	for _, c := range configs {
		if serv, ok := m.servMap[c.Key()]; ok {
			b.WriteString(serv.GenDockerComposeConfig(c))
		}
	}

	return b.String()
}
