package location

import (
	"context"
	"lelo-user/entity"
	locationRepository "lelo-user/repository/location"
)

type LocationUsecaseModule struct {
	LocationRepository locationRepository.LocationRepository
}

func NewLocationUsecase(
	repo locationRepository.LocationRepository,
) *LocationUsecaseModule {
	return &LocationUsecaseModule{
		LocationRepository: repo,
	}
}

type LocationUsecase interface {
	GetProvince(ctx context.Context, name string) ([]entity.ProvinceEntity, error)
	GetCityByProvinceId(ctx context.Context, provinceId int64, name string) ([]entity.CityEntity, error)
	GetDistrictByCityId(ctx context.Context, cityId int64, name string) ([]entity.DistrictEntity, error)
	GetVillageByDistrictId(ctx context.Context, cityId int64, name string) ([]entity.VillageEntity, error)
}
