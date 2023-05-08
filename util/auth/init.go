package auth

import (
	entity "lelo-user/entity"

	"github.com/gin-gonic/gin"
)

type UtilAuthModule struct{}

type UtilAuth interface {
	HashPassword(string) (string, error)
	CheckHashPassword(pass string, hashPass string) bool
	GenerateToken(userRole *entity.UserRoleEntityJoin) (*entity.TokenEntity, error)
	JwtTokenCheck(c *gin.Context)
}

func NewUtilAuth() *UtilAuthModule {
	return &UtilAuthModule{}

}
