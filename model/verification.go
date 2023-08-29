package model

import (
	"time"
)

type Verification struct {
	UserID          string    `gorm:"primaryKey" json:"user_id"`
	Type            string    `json:"type"`
	FileURL         string    `json:"file_url"`
	Status          string    `json:"status"`
	Comments        string    `json:"comments"`
	IsVerified      bool      `json:"is_verified"`
	IsEmailVerified bool      `json:"is_email_verified"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Verification) TableName() string {
	return "user_verification"
}
