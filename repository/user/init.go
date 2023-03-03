package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db pgx.Conn) *UserRepository {
	return &UserRepository{
		db: &db,
	}
}

type IUserRepository interface {
	GetUserById(ctx context.Context)
}
