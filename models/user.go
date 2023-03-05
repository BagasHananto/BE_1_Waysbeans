package models

import "time"

type User struct {
	Id       int                   `json:"id"`
	Username string                `json:"name" gorm:"type: varchar(255)"`
	Email    string                `json:"email" gorm:"type: varchar(255)"`
	Password string                `json:"password" gorm:"type: varchar(255)"`
	Profile  ProfileResponse       `json:"profile"`
	Products []ProductUserResponse `json:"products"`
	CreateAt time.Time             `json:"create_at"`
	UpdateAt time.Time             `json:"update_at"`
}

type UsersProfileResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
