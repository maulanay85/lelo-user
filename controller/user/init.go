package controller

import (
	userusecase "lelo-user/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserControllerModule struct {
	UserUsecase userusecase.UserUsecase
}

func NewUserController(
	userusecase userusecase.UserUsecase,
) *UserControllerModule {
	return &UserControllerModule{
		UserUsecase: userusecase,
	}
}

type UserController interface {
	ChangePassword(c *gin.Context)
	GetUserDataById(c *gin.Context)
}
