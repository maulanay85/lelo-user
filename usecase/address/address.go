package address

import (
	"context"
	"fmt"
	"lelo-user/entity"
	"lelo-user/util"

	utilcommon "lelo-user/util/common"

	log "github.com/sirupsen/logrus"
)

func (a *AddressUsecaseModule) GetAddressByUserId(ctx context.Context, userId int64) ([]entity.UserAddressResponseEntity, error) {
	log.Infof("GetAddressByUserId for id: %d", userId)
	addresses, err := a.Repo.GetAddressByUserId(ctx, userId)
	if err != nil {
		log.Errorf("[usecase] GetAddressByUserId err: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return []entity.UserAddressResponseEntity{}, errorWrap
	}
	return addresses, nil
}

func (a *AddressUsecaseModule) InsertAddressByUser(ctx context.Context, userId int64, data entity.UserAddressRequestEntity) (int64, error) {
	log.Infof("insert address for user: %d. data : %v", userId, data)
	var address entity.UserAddressEntity
	utilcommon.ConvertModel(&data, &address)
	existAddress, err := a.Repo.GetAddressByUserId(ctx, userId)
	if err != nil {
		log.Errorf("[usecase] InsertAddressByUser err: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	address.UserId = userId
	if len(existAddress) == 0 {
		address.IsMain = true
	}
	id, err := a.Repo.InsertAddressByUserId(ctx, userId, &address)
	if err != nil {
		log.Errorf("[usecase] InsertAddressByUser err: %v", err)
		errorWrap := fmt.Errorf("error: %w", util.ErrorInternalServer)
		return 0, errorWrap
	}
	return id, nil
}

func (a *AddressUsecaseModule) GetAddressByUserIdAndId(ctx context.Context, userId int64, id int64) (*entity.UserAddressResponseEntity, error) {
	log.Infof("get address for user: %d. id : %d", userId, id)

	address, err := a.Repo.GetAddressByUserIdAndId(ctx, userId, id)
	if err != nil {
		log.Errorf("[usecase] GetAddressByUserIdAndId err: %v", err)
		errorWrap := fmt.Errorf("data is not exist: %w", util.ErrorNotFound)
		return &entity.UserAddressResponseEntity{}, errorWrap
	}

	return address, nil

}

func (a *AddressUsecaseModule) SetMainAddressTx(ctx context.Context, userId, addressId int64) error {
	log.Infof("SetMainAddress for userId: %d to id: %d", userId, addressId)

	tx, err := a.UtilDbModule.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	mainId, err := a.Repo.GetMainAddressTx(ctx, tx, userId)
	if err != nil {
		log.Errorf("[usecase] SetMainAddress err: %v", err)
		errorWrap := fmt.Errorf("error update main address: %w", util.ErrorInternalServer)
		return errorWrap
	}

	err = a.Repo.RemoveMainAddressTx(ctx, tx, userId, mainId)
	if err != nil {
		log.Errorf("[usecase] SetMainAddress err: %v", err)
		errorWrap := fmt.Errorf("error update main address: %w", util.ErrorInternalServer)
		return errorWrap
	}
	err = a.Repo.SetMainAddressTx(ctx, tx, userId, addressId)
	if err != nil {
		log.Errorf("[usecase] SetMainAddress err: %v", err)
		errorWrap := fmt.Errorf("error update main address: %w", util.ErrorInternalServer)
		return errorWrap
	}
	return nil
}
