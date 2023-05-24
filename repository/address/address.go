package address

import (
	"context"
	"lelo-user/entity"
	"time"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

func (ua *AddressRepositoryModule) GetAddressByUserId(ctx context.Context, userId int64) ([]entity.UserAddressResponseEntity, error) {
	var userAddresses []entity.UserAddressResponseEntity

	rows, err := ua.db.Query(ctx,
		`SELECT
			ua.id,
			ua.user_id,
			ua.province_id,
			ua.city_id,
			ua.district_id,
			ua.village_id,
			ua.zip_code,
			ua.lat,
			ua.long,
			ua.is_main,
			ua.created_by,
			ua.updated_by,
			ua.created_time,
			ua.updated_time,
			ua.address,
			ua.is_deleted,
			ua.notes,
			p.name,
			c.name,
			d.name,
			v.name
		FROM t_mst_user_address ua
		JOIN t_mst_province p on ua.province_id = p.id
		JOIN t_mst_city c on ua.city_id = c.id
		JOIN t_mst_district d on ua.district_id = d.id
		JOIN t_mst_village v on ua.village_id = v.id
		WHERE ua.user_id = $1 AND ua.is_deleted = false AND ua.is_main = true
		UNION ALL
		SELECT
			ua.id,
			ua.user_id,
			ua.province_id,
			ua.city_id,
			ua.district_id,
			ua.village_id,
			ua.zip_code,
			ua.lat,
			ua.long,
			ua.is_main,
			ua.created_by,
			ua.updated_by,
			ua.created_time,
			ua.updated_time,
			ua.address,
			ua.is_deleted,
			ua.notes,
			p.name,
			c.name,
			d.name,
			v.name
		FROM t_mst_user_address ua
		JOIN t_mst_province p on ua.province_id = p.id
		JOIN t_mst_city c on ua.city_id = c.id
		JOIN t_mst_district d on ua.district_id = d.id
		JOIN t_mst_village v on ua.village_id = v.id
		WHERE ua.user_id = $2 AND ua.is_deleted = false AND ua.is_main = false
		`, userId, userId,
	)
	if err != nil {
		log.Errorf("[repository] GetAddressByUserId: err %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		address := entity.UserAddressResponseEntity{}
		err := rows.Scan(&address.Id,
			&address.UserId,
			&address.ProvinceId,
			&address.CityId,
			&address.DistrictId,
			&address.VillageId,
			&address.ZipCode,
			&address.Lat,
			&address.Long,
			&address.IsMain,
			&address.CreatedBy,
			&address.UpdatedBy,
			&address.CreatedTime,
			&address.UpdatedTime,
			&address.Address,
			&address.IsDeleted,
			&address.Notes,
			&address.ProvinceName,
			&address.CityName,
			&address.DistrcitName,
			&address.VillageName)
		if err != nil {
			log.Errorf("[repository] GetAddressByUserId: err %v", err)
			return nil, err
		}
		userAddresses = append(userAddresses, address)
	}
	return userAddresses, nil
}

func (ua *AddressRepositoryModule) GetAddressByUserIdAndId(ctx context.Context, userId int64, id int64) (*entity.UserAddressResponseEntity, error) {
	address := entity.UserAddressResponseEntity{}
	sql :=
		`SELECT
		ua.id,
		ua.user_id,
		ua.province_id,
		ua.city_id,
		ua.district_id,
		ua.village_id,
		ua.zip_code,
		ua.lat,
		ua.long,
		ua.is_main,
		ua.created_by,
		ua.updated_by,
		ua.created_time,
		ua.updated_time,
		ua.address,
		ua.is_deleted,
		ua.notes,
		p.name,
		c.name,
		d.name,
		v.name
	FROM t_mst_user_address ua
	JOIN t_mst_province p on ua.province_id = p.id
	JOIN t_mst_city c on ua.city_id = c.id
	JOIN t_mst_district d on ua.district_id = d.id
	JOIN t_mst_village v on ua.village_id = v.id
	WHERE ua.user_id = $1 AND ua.id = $2 AND ua.is_deleted = false`

	err := ua.db.QueryRow(ctx, sql, userId, id).Scan(
		&address.Id,
		&address.UserId,
		&address.ProvinceId,
		&address.CityId,
		&address.DistrictId,
		&address.VillageId,
		&address.ZipCode,
		&address.Lat,
		&address.Long,
		&address.IsMain,
		&address.CreatedBy,
		&address.UpdatedBy,
		&address.CreatedTime,
		&address.UpdatedTime,
		&address.Address,
		&address.IsDeleted,
		&address.Notes,
		&address.ProvinceName,
		&address.CityName,
		&address.DistrcitName,
		&address.VillageName,
	)
	if err != nil {
		log.Errorf("[repository] GetAddressByUserIdAndId userId: %d, id: %d, err: %v", userId, id, err)
		return nil, err
	}
	return &address, nil
}

func (ua *AddressRepositoryModule) InsertAddressByUserId(ctx context.Context, userId int64, data *entity.UserAddressEntity) (int64, error) {
	var id int64
	err := ua.db.QueryRow(ctx,
		`INSERT INTO t_mst_user_address
		(user_id, province_id, city_id, district_id, village_id, zip_code, lat, long, is_main, created_by, updated_by, address, is_deleted, notes)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
	RETURNING id`, data.UserId, data.ProvinceId, data.CityId, data.DistrictId, data.VillageId, data.ZipCode, data.Lat, data.Long, data.IsMain, userId, userId, data.Address, false, data.Notes,
	).Scan(&id)
	if err != nil {
		log.Errorf("[repository]: InsertAddressByUserId for userId %d error: %v", userId, err)
		return 0, err
	}
	return id, nil
}

func (ua *AddressRepositoryModule) GetMainAddressTx(ctx context.Context, tx pgx.Tx, userId int64) (int64, error) {
	var id int64
	err := tx.QueryRow(ctx,
		`
		SELECT id
		FROM t_mst_user_address
		WHERE is_main = true and user_id = $1
	`, userId).Scan(&id)
	if err != nil {
		log.Errorf("[repository]: GetMainAddressTx for userId %d error: %v", userId, err)
		return 0, err
	}
	return id, nil
}

func (ua *AddressRepositoryModule) RemoveMainAddressTx(ctx context.Context, tx pgx.Tx, userId, addressId int64) error {
	_, err := tx.Exec(ctx,
		`
		UPDATE t_mst_user_address SET is_main = false, updated_time = $1 WHERE id = $2 AND user_id = $3
	`, time.Now(), addressId, userId)
	if err != nil {
		log.Errorf("[repository]: RemoveMainAddressTx for id %d error: %v", addressId, err)
		return err
	}
	return nil
}

func (ua *AddressRepositoryModule) SetMainAddressTx(ctx context.Context, tx pgx.Tx, userId, addressId int64) error {
	_, err := tx.Exec(ctx,
		`
			UPDATE t_mst_user_address SET is_main = true, updated_time = $1 WHERE id = $2 AND user_id = $3
		`, time.Now(), addressId, userId)
	if err != nil {
		log.Errorf("[repository]: SetMainAddressTx for id %d error: %v", addressId, err)
		return err
	}
	return nil
}
