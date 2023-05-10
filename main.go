package main

import (
	"context"
	"fmt"
	"io"
	"os"

	configuration "lelo-user/config"
	"lelo-user/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	module "lelo-user/module"
)

func main() {

	// handle panic
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("panic happen: %#v", err)
		}
	}()

	// Read All Config
	err := configuration.ReadConfiguration(".")

	if err != nil {
		log.Errorf("error read config file: %#v", err)
		return
	}

	logname := "lelo-user.log"
	f, err := os.OpenFile(logname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error create file: %s", logname)
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if configuration.ConfigData.Env == "dev" {
		log.SetLevel(log.TraceLevel)
	}

	gin.SetMode(gin.ReleaseMode)

	// initiate context
	ctx := context.Background()

	// init db
	db, err := configuration.InitDb(ctx, &configuration.ConfigData, &configuration.CredentialData)
	if err != nil {
		log.Errorf("error init db: %#v", err)
		return
	}

	defer db.Close()

	err = module.InitModule(ctx, &configuration.ConfigData, &configuration.CredentialData, db)
	if err != nil {
		log.Errorf("error init modules: %#v", err)
		return
	}

	if err = routes.NewRoutes(gin.New(), int32(configuration.ConfigData.Port)).Run(); err != nil {
		log.Errorf("error running server : %#v", err)
	}

}
