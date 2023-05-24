package address

import (
	"context"
	"lelo-user/entity"
	addressRepository "lelo-user/repository/address"
	utilDB "lelo-user/util/db"
)

type AddressUsecaseModule struct {
	Repo         addressRepository.AddressRepository
	UtilDbModule utilDB.UtilDb
}

func NewAddressUsecaseModule(
	repo addressRepository.AddressRepository,
	dbutil utilDB.UtilDb,
) *AddressUsecaseModule {
	return &AddressUsecaseModule{
		Repo:         repo,
		UtilDbModule: dbutil,
	}
}

type AddressUsecase interface {
	GetAddressByUserId(ctx context.Context, userId int64) ([]entity.UserAddressResponseEntity, error)
	InsertAddressByUser(ctx context.Context, userId int64, data entity.UserAddressRequestEntity) (int64, error)
	GetAddressByUserIdAndId(ctx context.Context, userId int64, id int64) (*entity.UserAddressResponseEntity, error)
	SetMainAddressTx(ctx context.Context, userId, addressId int64) error
}
