package entity

import (
	"time"
)

type UserEntity struct {
	Id          int64  `json:"id" db:"id"`
	Fullname    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Pass        string `json:"-"`
	Status      int    `json:"status"`
	Bod         time.Time
}
