package user

import (
	"context"

	entity "lelo-user/entity"

	log "github.com/sirupsen/logrus"
)

func (u *UserRepositoryModule) CheckEmailExist(ctx context.Context, email string) (int, error) {
	var total int
	err := u.db.QueryRow(ctx, `
		SELECT 
			count(*) as total
		FROM user_management where email = $1`, email).Scan(&total)
	if err != nil {
		log.Errorf("[repository]: CheckEmailExit err: %v", err)
		return 0, err
	}
	return total, nil
}

func (u *UserRepositoryModule) InsertUser(ctx context.Context, user *entity.UserEntity) (int32, error) {
	var id int32
	err := u.db.QueryRow(ctx,
		`INSERT INTO user_management
			(fullname, email, pass)
		 VALUES
		 	($1, $2, $3) Returning id
		`, user.Fullname, user.Email, user.Pass,
	).Scan(&id)
	if err != nil {
		log.Errorf("[repository]: InsertUser err: %v", err)
		return 0, err
	}
	return id, nil
}

func (u *UserRepositoryModule) ChangePassword(ctx context.Context, email string, pass string) error {
	_, err := u.db.Exec(ctx,
		`UPDATE user_management
			SET pass = $1
		WHERE email = $2
		`, email, pass)
	if err != nil {
		log.Errorf("[repository]: ChangePassword err: %v", err)
		return err
	}
	return nil
}

func (u *UserRepositoryModule) GetPasswordByEmail(ctx context.Context, email string) (string, error) {
	var pass string
	err := u.db.QueryRow(ctx, `
		SELECT pass
			FROM user_management
		WHERE email = $1
	`, email).Scan(&pass)
	if err != nil {
		log.Errorf("[repository]: GetPasswordByEmail err: %v", err)
		return "", err
	}
	return pass, err
}
