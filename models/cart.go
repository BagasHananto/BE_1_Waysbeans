package models

import "time"

type Cart struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user"`
	ProductID int                  `json:"product_id" gorm:"type: int"`
	Product   ProductResponse      `json:"product"`
	OrderQty  int                  `json:"orderQuantity" gorm:"type: int"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type CartToUser struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	OrderQty  int `json:"orderQuantity" gorm:"type: int"`
	UserID    int `json:"user_id"`
}

type CartToProduct struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	OrderQty  int `json:"orderQuantity" gorm:"type: int"`
	UserID    int `json:"user_id"`
}

func (CartToUser) TableName() string {
	return "carts"
}

func (CartToProduct) TableName() string {
	return "carts"
}
