package profiledto

import "waysbeans/models"

type ProfileResponse struct {
	Id      int                         `json:"id" grom:"primary_key:auto_increment"`
	Phone   string                      `json:"phone" gorm:"type: varchar(255)"`
	Address string                      `json:"address" gorm:"type: text"`
	User    models.UsersProfileResponse `json:"user"`
}
