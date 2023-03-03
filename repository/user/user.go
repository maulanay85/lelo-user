package user

import (
	"context"

	entity "lelo-user/entity"

	log "github.com/sirupsen/logrus"
)

func (u *UserRepositoryModule) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var userEntity entity.UserEntity
	err := u.db.QueryRow(ctx,
		`select 
			id, 
			fullName, 
			email,
			phone_number,
			bod, 
			status 
			from user_management where email = $1`, email).
		Scan(&userEntity.Id, &userEntity.Fullname, &userEntity.Email, &userEntity.PhoneNumber, &userEntity.Bod, &userEntity.Status)
	if err != nil {
		log.Errorf("error GetUserByEmail: %v", err)
		return nil, err
	}

	return &userEntity, nil
}
