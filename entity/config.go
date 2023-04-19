package entity

type Config struct {
	Port        int16    `yaml:"port"`
	ServiceName string   `yaml:"service_name"`
	Database    Database `yaml:"database"`
	Env         string   `yaml:"env"`
	Exp         Exp      `yaml:"exp"`
}

type Database struct {
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	Port int32  `yaml:"port"`
}

type Exp struct {
	Token        int32 `yaml:"token"`
	RefreshToken int32 `yaml:"refresh_token"`
}
