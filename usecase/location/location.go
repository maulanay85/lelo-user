package location

import (
	"context"
	"fmt"
	"lelo-user/entity"
	"lelo-user/util"

	log "github.com/sirupsen/logrus"
)

func (l *LocationUsecaseModule) GetProvince(ctx context.Context, name string) ([]entity.ProvinceEntity, error) {
	locations, err := l.LocationRepository.GetProvince(ctx, name)
	if err != nil {
		log.Errorf("[usecase] GetProvince: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return []entity.ProvinceEntity{}, errorWrap
	}

	return locations, nil
}

func (l *LocationUsecaseModule) GetCityByProvinceId(ctx context.Context, provinceId int64, name string) ([]entity.CityEntity, error) {
	cities, err := l.LocationRepository.GetCityByProvinceId(ctx, provinceId, name)
	if err != nil {
		log.Errorf("[usecase] GetProvince: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return []entity.CityEntity{}, errorWrap
	}

	return cities, nil
}
