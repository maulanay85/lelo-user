package entity

type RoleEntity struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"status"`
	BaseEntity
}
