package model

import (
	"encoding/json"
	"time"
)

type School struct {
	UserID         string          `gorm:"primaryKey" json:"user_id"`
	SchoolID       string          `json:"school_id"`
	School         json.RawMessage `gorm:"-" json:"school"`
	GraduationYear int             `json:"graduation_year"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (School) TableName() string {
	return "user_school"
}
