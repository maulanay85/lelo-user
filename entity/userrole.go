package entity

type UserRoleEntity struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	RoleId int32 `json:"roleId"`
	Status int   `json:"status"`
	BaseEntity
}

type UserRoleEntityJoin struct {
	Id          int64  `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	RoleId      int32  `json:"roleId"`
	RoleName    string `json:"roleName"`
	Code        string `json:"code"`
}
