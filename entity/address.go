package entity

type UserAddressEntity struct {
	Id         int64   `json:"id" db:"id"`
	UserId     int64   `json:"user_id" db:"user_id"`
	ProvinceId int64   `json:"province_id" db:"province_id"`
	CityId     int64   `json:"city_id" db:"city_id"`
	DistrictId int64   `json:"district_id" db:"district_id"`
	VillageId  int64   `json:"village_id" db:"village_id"`
	ZipCode    string  `json:"zip_code" db:"zip_code"`
	Lat        float32 `json:"lat" db:"lat"`
	Long       float32 `json:"long" db:"long"`
	IsMain     bool    `json:"is_main" db:"is_main"`
	IsDeleted  bool    `json:"is_deleted" db:"is_deleted"`
	Address    string  `json:"address"`
	Notes      string  `json:"notes"`
	BaseEntity
}

type UserAddressResponseEntity struct {
	Id           int64   `json:"id" db:"id"`
	UserId       int64   `json:"user_id" db:"user_id"`
	ProvinceId   int64   `json:"province_id" db:"province_id"`
	CityId       int64   `json:"city_id" db:"city_id"`
	DistrictId   int64   `json:"district_id" db:"district_id"`
	VillageId    int64   `json:"village_id" db:"village_id"`
	ZipCode      string  `json:"zip_code" db:"zip_code"`
	Lat          float32 `json:"lat" db:"lat"`
	Long         float32 `json:"long" db:"long"`
	IsMain       bool    `json:"is_main" db:"is_main"`
	IsDeleted    bool    `json:"is_deleted" db:"is_deleted"`
	ProvinceName string  `json:"province_name"`
	CityName     string  `json:"city_name"`
	DistrcitName string  `json:"district_name"`
	VillageName  string  `json:"village_name"`
	Address      string  `json:"address"`
	Notes        string  `json:"notes"`
	BaseEntity
}

type UserAddressRequestEntity struct {
	ProvinceId int64   `json:"province_id" db:"province_id" binding:"required"`
	CityId     int64   `json:"city_id" db:"city_id" binding:"required"`
	DistrictId int64   `json:"district_id" db:"district_id" binding:"required"`
	VillageId  int64   `json:"village_id" db:"village_id" binding:"required"`
	ZipCode    string  `json:"zip_code" db:"zip_code" binding:"required"`
	Lat        float32 `json:"lat" db:"lat"`
	Long       float32 `json:"long" db:"long"`
	Address    string  `json:"address" binding:"required"`
	Notes      string  `json:"notes"`
}
