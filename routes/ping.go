package routes

import (
	pingcontroller "lelo-user/controller"

	"github.com/gin-gonic/gin"
)

func (r routes) addPing(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", pingcontroller.PongFunction)
}
