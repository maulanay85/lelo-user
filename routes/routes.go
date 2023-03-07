package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type routes struct {
	router *gin.Engine
	port   int32
}

func NewRoutes(router *gin.Engine, port int32) routes {
	r := routes{
		router: router,
		port:   port,
	}
	v1 := r.router.Group("/v1")

	r.addPing(v1)
	r.addAuth(v1)

	return r
}

func (r routes) Run() error {
	portStr := fmt.Sprintf(":%d", r.port)
	log.Infof("running at %s", portStr)
	return r.router.Run(portStr)
}
