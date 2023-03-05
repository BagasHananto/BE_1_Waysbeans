package cartdto

type CartRequest struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	OrderQty  int `json:"orderQuantity" gorm:"type: int" validate:"required"`
	UserID    int `json:"user_id"`
}
