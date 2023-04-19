package user

import (
	"context"
	entity "lelo-user/entity"
	userRepository "lelo-user/repository/user"
	userRoleRepository "lelo-user/repository/userrole"
	utilAuth "lelo-user/util/auth"
)

type UserUsecaseModule struct {
	Repo         userRepository.UserRepository
	UtilAuth     utilAuth.UtilAuth
	UserRuleRepo userRoleRepository.UserRoleRepository
}

func NewUserusecase(
	repo userRepository.UserRepository,
	utilauth utilAuth.UtilAuth,
	userRoleRepo userRoleRepository.UserRoleRepository,
) *UserUsecaseModule {
	return &UserUsecaseModule{
		Repo:         repo,
		UtilAuth:     utilauth,
		UserRuleRepo: userRoleRepo,
	}
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, fullname string, email string, pass string) (int32, error)
	Login(ctx context.Context, email string, pass string) (*entity.TokenEntity, error)
}
