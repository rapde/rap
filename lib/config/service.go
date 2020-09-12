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
	Image      string            `mapstructure:"image"`      // 镜像
	Command    string            `mapstructure:"command"`    // 命令
	Volumes    []string          `mapstructure:"volumes"`    // 容器卷信息
	Ports      []string          `mapstructure:"ports"`      // 容器端口映射
	Enviroment map[string]string `mapstructure:"enviroment"` // 容器环境变量
}

// Key 生成服务 key，用于区分不同的服务类型
func (s *Service) Key() ServiceKey {
	return NewServiceKey(s.Service, s.Version)
}

// 合并两个 service 以 new 为准进行覆盖, 如果字段为空取默认值
func (s *Service) Merge(new *Service) *Service {
	if new == nil {
		return s
	}

	if new.Config == nil {
		new.Config = &ServiceConfig{
			Image:      s.Config.Image,
			Command:    s.Config.Command,
			Volumes:    s.Config.Volumes,
			Ports:      s.Config.Ports,
			Enviroment: s.Config.Enviroment,
		}

		return new
	}

	if new.Config.Image == "" {
		new.Config.Image = s.Config.Image
	}

	if new.Config.Command == "" {
		new.Config.Command = s.Config.Command
	}

	if new.Config.Volumes == nil {
		new.Config.Volumes = s.Config.Volumes
	}

	if new.Config.Ports == nil {
		new.Config.Ports = s.Config.Ports
	}

	if new.Config.Enviroment == nil {
		new.Config.Enviroment = s.Config.Enviroment
	}

	return new
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
