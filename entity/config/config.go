package config

type Config struct {
	Port        int16    `yaml:"port"`
	ServiceName string   `yaml:"service_name"`
	Database    Database `yaml:"database"`
	Env         string   `yaml:"env"`
}

type Database struct {
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	Port int32  `yaml:"port"`
}
