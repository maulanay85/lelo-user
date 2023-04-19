package entity

type Credential struct {
	Database DatabaseCredential `yaml:"database"`
	Jwt      JwtCredential      `yaml:"jwt"`
}

type DatabaseCredential struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type JwtCredential struct {
	SecretKey string `yaml:"secret-key"`
}
