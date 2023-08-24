package model

import (
	"time"
)

type Verification struct {
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Type      string    `json:"type"`
	FileURL   string    `json:"file_url"`
	Status    string    `json:"status"`
	Comments  string    `json:"comments"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Verification) TableName() string {
	return "user_verification"
}
