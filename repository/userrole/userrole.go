package userrole

import (
	"context"
	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func (u *UserRoleRepositoryModule) GetUserRoleByUserId(ctx context.Context, userId int64) (userRole *entity.UserRoleEntityJoin, err error) {
	var userData entity.UserRoleEntityJoin

	err = u.db.QueryRow(ctx,
		`SELECT 
			u.id	id,
			r.code	code,
			u.email	email,
			u.fullname	fullname,
			u.phone_number	phoneNumber,
			ur.role_id	roleId,
			r.name	roleName
		FROM t_mst_user u
		LEFT JOIN t_map_user_role ur ON u.id = ur.user_id
		LEFT JOIN t_mst_role r on r.id = ur.role_id
		WHERE u.status = 1
		AND u.id = $1
	`, userId).Scan(
		&userData.Id,
		&userData.Code,
		&userData.Email,
		&userData.Fullname,
		&userData.PhoneNumber,
		&userData.RoleId,
		&userData.RoleName,
	)
	if err != nil {
		log.Errorf("[repository]: GetUserRoleByUserId err: %v", err)
		return &userData, err
	}
	return &userData, nil
}

func (u *UserRoleRepositoryModule) InsertUserRoleTx(ctx context.Context, tx pgx.Tx, userRole *entity.UserRoleEntity) (int64, error) {
	var id int64
	err := tx.QueryRow(ctx,
		`INSERT INTO t_map_user_role
			(user_id, role_id, status)
		VALUES($1, $2, 1) RETURNING id
		`, userRole.UserId, userRole.RoleId,
	).Scan(&id)
	if err != nil {
		log.Errorf("[repository]: InsertUserRoleTx err: %v", err)
		return 0, err
	}
	return id, nil
}
