package module

import (
	"context"
	entity "lelo-user/entity"

	userrepository "lelo-user/repository/user"
	userusecase "lelo-user/usecase/user"
	utilauth "lelo-user/util/auth"

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

	return nil
}
