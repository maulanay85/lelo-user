package user

import (
	"context"

	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func (u *UserRepositoryModule) CheckEmailExist(ctx context.Context, email string) (int, error) {
	var total int
	err := u.db.QueryRow(ctx, `
		SELECT 
			count(*) as total
		FROM t_mst_user where email = $1`, email).Scan(&total)
	if err != nil {
		log.Errorf("[repository]: CheckEmailExit err: %v", err)
		return 0, err
	}
	return total, nil
}

func (u *UserRepositoryModule) InsertUser(ctx context.Context, user *entity.UserEntity) (int32, error) {
	var id int32
	err := u.db.QueryRow(ctx,
		`INSERT INTO t_mst_user
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

func (u *UserRepositoryModule) InsertUserTx(ctx context.Context, tx pgx.Tx, user *entity.UserEntity) (int64, error) {
	var id int64

	err := tx.QueryRow(ctx,
		`INSERT INTO t_mst_user
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
		`UPDATE t_mst_user
			SET pass = $1
		WHERE email = $2
		`, pass, email)
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
			FROM t_mst_user
		WHERE email = $1
	`, email).Scan(&pass)
	if err != nil {
		log.Errorf("[repository]: GetPasswordByEmail err: %v", err)
		return "", err
	}
	return pass, err
}

func (u *UserRepositoryModule) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.db.QueryRow(ctx,
		`select
			id,
			fullname,
			email,
			pass
		from
			t_mst_user
		where email = $1
		and status = 1
		`, email,
	).Scan(&user.Id, &user.Fullname, &user.Email, &user.Pass)
	if err != nil {
		log.Errorf("[repository]: GetUserByEmail err: %v", err)
		return &user, err
	}

	return &user, nil
}

func (u *UserRepositoryModule) GetUserById(ctx context.Context, id int64) (*entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.db.QueryRow(ctx,
		`select
			id,
			fullname,
			email
		from
			t_mst_user
		where email = $1
		and status = 1
		`, id,
	).Scan(&user.Id, &user.Fullname, &user.Email)
	if err != nil {
		log.Errorf("[repository]: GetUserByEmail err: %v", err)
		return &user, err
	}

	return &user, nil
}
