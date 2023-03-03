package user

import (
	"context"
	"fmt"
	entity "lelo-user/entity"

	util "lelo-user/util"

	log "github.com/sirupsen/logrus"
)

func (u *UserUsecaseModule) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	userEntity, err := u.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Errorf("error on userusecase.getUserByEmail: %v", err)

		wrap := fmt.Errorf("error : %w", util.ErrorInternalServer)
		return nil, wrap
	}
	return userEntity, nil
}
