package controller

import (
	addressusecae "lelo-user/usecase/address"
	userusecase "lelo-user/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserControllerModule struct {
	UserUsecase    userusecase.UserUsecase
	AddressUsecase addressusecae.AddressUsecase
}

func NewUserController(
	userusecase userusecase.UserUsecase,
	addressusecase addressusecae.AddressUsecase,
) *UserControllerModule {
	return &UserControllerModule{
		UserUsecase:    userusecase,
		AddressUsecase: addressusecase,
	}
}

type UserController interface {
	ChangePassword(c *gin.Context)
	GetUserDataById(c *gin.Context)
	GetUserAddress(c *gin.Context)
	InsertUserAddress(c *gin.Context)
}
