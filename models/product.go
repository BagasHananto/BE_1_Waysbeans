package models

import "time"

type Product struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name      string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Price     int                  `json:"price" form:"price" gorm:"type: int"`
	Desc      string               `json:"desc" gorm:"type:text" form:"desc"`
	Photo     string               `json:"photo" form:"photo" gorm:"type: varchar(255)"`
	Stock     int                  `json:"stock" form:"stock"`
	UserID    int                  `json:"user_id" form:"user_id"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"create_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type ProductResponse struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Price  int                  `json:"price"`
	Desc   string               `json:"desc"`
	Stock  int                  `json:"stock"`
	Photo  string               `json:"photo"`
	UserID int                  `json:"user_id" form:"user_id"`
	User   UsersProfileResponse `json:"user"`
}

type ProductUserResponse struct {
	ID     int                  `json:"id"`
	Name   string               `json:"name"`
	Price  int                  `json:"price"`
	Desc   string               `json:"desc"`
	Stock  int                  `json:"stock"`
	Photo  string               `json:"photo"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
