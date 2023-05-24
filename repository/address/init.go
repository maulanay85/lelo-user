package address

import (
	"context"
	"lelo-user/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AddressRepositoryModule struct {
	db *pgxpool.Pool
}

func NewAddressRepository(
	db pgxpool.Pool,
) *AddressRepositoryModule {
	return &AddressRepositoryModule{
		db: &db,
	}
}

type AddressRepository interface {
	GetAddressByUserId(ctx context.Context, userId int64) ([]entity.UserAddressResponseEntity, error)
	GetAddressByUserIdAndId(ctx context.Context, userId int64, id int64) (*entity.UserAddressResponseEntity, error)
	InsertAddressByUserId(ctx context.Context, userId int64, data *entity.UserAddressEntity) (int64, error)
	// UpdateAddressByUserIdAndId(ctx context.Context, userId int64, data *entity.UserAddressEntity) error
}
