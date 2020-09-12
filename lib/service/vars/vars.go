package vars

import "github.com/rapde/rap/lib/config"

var (
	ServiceMap map[string]*config.Service
	Mysql5_7   = &config.Service{
		Name:    "",
		Service: "mysql",
		Version: "5.7",
		Config: &config.ServiceConfig{
			Image:      "mysql:5.7",
			Command:    "--default-authentication-plugin=mysql_native_password",
			Volumes:    []string{".rap/mysql:/var/lib/mysql"},
			Ports:      []string{"3306:3306"},
			Enviroment: map[string]string{"MYSQL_ROOT_PASSWORD": "123"},
		},
	}
)

func init() {
	ServiceMap = make(map[string]*config.Service)
	ServiceMap[string(Mysql5_7.Key())] = Mysql5_7
}
