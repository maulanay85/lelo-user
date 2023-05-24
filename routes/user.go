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
		user.GET("/self", module.UserController.GetUserDataById)
		user.GET("/address", module.UserController.GetUserAddress)
		user.POST("/address", module.UserController.InsertUserAddress)
		user.GET("/address/:addressId", module.UserController.GetAddressByUserIdAndId)
		user.PATCH("/address/:addressId/main", module.UserController.SetMainAddress)
	}
}
