package module

import (
	"context"
	entity "lelo-user/entity"

	userrepository "lelo-user/repository/user"
	routes "lelo-user/routes"
	userusecase "lelo-user/usecase/user"
	utilauth "lelo-user/util/auth"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

var UserRepository *userrepository.UserRepositoryModule
var UserusecaseModule *userusecase.UserUsecaseModule
var UtilAuthModule *utilauth.UtilAuthModule

func InitModule(ctx context.Context, configEntity *entity.Config, credentialEntity *entity.Credential, db *pgxpool.Pool) error {

	UtilAuthModule = utilauth.NewUtilAuth()

	// init repo
	UserRepository = userrepository.NewUserRepository(*db)

	// init usecase
	UserusecaseModule = userusecase.NewUserusecase(UserRepository, UtilAuthModule)

	// init server

	Routes := routes.NewRoutes(gin.New(), int32(configEntity.Port))
	Routes.Run()
	return nil
}
