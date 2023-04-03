package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	HostPort    string            `yaml:"host-port" json:"host-port"`
	LogConfig   LogConfig         `yaml:"log" json:"log"`
	DictMaxSize int               `yaml:"dict-max-size" json:"dict-max-size"`
	BasicAuth   map[string]string `yaml:"basic-auth" json:"basic-auth"`
	RedisConfig RedisConfig       `yaml:"redis" json:"redis"`
}

type LogConfig struct {
	Level           int    `yaml:"level" json:"level"`
	File            string `yaml:"file" json:"file"`
	AccessLogFormat string `yaml:"access-log-format" json:"access-log-format"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr" json:"addr"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}

var ConfigInstance ServerConfig

func init() {

	var configFile string
	flag.StringVar(&configFile, "c", "./config.yaml", "short for conf, specific the config file or default './config'")
	flag.StringVar(&configFile, "conf", "./config.yaml", "specific the config file or default './config'")
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(bytes, &ConfigInstance)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}
	if ConfigInstance.LogConfig.AccessLogFormat == "" {
		ConfigInstance.LogConfig.AccessLogFormat = "[${time}] ${status} - ${latency} ${method} ${path}"
	}

	if ConfigInstance.DictMaxSize <= 0 {
		ConfigInstance.DictMaxSize = 2000
	}
}
