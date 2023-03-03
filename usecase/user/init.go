package user

import (
	"context"
	entity "lelo-user/entity"
	userRepository "lelo-user/repository/user"
)

type UserUsecaseModule struct {
	Repo userRepository.UserRepository
}

func NewUserusecase(repo userRepository.UserRepository) *UserUsecaseModule {
	return &UserUsecaseModule{
		Repo: repo,
	}
}

type UserUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}
