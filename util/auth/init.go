package auth

import (
	entity "lelo-user/entity"
)

type UtilAuthModule struct{}

type UtilAuth interface {
	HashPassword(string) (string, error)
	CheckHashPassword(pass string, hashPass string) bool
	GenerateToken(userRole *entity.UserRoleEntityJoin) (*entity.TokenEntity, error)
}

func NewUtilAuth() *UtilAuthModule {
	return &UtilAuthModule{}

}
