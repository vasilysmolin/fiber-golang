package models

import (
    "gorm.io/gorm"
)

type Profile struct {
	ID  uint64  `gorm:"primary_key;column:id" json:"id"`
	UserID   int64 `gorm:"column:user_id;uniqueIndex:idx_user_id,unique" json:"user_id"`
    UpdatedAt time.Time `gorm:"autoCreateTime:nano;autoUpdateTime:nano" json:"updated_at,omitempty"`
    CreatedAt time.Time `gorm:"autoCreateTime:nano;" json:"created_at,omitempty"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
}



