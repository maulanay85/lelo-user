package controller

import (
	"fmt"
	"lelo-user/entity"

	util "lelo-user/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (uc *AuthControllerModule) RegisterUser(c *gin.Context) {
	var registerData entity.RegisterUserEntity
	if err := c.ShouldBindJSON(&registerData); err != nil {
		log.Errorf("[controller] error bind json on register user: %#v", err)
		wrap := fmt.Errorf("err: %w", util.ErrorErrorBadRequest)
		util.SendErrorResponse(c, wrap)
		return
	}
	id, err := uc.UserUsecase.RegisterUser(c, registerData.Fullname, registerData.Email, registerData.Pass)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, id)

}

func (uc *AuthControllerModule) LoginUser(c *gin.Context) {
	var loginData entity.LoginEntity
	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Errorf("[controller] error bind json on login: %#v", err)
		wrap := fmt.Errorf("err: %w", util.ErrorErrorBadRequest)
		util.SendErrorResponse(c, wrap)
		return
	}
	token, err := uc.UserUsecase.Login(c, loginData.Email, loginData.Pass)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}

	util.SendSuccess(c, token)
}
