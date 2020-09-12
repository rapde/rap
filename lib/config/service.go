package config

import (
	"fmt"
	"strings"
)

// Service 依赖服务定义
type Service struct {
	Name    string         // 用户指定的服务名称，如：userdb...
	Service string         // 依赖服务名，如：mysql、mongo、redis...
	Version string         // 依赖服务版本，如 5.7
	Config  *ServiceConfig // 依赖服务的额外配置
}

// ServiceConfig Service 配置
type ServiceConfig struct {
	Volumes    []string          // 容器卷信息
	Ports      []string          // 容器端口映射
	Enviroment map[string]string // 容器环境变量
}

// Key 生成服务 key，用于区分不同的服务类型
func (s *Service) Key() ServiceKey {
	return NewServiceKey(s.Service, s.Version)
}

// ServiceKey 服务标识符
type ServiceKey string

// Name 获取服务名称
func (k ServiceKey) Name() string {
	return strings.Split(string(k), "@")[0]
}

// Version 获取服务版本
func (k ServiceKey) Version() string {
	return strings.Split(string(k), "@")[1]
}

// NewServiceKey 根据服务名称和版本组装服务标识符
func NewServiceKey(name string, version string) ServiceKey {
	return ServiceKey(fmt.Sprintf("%s@%s", name, version))
}
