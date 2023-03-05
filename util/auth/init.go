package auth

type UtilAuthModule struct{}

type UtilAuth interface {
	HashPassword(string) (string, error)
	CheckHashPassword(pass string, hashPass string) bool
}

func NewUtilAuth() *UtilAuthModule {
	return &UtilAuthModule{}

}
