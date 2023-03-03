package module

import (
	"context"
	"errors"
	dbModule "lelo-user/config"
	entity "lelo-user/entity"

	userrepository "lelo-user/repository/user"
	userusecase "lelo-user/usecase/user"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var UserRepository *userrepository.UserRepositoryModule
var UserusecaseModule *userusecase.UserUsecaseModule

func InitModule(ctx context.Context, configEntity *entity.Config, credentialEntity *entity.Credential, db *pgxpool.Pool) error {
	// init db
	db, err := dbModule.InitDb(ctx, configEntity, credentialEntity)
	if err != nil {
		log.Errorf("error init db: %#v", err)
		return errors.New(err.Error())
	}

	// init repo
	UserRepository = userrepository.NewUserRepository(*db)

	// init usecase
	UserusecaseModule = userusecase.NewUserusecase(UserRepository)

	return nil
}
