package entity

type UserRoleEntity struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	RoleId int32 `json:"roleId"`
	Status int   `json:"status"`
	BaseEntity
}
