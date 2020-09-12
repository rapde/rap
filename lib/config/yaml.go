package config

import (
	"log"

	"github.com/rapde/rap/lib/utils"
	"github.com/spf13/viper"
)

// RapYaml rap.yml struct
// config file should convert to rapyaml and save
type RapYaml struct {
	// Key is Service.Name, Value is Service.Service@Service.Version
	// database: mysql@5.7
	Dependencies map[string]ServiceKey     `mapstructure:"depns"`
	Configs      map[string]*ServiceConfig `mapstructure:"configs"`
}

// ToConfig convert rapyaml to app config
func (r *RapYaml) toConfig() *Config {
	svrs := make([]*Service, 0, len(r.Dependencies))

	for name, service := range r.Dependencies {
		svr := &Service{
			Name:    name,
			Service: service.Name(),
			Version: service.Version(),
			Config:  r.Configs[name],
		}
		svrs = append(svrs, svr)
	}

	return &Config{Services: svrs}
}

// LoadFromPath load from file path
func LoadFromPath(path string) {
	viper.SetConfigName("rap")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(utils.GetWorkDir())

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read %s failed %v", path, err)
	}

	viper.SetConfigType("yaml")

	rapyaml := &RapYaml{}
	err = viper.Unmarshal(rapyaml)
	if err != nil {
		log.Fatalf("Unmarshal %s failed %v", path, err)
	}

	AppConfig = rapyaml.toConfig()
	AppConfig.Path = path
}
