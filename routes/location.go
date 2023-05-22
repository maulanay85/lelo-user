package routes

import (
	"lelo-user/module"

	"github.com/gin-gonic/gin"
)

func (r routes) addLocation(rg *gin.RouterGroup) {
	authModule := module.UtilAuthModule
	location := rg.Group("/location")
	location.Use(authModule.JwtTokenCheck)
	{
		location.GET("/province", module.LocationController.GetProvince)
		location.GET("/province/:provinceId/city", module.LocationController.GetCity)
		location.GET("/city/:cityId/district", module.LocationController.GetDistrictByCityId)
	}
}
