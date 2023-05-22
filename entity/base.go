package entity

import (
	"time"
)

type BaseEntity struct {
	CreatedBy   int32     `json:"created_by,omitempty"`
	UpdatedBy   int32     `json:"updated_by,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	UpdatedTime time.Time `json:"Updated_time,omitempty"`
}
