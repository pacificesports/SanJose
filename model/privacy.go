package model

import "time"

type Privacy struct {
	UserID                   string    `gorm:"primaryKey" json:"user_id"`
	ShowEmail                bool      `json:"show_email"`
	ShowPhoneNumber          bool      `json:"show_phone_number"`
	ShowPronouns             bool      `json:"show_pronouns"`
	PushNotificationsEnabled bool      `json:"push_notifications_enabled"`
	PushNotificationToken    string    `json:"push_notification_token"`
	MatchRemindersEnabled    bool      `json:"match_reminders_enabled"`
	UpdatedAt                time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt                time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Privacy) TableName() string {
	return "user_privacy"
}
