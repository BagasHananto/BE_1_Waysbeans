package productdto

type ProductRequest struct {
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Desc   string `json:"desc" validate:"required"`
	Stock  int    `json:"stock" validate:"required"`
	Photo  string `json:"photo" validate:"required"`
	UserID int    `json:"user_id"`
}
