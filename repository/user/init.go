package user

import (
	"context"
	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4"
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
	InsertUserTx(ctx context.Context, tx pgx.Tx, user *entity.UserEntity) (int64, error)
	ChangePassword(ctx context.Context, email string, pass string) error
	GetPasswordByEmail(ctx context.Context, email string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	GetUserById(ctx context.Context, id int64) (*entity.UserEntity, error)
}
