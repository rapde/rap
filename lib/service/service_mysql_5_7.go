package service

import (
	"bytes"
	"html/template"

	"github.com/rapde/rap/lib/config"
)

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
	// TODO 读取配置文件中客户自定义的 config
	// TODO 合并自定义 config 到默认 config
	// TODO 根据 config 生成 DockerCompose 文件内容
	var buf bytes.Buffer
	mysql57Tmpl.Execute(&buf, servConf)
	return buf.String()
}

var mysql57Tmpl = template.Must(template.New("mysql57").Parse(`
  {{.Name}}:
    image: "mysql:5.7"
    command: --default-authentication-plugin=mysql_native_password
    volumes:
      - .rap/mysql:/var/lib/mysql
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: 123
`))
