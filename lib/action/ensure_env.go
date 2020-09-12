package action

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/rapde/rap/lib/config"
	"github.com/rapde/rap/lib/service"
	"github.com/rapde/rap/lib/utils"
)

// EnsureEnv 检查并初始化 rap 环境
func EnsureEnv() {
	wd := utils.GetWorkDir()

	// 1. 如果 rap.yaml 文件不存在，创建一个
	rapYamlPath := path.Join(wd, "rap.yaml")
	if !utils.Exists(rapYamlPath) {
		c := config.Config{
			Path: rapYamlPath,
		}
		c.Save()
	}

	if !utils.IsFile(rapYamlPath) {
		log.Fatalln("rap.yaml conflict")
	}

	// 加载配置文件
	config.LoadFromPath(rapYamlPath)

	// 2. 如果 .rap 目录不存在，创建一个
	rapDirPath := path.Join(wd, ".rap")
	if !utils.Exists(rapDirPath) {
		err := os.Mkdir(rapDirPath, 0760)
		if err != nil {
			log.Fatalf("Failed to create .rap folder %v", err)
		}
	}

	if !utils.IsDir(rapDirPath) {
		log.Fatalln(".rap folder conflict")
	}

	// 3. 检查 Docker 是否存在
	if _, ok := utils.GetBinPath("docker"); !ok {
		log.Fatalln("please install docker first: https://docs.docker.com/engine/install/")
	}

	// 4. 检查 docker-compose 是否存在
	if _, ok := utils.GetBinPath("docker-compose"); !ok {
		log.Fatalln("please install docker-compose first: https://docs.docker.com/compose/install/")
	}

	// 5. 根据 rap.yaml 生成 docker-compose file
	dockerComposeFilePath := path.Join(wd, ".rap/docker-compose.yml")
	err := ioutil.WriteFile(dockerComposeFilePath, []byte(service.Mgr().GenDockerComposeConfig(config.AppConfig.Services)), 0660)
	if err != nil {
		log.Fatalf("failed to save docker-compose.yml %v", err)
	}
}
