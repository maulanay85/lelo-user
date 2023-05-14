package entity

type UserAddressEntity struct {
	Id         int64   `json:"id" db:"id"`
	UserId     int64   `json:"user_id" db:"user_id"`
	ProvinceId int32   `json:"province_id" db:"province_id"`
	CityId     int32   `json:"city_id" db:"city_id"`
	DistrictId int32   `json:"district_id" db:"district_id"`
	VillageId  int32   `json:"village_id" db:"village_id"`
	ZipCode    string  `json:"zip_code" db:"zip_code"`
	Lat        float32 `json:"lat" db:"lat"`
	Long       float32 `json:"long" db:"long"`
	IsMain     bool    `json:"is_main" db:"is_main"`
	IsDeleted  bool    `json:"is_deleted" db:"is_deleted"`
	BaseEntity
}
