// Package service 服务接口定义
package service

import (
	"bytes"
	"github.com/rapde/rap/lib/config"
	"text/template"
)

// IService 服务接口
// 每个服务需要实现自己的 docker-compose 配置
type IService interface {
	// ServiceKey 获取该 service 实现对应的服务标识符
	ServiceKey() config.ServiceKey
	// GenDockerComposeConfig 根据服务配置，生成 DockerCompose 文件配置
	GenDockerComposeConfig(servConf *config.Service) string
}

type Service struct {
	*config.Service
}

func NewService(conf *config.Service) *Service {
	return &Service{
		conf,
	}
}

func (s *Service) ServiceKey() config.ServiceKey {
	return s.Key()
}

func (s *Service) GenDockerComposeConfig(userConf *config.Service) string {
	var buf bytes.Buffer
	res := s.Merge(userConf)
	tmpl.Execute(&buf, res)
	return buf.String()
}

var tmpl = template.Must(template.New("tmpl").Parse(`
  {{.Name}}:
    image: "{{.Config.Image}}"
    command: {{.Config.Command}}
    volumes:
		{{range .Config.Volumes}}
      - {{.}}
		{{end}}
    ports:
		{{range .Config.Ports}}
      - "{{.}}"
		{{end}}
    environment:
		{{range $key, $value := .Config.Environments}}
        {{$key}}: {{$value}}
		{{end}}
`))
