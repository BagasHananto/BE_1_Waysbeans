package handlers

import (
	"net/http"
	"strconv"
	cartdto "waysbeans/dto/cart"
	dto "waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// SETUP CONTROL STRUCT
type handlerCart struct {
	CartRepository repositories.CartRepository
}

// SETUP CONTROL FUNCTION
func ControlCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

// FUNCTION FIND CARTS
func (h *handlerCart) FindCarts(c echo.Context) error {
	carts, err := h.CartRepository.FindCart()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: carts})
}

// FUNCTION GET CART BY ID
func (h *handlerCart) GetCart(c echo.Context) error {
	// get url param ID
	id, _ := strconv.Atoi(c.Param("id"))

	// repository get profile
	var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertCart(cart)})
}

func (h *handlerCart) CreateCart(c echo.Context) error {
	productId, _ := strconv.Atoi(c.FormValue("product_id"))
	orderQty, _ := strconv.Atoi(c.FormValue("orderQuantity"))

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := cartdto.CartRequest{
		ProductID: productId,
		OrderQty:  orderQty,
		UserID:    int(userId),
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	cart := models.Cart{
		ProductID: request.ProductID,
		OrderQty:  request.OrderQty,
		UserID:    request.UserID,
	}

	cart, err = h.CartRepository.CreateCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	cart, _ = h.CartRepository.GetCart(cart.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertCart(cart)})
}

func convertCart(u models.Cart) cartdto.CartResponse {
	return cartdto.CartResponse{
		ProductID: u.ProductID,
		OrderQty:  u.OrderQty,
		UserID:    u.UserID,
	}
}
