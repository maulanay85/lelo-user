package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	entity "lelo-user/entity"
)

type UserRepositoryModule struct {
	db *pgxpool.Pool
}

func NewUserRepository(db pgxpool.Pool) *UserRepositoryModule {
	return &UserRepositoryModule{
		db: &db,
	}
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	// GetUserByPhone(ctx context.Context, phoneNumber string) (pgx.Row, error)
	// InsertUser(ctx context.Context, user *entity.UserEntity) (int32, error)
}
