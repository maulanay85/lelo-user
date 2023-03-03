package module

import (
	"context"
	"errors"
	dbModule "lelo-user/config"
	configEntity "lelo-user/entity/config"
	userrepoitory "lelo-user/repository/user"

	log "github.com/sirupsen/logrus"
)

func InitModule(ctx context.Context, configEntity *configEntity.Config, credentialEntity *configEntity.Credential) error {
	// init db
	db, err := dbModule.InitDb(ctx, configEntity, credentialEntity)
	if err != nil {
		log.Errorf("error init db: %#v", err)
		return errors.New(err.Error())
	}

	// init usecase
	userrepoitory.NewUserRepository(*db)
	return nil
}
