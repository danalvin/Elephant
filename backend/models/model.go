package model

import "time"

// BaseModel -
type BaseModel struct {
	ID int64 `json:"id"`

	CreatedBy int64 `json:"created_by"`
	UpdatedBy int64 `json:"updated_by"`
	DeletedBy int64 `json:"deleted_by"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
