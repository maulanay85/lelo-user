package user

import (
	"context"
	entity "lelo-user/entity"
	roleRepository "lelo-user/repository/role"
	userRepository "lelo-user/repository/user"
	userRoleRepository "lelo-user/repository/userrole"
	utilAuth "lelo-user/util/auth"
	utilDB "lelo-user/util/db"
)

type UserUsecaseModule struct {
	Repo           userRepository.UserRepository
	UtilAuth       utilAuth.UtilAuth
	UserRuleRepo   userRoleRepository.UserRoleRepository
	UtilDbModule   utilDB.UtilDb
	RoleRepository roleRepository.RoleRepository
}

func NewUserusecase(
	repo userRepository.UserRepository,
	utilauth utilAuth.UtilAuth,
	userRoleRepo userRoleRepository.UserRoleRepository,
	utilDbModule utilDB.UtilDb,
	roleRepository roleRepository.RoleRepository,
) *UserUsecaseModule {
	return &UserUsecaseModule{
		Repo:           repo,
		UtilAuth:       utilauth,
		UserRuleRepo:   userRoleRepo,
		UtilDbModule:   utilDbModule,
		RoleRepository: roleRepository,
	}
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, fullname string, email string, pass string) (int64, error)
	Login(ctx context.Context, email string, pass string) (*entity.TokenEntity, error)
	ChangePassword(ctx context.Context, email, currPassword, newPassword string) (int64, error)
	RefreshToken(ctx context.Context, rt string) (*entity.TokenEntity, error)
}
