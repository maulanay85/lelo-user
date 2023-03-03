package entity

type Credential struct {
	Database DatabaseCredential `yaml:"database"`
}

type DatabaseCredential struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
