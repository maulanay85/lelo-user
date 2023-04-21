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
	dbauth "lelo-user/util/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

var UserRepository *userrepository.UserRepositoryModule
var RoleRepository *rolerepository.RoleRepositoryModule
var UserRoleRepository *userrolerepository.UserRoleRepositoryModule

var UserUsecaseModule *userusecase.UserUsecaseModule
var UtilAuthModule *utilauth.UtilAuthModule
var AuthModule *authcontroller.AuthControllerModule
var UtilDbModule *dbauth.UtilDBModule

func InitModule(ctx context.Context, configEntity *entity.Config, credentialEntity *entity.Credential, db *pgxpool.Pool) error {

	// init util
	UtilAuthModule = utilauth.NewUtilAuth()
	UtilDbModule = dbauth.NewUtilDb(db)

	// init repo
	UserRepository = userrepository.NewUserRepository(*db)
	RoleRepository = rolerepository.NewRoleRepository(*db)
	UserRoleRepository = userrolerepository.NewUserRoleRepository(*db)
	// init usecase
	UserUsecaseModule = userusecase.NewUserusecase(UserRepository, UtilAuthModule, UserRoleRepository, UtilDbModule)

	// init controller
	AuthModule = authcontroller.NewAuthController(UserUsecaseModule)
	// init server

	return nil
}
