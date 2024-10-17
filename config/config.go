package config

import (
	"gin_study/logger"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server    *Server `yaml:"server"`
	Mysql     *Mysql  `yaml:"mysql"`
	Jwt       *Jwt    `yaml:"jwt"`
	Redis     *Redis  `yaml:"redis"`
	EmailSmpt *Email  `yaml:"emailSmpt"`
}

type Mysql struct {
	Dsn         string `yaml:"dsn"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

type Server struct {
	Ip          string       `yaml:"ip"`
	Port        string       `yaml:"port"`
	EnableSSL   bool         `yaml:"enableSSL"`
	Certificate *Certificate `yaml:"certificate"`
	EnableH2C   bool         `yaml:"enableH2C"`
}

type Certificate struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type Jwt struct {
	Secret   string `yaml:"secret"`
	RedisKey string `yaml:"redisKey"`
}

type Redis struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Db          int    `yaml:"db"`
	Password    string `yaml:"password"`
	IdleTimeout int64  `yaml:"idleTimeout"`
}

type Email struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	RedisKey string `yaml:"redisKey"`
}

// 配置对象
var Config AppConfig

func GetConfig(env string) {
	// 打开 YAML 文件
	file, err := os.Open("./config/application-" + env + ".yaml")
	if err != nil {
		logger.Error(zap.Any("Error opening file", err))
		return
	}
	defer file.Close()
	// 创建解析器
	decoder := yaml.NewDecoder(file)
	// 解析 YAML 数据
	err = decoder.Decode(&Config)
	if err != nil {
		logger.Error(zap.Any("Error decoding YAML", err))
		return
	}
}
