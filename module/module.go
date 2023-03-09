package module

import (
	"context"
	entity "lelo-user/entity"

	authcontroller "lelo-user/controller/auth"
	rolerepository "lelo-user/repository/role"
	userrepository "lelo-user/repository/user"
	userrolerepository "lelo-user/repository/userrole"
	userusecase "lelo-user/usecase/user"
	utilauth "lelo-user/util/auth"

	"github.com/jackc/pgx/v4/pgxpool"
)

var UserRepository *userrepository.UserRepositoryModule
var RoleRepository *rolerepository.RoleRepositoryModule
var UserRoleRepository *userrolerepository.UserRoleRepositoryModule

var UserUsecaseModule *userusecase.UserUsecaseModule
var UtilAuthModule *utilauth.UtilAuthModule
var AuthModule *authcontroller.AuthControllerModule

func InitModule(ctx context.Context, configEntity *entity.Config, credentialEntity *entity.Credential, db *pgxpool.Pool) error {

	UtilAuthModule = utilauth.NewUtilAuth()

	// init repo
	UserRepository = userrepository.NewUserRepository(*db)
	RoleRepository = rolerepository.NewRoleRepository(*db)
	UserRoleRepository = userrolerepository.NewUserRoleRepository(*db)
	// init usecase
	UserUsecaseModule = userusecase.NewUserusecase(UserRepository, UtilAuthModule)

	// init controller
	AuthModule = authcontroller.NewAuthController(UserUsecaseModule)
	// init server

	return nil
}
