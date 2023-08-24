package model

import "time"

type Role struct {
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Role      string    `gorm:"primaryKey" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Role) TableName() string {
	return "user_role"
}
