package role

import "github.com/jackc/pgx/v4/pgxpool"

type RoleRepositoryModule struct {
	db *pgxpool.Pool
}

func NewRoleRepository(
	db pgxpool.Pool,
) *RoleRepositoryModule {
	return &RoleRepositoryModule{
		db: &db,
	}
}

type RoleRepository interface {
}
