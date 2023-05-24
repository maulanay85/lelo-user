package controller

import (
	"fmt"
	"strconv"

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

func (uc *UserControllerModule) GetUserDataById(c *gin.Context) {
	id, _ := c.Get("id")
	user, err := uc.UserUsecase.GetUserDataById(c, int64(id.(float64)))
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, user)
}

func (uc *UserControllerModule) GetUserAddress(c *gin.Context) {
	id, _ := c.Get("id")
	address, err := uc.AddressUsecase.GetAddressByUserId(c, int64(id.(float64)))
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, address)
}

func (uc *UserControllerModule) InsertUserAddress(c *gin.Context) {
	id, _ := c.Get("id")
	var data entity.UserAddressRequestEntity
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Errorf("[controller] error bind json on InsertUserAddress: %#v", err)
		wrap := fmt.Errorf("err: %w", util.ErrorErrorBadRequest)
		util.SendErrorResponse(c, wrap)
		return
	}
	id, err := uc.AddressUsecase.InsertAddressByUser(c, int64(id.(float64)), data)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, id)
}

func (uc *UserControllerModule) GetAddressByUserIdAndId(c *gin.Context) {
	addressId := c.Param("addressId")
	addressIdInt, _ := strconv.ParseInt(addressId, 10, 64)
	id, _ := c.Get("id")
	data, err := uc.AddressUsecase.GetAddressByUserIdAndId(c, int64(id.(float64)), addressIdInt)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, data)
}

func (uc *UserControllerModule) SetMainAddress(c *gin.Context) {
	addressId := c.Param("addressId")
	addressIdInt, _ := strconv.ParseInt(addressId, 10, 64)
	userId, _ := c.Get("id")

	err := uc.AddressUsecase.SetMainAddressTx(c, int64(userId.(float64)), addressIdInt)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, nil)
}
