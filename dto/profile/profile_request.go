package profiledto

type ProfileRequest struct {
	Phone   string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: text"`
}

type UpdateProfileRequest struct {
	Phone   string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: text"`
}
