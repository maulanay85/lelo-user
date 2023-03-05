package entity

import (
	"time"
)

type BaseEntity struct {
	CreatedBy   int32     `json:"created_by"`
	UpdatedBy   int32     `json:"updated_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"Updated_date"`
}
