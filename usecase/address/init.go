package address

import (
	"context"
	"lelo-user/entity"
	addressRepository "lelo-user/repository/address"
)

type AddressUsecaseModule struct {
	Repo addressRepository.AddressRepository
}

func NewAddressUsecaseModule(
	repo addressRepository.AddressRepository,
) *AddressUsecaseModule {
	return &AddressUsecaseModule{
		Repo: repo,
	}
}

type AddressUsecase interface {
	GetAddressByUserId(ctx context.Context, userId int64) ([]entity.UserAddressResponseEntity, error)
	InsertAddressByUser(ctx context.Context, userId int64, data entity.UserAddressRequestEntity) (int64, error)
	GetAddressByUserIdAndId(ctx context.Context, userId int64, id int64) (*entity.UserAddressResponseEntity, error)
}
