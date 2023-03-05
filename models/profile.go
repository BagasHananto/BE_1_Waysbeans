package models

type Profile struct {
	Id      int                  `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string               `json:"phone" gorm:"type: varchar(255)"`
	Address string               `json:"address" gorm:"type: text"`
	UserID  int                  `json:"user_id"`
	User    UsersProfileResponse `json:"user"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	Phone   string `json:"phone"`
	Address string `json:"address"`
	UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
