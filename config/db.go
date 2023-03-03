package config

import (
	"context"
	"fmt"
	entity "lelo-user/entity"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func InitDb(ctx context.Context, config *entity.Config, credential *entity.Credential) (*pgxpool.Pool, error) {
	urlDatabase := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		credential.Database.User,
		credential.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)
	db, err := pgxpool.Connect(ctx, urlDatabase)

	if err != nil {
		log.Errorf("error open connection: %v", err)
		return nil, err
	}
	log.Info("success connect to db!")
	return db, nil
}
