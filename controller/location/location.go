package location

import (
	"lelo-user/util"

	"github.com/gin-gonic/gin"
)

func (lc *LocationControllerModule) GetProvince(c *gin.Context) {
	name := c.Query("name")

	provinces, err := lc.LocationUsecase.GetProvince(c, name)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, provinces)
}
