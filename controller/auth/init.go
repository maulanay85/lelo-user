package controller

import (
	userusecase "lelo-user/usecase/user"

	"github.com/gin-gonic/gin"
)

type AuthControllerModule struct {
	UserUsecase userusecase.UserUsecase
}

func NewAuthController(
	userusecase userusecase.UserUsecase,
) *AuthControllerModule {
	return &AuthControllerModule{
		UserUsecase: userusecase,
	}
}

type AuthController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	RefreshToken(c *gin.Context)
}
