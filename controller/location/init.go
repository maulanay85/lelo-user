package location

import (
	"github.com/gin-gonic/gin"

	locationusecase "lelo-user/usecase/location"
)

type LocationControllerModule struct {
	LocationUsecase locationusecase.LocationUsecase
}

func NewLocationController(
	locationusecase locationusecase.LocationUsecase,
) *LocationControllerModule {
	return &LocationControllerModule{
		LocationUsecase: locationusecase,
	}
}

type LocationController interface {
	GetProvince(c *gin.Context)
	GetCity(c *gin.Context)
}
