package user

import (
	"context"
	userRepository "lelo-user/repository/user"
	utilAuth "lelo-user/util/auth"
)

type UserUsecaseModule struct {
	Repo     userRepository.UserRepository
	UtilAuth utilAuth.UtilAuth
}

func NewUserusecase(
	repo userRepository.UserRepository,
	utilauth utilAuth.UtilAuth,
) *UserUsecaseModule {
	return &UserUsecaseModule{
		Repo:     repo,
		UtilAuth: utilauth,
	}
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, fullname string, email string, pass string) (int32, error)
}
