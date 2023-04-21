package userrole

import (
	"context"

	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRoleRepositoryModule struct {
	db *pgxpool.Pool
}

func NewUserRoleRepository(
	db pgxpool.Pool,
) *UserRoleRepositoryModule {
	return &UserRoleRepositoryModule{
		db: &db,
	}
}

type UserRoleRepository interface {
	GetUserRoleByUserId(ctx context.Context, userId int64) (userRole *entity.UserRoleEntityJoin, err error)
	InsertUserRoleTx(ctx context.Context, tx pgx.Tx, userRole *entity.UserRoleEntity) (int64, error)
}
