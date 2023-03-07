package routes

import (
	module "lelo-user/module"

	"github.com/gin-gonic/gin"
)

func (r routes) addAuth(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.POST("/register", module.AuthModule.RegisterUser)
}
