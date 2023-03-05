package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	productdto "waysbeans/dto/product"
	dto "waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(c echo.Context) error {
	var path_file = "http://localhost:5000/uploads/"
	products, err := h.ProductRepository.FindProducts()
	for i, p := range products {
		products[i].Photo = path_file + p.Photo
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: products})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	var path_file = "http://localhost:5000/uploads/"
	product, err := h.ProductRepository.GetProduct(id)

	product.Photo = path_file + product.Photo
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseProduct(product)})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("Data File", dataFile)

	price, _ := strconv.Atoi(c.FormValue("price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := productdto.ProductRequest{
		Name:   c.FormValue("name"),
		Price:  price,
		Desc:   c.FormValue("desc"),
		Stock:  stock,
		Photo:  dataFile,
		UserID: int(userId),
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	product := models.Product{
		Name:   request.Name,
		Price:  request.Price,
		Desc:   request.Desc,
		Stock:  request.Stock,
		Photo:  request.Photo,
		UserID: request.UserID,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseProduct(product)})
}

func convertResponseProduct(u models.Product) models.ProductResponse {
	return models.ProductResponse{
		Name:  u.Name,
		Price: u.Price,
		Desc:  u.Desc,
		Stock: u.Stock,
		Photo: u.Photo,
	}
}
