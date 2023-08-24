package model

import "time"

type Connection struct {
	UserID     string    `gorm:"primaryKey" json:"user_id"`
	ID         string    `gorm:"primaryKey" json:"key"`
	Connection string    `json:"connection"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Connection) TableName() string {
	return "user_connection"
}
