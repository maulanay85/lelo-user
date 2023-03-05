package user

import (
	"context"
	"fmt"

	entity "lelo-user/entity"
	util "lelo-user/util"

	log "github.com/sirupsen/logrus"
)

func (u *UserUsecaseModule) RegisterUser(ctx context.Context, fullname string, email string, pass string) (int32, error) {
	id, err := u.Repo.CheckEmailExist(ctx, email)
	if err != nil {
		log.Errorf("[usecase] RegisterUser: %v", err)
		errorWrap := fmt.Errorf("error CheckEmailExist: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	if id != 0 {
		log.Error("[usecase] RegisterUser: email is exist")
		errorWrap := fmt.Errorf("email is exist: %w", util.ErrorPreCondition)
		return 0, errorWrap
	}
	// hash password
	hash, err := u.UtilAuth.HashPassword(pass)
	if err != nil {
		log.Errorf("[usecase] RegisterUser: %v", err)
		errorWrap := fmt.Errorf("error HashPassword: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	// insert user
	user := entity.UserEntity{
		Email:    email,
		Fullname: fullname,
		Pass:     hash,
	}
	iduser, err := u.Repo.InsertUser(ctx, &user)
	if err != nil {
		log.Errorf("[usecase] RegisterUser: %v", err)
		errorWrap := fmt.Errorf("error InsertUser: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	return iduser, nil
}
