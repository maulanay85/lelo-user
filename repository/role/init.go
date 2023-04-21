package role

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	entity "lelo-user/entity"
)

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
	GetRoleByCode(ctx context.Context, tx pgx.Tx, code string) (*entity.RoleEntity, error)
}
