package location

import (
	"context"
	"lelo-user/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type LocationRepoistoryModule struct {
	db *pgxpool.Pool
}

func NewLocationRepository(
	db pgxpool.Pool,
) *LocationRepoistoryModule {
	return &LocationRepoistoryModule{
		db: &db,
	}
}

type LocationRepository interface {
	GetProvince(ctx context.Context, name string) ([]entity.ProvinceEntity, error)
	GetCityByProvinceId(ctx context.Context, provinceId int64, name string) ([]entity.CityEntity, error)
	GetDistrictByCityId(ctx context.Context, cityId int64, name string) ([]entity.DistrictEntity, error)
	GetVillageByDistrictId(ctx context.Context, districtId int64, name string) ([]entity.VillageEntity, error)
}
