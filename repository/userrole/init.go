package userrole

import "github.com/jackc/pgx/v4/pgxpool"

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
}
