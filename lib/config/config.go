package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

// AppConfig global app config
var AppConfig *Config

// Config rap 服务配置，解析自 rap.yaml
type Config struct {
	// config file path
	Path     string
	Services []*Service // 服务列表
}

// Services 服务列表
type Services []*Service

// Len implement sort interface
func (s Services) Len() int {
	return len(s)
}

// Less implement sort interface
func (s Services) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap implement sort interface
func (s Services) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort yaml file
func (c *Config) toRapYamlBuf() *bytes.Buffer {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("depns:\n")

	sort.Sort(Services(c.Services))

	for _, v := range c.Services {
		buf.WriteString(fmt.Sprintf("  %s: %s@%s\n", v.Name, v.Service, v.Version))
	}

	buf.WriteString("configs:\n")
	for _, v := range c.Services {
		if v.Config == nil {
			continue
		}
		buf.WriteString(fmt.Sprintf("  %s:\n", v.Name))
		conf := v.Config
		if len(conf.Volumes) > 0 {
			buf.WriteString("    volumes:\n")
			for _, vv := range conf.Volumes {
				buf.WriteString(fmt.Sprintf("      - %s\n", vv))
			}
		}

		if len(conf.Ports) > 0 {
			buf.WriteString("    ports:\n")
			for _, vv := range conf.Ports {
				buf.WriteString(fmt.Sprintf("      - %s\n", vv))
			}
		}

		if len(conf.Environment) > 0 {
			buf.WriteString("    enviroment:\n")
			for k, vv := range conf.Environment {
				buf.WriteString(fmt.Sprintf("      %s: %s", k, vv))
			}
		}
	}

	return buf
}

// Save save to disk
func (c *Config) Save() {
	buf := c.toRapYamlBuf()
	f, err := os.OpenFile(c.Path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Save %s failed %+v", c.Path, err)
	}
	defer f.Close()

	f.Write(buf.Bytes())
}

// Parse 从字符串解析配置信息
func Parse(configStr string) (config *Config, err error) {

	return
}
