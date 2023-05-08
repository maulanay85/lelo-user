package entity

type UserEntity struct {
	Id          int64  `json:"id" db:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Pass        string `json:"-"`
	Status      int    `json:"status"`
	BaseEntity
}

type RegisterUserEntity struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Pass     string `json:"pass" binding:"required"`
}

type LoginEntity struct {
	Email string `json:"email" binding:"required,email"`
	Pass  string `json:"pass" binding:"required"`
}

type ChangePasswordEntity struct {
	Email    string `json:"email" binding:"required,email"`
	CurrPass string `json:"currPass" binding:"required"`
	NewPass  string `json:"newPass" binding:"required"`
}
