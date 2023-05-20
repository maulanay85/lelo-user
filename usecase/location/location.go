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
