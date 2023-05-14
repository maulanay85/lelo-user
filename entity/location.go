package entity

type VillageEntity struct {
	Id         int32  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	DistrictId int32  `json:"district_id" db:"district_id"`
	IsDeleted  bool   `json:"is_deleted" db:"is_deleted"`
	BaseEntity
}

type DistrictEntity struct {
	Id        int32  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CityId    int32  `json:"city_id" db:"city_id"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted"`
	BaseEntity
}

type CityEntity struct {
	Id         int32  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProvinceId int32  `json:"province_id" db:"province_id"`
	IsDeleted  bool   `json:"is_deleted" db:"is_deleted"`
	BaseEntity
}

type ProvinceEntity struct {
	Id        int32  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted"`
	BaseEntity
}
