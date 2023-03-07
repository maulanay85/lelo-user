package main

import (
	"context"
	"os"

	config "lelo-user/config"
	"lelo-user/routes"

	dbModule "lelo-user/config"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	module "lelo-user/module"
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.TraceLevel)

	// handle panic
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("panic happen: %#v", err)
		}
	}()

	// Read All Config
	err := config.ReadConfiguration()

	if err != nil {
		log.Errorf("error read config file: %#v", err)
		return
	}
	gin.SetMode(gin.ReleaseMode)

	// initiate context
	ctx := context.Background()

	// init db
	db, err := dbModule.InitDb(ctx, &config.ConfigData, &config.CredentialData)
	if err != nil {
		log.Errorf("error init db: %#v", err)
		return
	}

	defer db.Close()

	err = module.InitModule(ctx, &config.ConfigData, &config.CredentialData, db)
	if err != nil {
		log.Errorf("error init modules: %#v", err)
		return
	}

	if err = routes.NewRoutes(gin.New(), int32(config.ConfigData.Port)).Run(); err != nil {
		log.Errorf("error running server : %#v", err)
	}

}
