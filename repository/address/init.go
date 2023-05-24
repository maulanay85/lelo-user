package address

import (
	"context"
	"lelo-user/entity"

	"github.com/jackc/pgx/v4"
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
	GetMainAddressTx(ctx context.Context, tx pgx.Tx, userId int64) (int64, error)
	RemoveMainAddressTx(ctx context.Context, tx pgx.Tx, userId, addressId int64) error
	SetMainAddressTx(ctx context.Context, tx pgx.Tx, userId, addressId int64) error
	// UpdateAddressByUserIdAndId(ctx context.Context, userId int64, data *entity.UserAddressEntity) error
}
