package config

import (
	"gin_study/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server *Server `yaml:"server" valid:"required"`
	Mysql  *Mysql  `yaml:"mysql" valid:"required"`
	Jwt    *Jwt    `yaml:"jwt" valid:"required"`
}

type Mysql struct {
	Dsn         string `yaml:"dsn" valid:"required"`
	MaxIdle     int    `yaml:"maxIdle" valid:"required"`
	MaxOpenConn int    `yaml:"maxOpenConn" valid:"required"`
}

type Server struct {
	Ip          string       `yaml:"ip" valid:"required"`
	Port        string       `yaml:"port" valid:"required"`
	EnableSSL   bool         `yaml:"enableSSL"`
	Certificate *Certificate `yaml:"certificate"`
	EnableH2C   bool         `yaml:"enableH2C"`
}

type Certificate struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

// 配置对象
var Config AppConfig

func GetConfig(env string) {
	// 打开 YAML 文件
	file, err := os.Open("./config/config-" + env + ".yaml")
	if err != nil {
		logger.Error(map[string]interface{}{"Error opening file": err.Error})
		return
	}
	defer file.Close()
	// 创建解析器
	decoder := yaml.NewDecoder(file)
	// 解析 YAML 数据
	err = decoder.Decode(&Config)
	if err != nil {
		logger.Error(map[string]interface{}{"Error decoding YAML": err.Error})
		return
	}
}
