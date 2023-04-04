package config

type ServerConfig struct {
	HostPort    string            `yaml:"host-port" json:"host-port" mapstructure:"host-port"`
	LogConfig   LogConfig         `yaml:"log" json:"log" mapstructure:"log" `
	DictMaxSize int               `yaml:"dict-max-size" json:"dict-max-size" mapstructure:"dict-max-size"`
	BasicAuth   map[string]string `yaml:"basic-auth" json:"basic-auth" mapstructure:"basic-auth"`
	RedisConfig RedisConfig       `yaml:"redis" json:"redis" mapstructure:"redis"`
}

type LogConfig struct {
	Level           int    `yaml:"level" json:"level" mapstructure:"level"`
	File            string `yaml:"file" json:"file" mapstructure:"file"`
	AccessLogFormat string `yaml:"access-log-format" json:"access-log-format" mapstructure:"access-log-format"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr" json:"addr" mapstructure:"addr"`
	Password string `yaml:"password" json:"password" mapstructure:"password"`
	Db       int    `yaml:"db" json:"db" mapstructure:"db"`
}
