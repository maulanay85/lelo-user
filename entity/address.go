package entity

type UserAddressEntity struct {
	Id int64 `json:"id" db:"id"`
	BaseEntity
}
