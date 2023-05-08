package routes

import (
	"lelo-user/module"

	"github.com/gin-gonic/gin"
)

func (r routes) addUser(rg *gin.RouterGroup) {
	authModule := module.UtilAuthModule

	user := rg.Group("/user")
	user.Use(authModule.JwtTokenCheck)
	{
		user.PATCH("/password", module.UserController.ChangePassword)
	}
}
