package cartdto

type CartResponse struct {
	ProductID int `json:"product_id" gorm:"type: int"`
	OrderQty  int `json:"orderQuantity" gorm:"type: int"`
	UserID    int `json:"user_id"`
}
