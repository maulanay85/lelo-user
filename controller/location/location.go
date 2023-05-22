package location

import (
	"lelo-user/util"
	"strconv"

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

func (lc *LocationControllerModule) GetCity(c *gin.Context) {
	name := c.Query("name")
	provinceId := c.Param("provinceId")
	id, _ := strconv.ParseInt(provinceId, 10, 64)

	cities, err := lc.LocationUsecase.GetCityByProvinceId(c, id, name)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, cities)
}

func (lc *LocationControllerModule) GetDistrictByCityId(c *gin.Context) {
	name := c.Query("name")
	cityId := c.Param("cityId")
	id, _ := strconv.ParseInt(cityId, 10, 64)

	districts, err := lc.LocationUsecase.GetDistrictByCityId(c, id, name)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, districts)
}

func (lc *LocationControllerModule) GetVillageByDistrictId(c *gin.Context) {
	name := c.Query("name")
	districtId := c.Param("districtId")
	id, _ := strconv.ParseInt(districtId, 10, 64)

	districts, err := lc.LocationUsecase.GetVillageByDistrictId(c, id, name)
	if err != nil {
		util.SendErrorResponse(c, err)
		return
	}
	util.SendSuccess(c, districts)
}
