package user

import (
	"context"
	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4/pgxpool"
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
	CheckEmailExist(ctx context.Context, email string) (int, error)
	InsertUser(ctx context.Context, user *entity.UserEntity) (int32, error)
	ChangePassword(ctx context.Context, email string, pass string) error
	GetPasswordByEmail(ctx context.Context, email string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}
