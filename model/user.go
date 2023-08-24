package model

import "time"

type User struct {
	ID                string       `gorm:"primaryKey" json:"id"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	PreferredName     string       `json:"preferred_name"`
	Pronouns          string       `json:"pronouns"`
	Email             string       `gorm:"unique" json:"email"`
	ProfilePictureURL string       `json:"profile_picture_url"`
	Bio               string       `json:"bio"`
	Gender            string       `json:"gender"`
	Roles             []Role       `gorm:"-" json:"roles"`
	Privacy           Privacy      `gorm:"-" json:"privacy"`
	School            School       `gorm:"-" json:"school"`
	Verification      Verification `gorm:"-" json:"verification"`
	Connections       []Connection `gorm:"-" json:"connections"`
	UpdatedAt         time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt         time.Time    `gorm:"autoCreateTime" json:"created_at"`
}

func (User) TableName() string {
	return "user"
}

func (user User) String() string {
	return "(" + user.ID + ")" + " " + user.FirstName + " " + user.LastName
}
