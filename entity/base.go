package entity

import (
	"time"
)

type BaseEntity struct {
	CreatedBy   int32     `json:"created_by"`
	UpdatedBy   int32     `json:"updated_by"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"Updated_time"`
}
