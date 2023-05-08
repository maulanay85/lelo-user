package controller

import (
	"fmt"

	"lelo-user/entity"
	util "lelo-user/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (uc *UserControllerModule) ChangePassword(c *gin.Context) {
	var changePasswordEntity entity.ChangePasswordEntity
	if err := c.ShouldBindJSON(&changePasswordEntity); err != nil {
		log.Errorf("[controller] error bind json on ChangePassword: %#v", err)
		wrap := fmt.Errorf("err: %w", util.ErrorErrorBadRequest)
		util.SendErrorResponse(c, wrap)
		return
	}
	id, err := uc.UserUsecase.ChangePassword(c,
		changePasswordEntity.Email,
		changePasswordEntity.CurrPass,
		changePasswordEntity.NewPass)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, id)
}
