// Package service 服务接口定义
package service

import "github.com/rapde/rap/lib/config"

// IService 服务接口
// 每个服务需要实现自己的 docker-compose 配置
type IService interface {
	// ServiceKey 获取该 service 实现对应的服务标识符
	ServiceKey() config.ServiceKey
	// GenDockerComposeConfig 根据服务配置，生成 DockerCompose 文件配置
	GenDockerComposeConfig(servConf *config.Service) string
}
