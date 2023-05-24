package module

import (
	"context"
	entity "lelo-user/entity"

	authcontroller "lelo-user/controller/auth"
	locationcontroller "lelo-user/controller/location"
	usercontroller "lelo-user/controller/user"
	addressrepository "lelo-user/repository/address"
	locationrepository "lelo-user/repository/location"
	rolerepository "lelo-user/repository/role"
	userrepository "lelo-user/repository/user"
	userrolerepository "lelo-user/repository/userrole"
	addressusecase "lelo-user/usecase/address"
	locationusecase "lelo-user/usecase/location"
	userusecase "lelo-user/usecase/user"
	utilauth "lelo-user/util/auth"
	dbauth "lelo-user/util/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

// db
var UserRepository *userrepository.UserRepositoryModule
var RoleRepository *rolerepository.RoleRepositoryModule
var UserRoleRepository *userrolerepository.UserRoleRepositoryModule
var LocationRepository *locationrepository.LocationRepoistoryModule
var AddressRepository *addressrepository.AddressRepositoryModule

// usecase
var UserUsecaseModule *userusecase.UserUsecaseModule
var UtilAuthModule *utilauth.UtilAuthModule
var UtilDbModule *dbauth.UtilDBModule
var LocationUsecaseModule *locationusecase.LocationUsecaseModule
var AddressUsecaseModule *addressusecase.AddressUsecaseModule

// controller
var AuthModule *authcontroller.AuthControllerModule
var UserController *usercontroller.UserControllerModule
var LocationController *locationcontroller.LocationControllerModule

func InitModule(ctx context.Context, configEntity *entity.Config, credentialEntity *entity.Credential, db *pgxpool.Pool) error {

	// init util
	UtilAuthModule = utilauth.NewUtilAuth()
	UtilDbModule = dbauth.NewUtilDb(db)

	// init repo
	UserRepository = userrepository.NewUserRepository(*db)
	RoleRepository = rolerepository.NewRoleRepository(*db)
	UserRoleRepository = userrolerepository.NewUserRoleRepository(*db)
	LocationRepository = locationrepository.NewLocationRepository(*db)
	AddressRepository = addressrepository.NewAddressRepository(*db)
	// init usecase
	UserUsecaseModule = userusecase.NewUserusecase(UserRepository, UtilAuthModule, UserRoleRepository, UtilDbModule, RoleRepository)
	LocationUsecaseModule = locationusecase.NewLocationUsecase(LocationRepository)
	AddressUsecaseModule = addressusecase.NewAddressUsecaseModule(AddressRepository, UtilDbModule)
	// init controller
	AuthModule = authcontroller.NewAuthController(UserUsecaseModule)
	UserController = usercontroller.NewUserController(UserUsecaseModule, AddressUsecaseModule)
	LocationController = locationcontroller.NewLocationController(LocationUsecaseModule)
	// init server

	return nil
}
