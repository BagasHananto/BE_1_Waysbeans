package handlers

import (
	profiledto "waysbeans/dto/profile"
	dto "waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) FindProfile(c echo.Context) error {
	profile, err := h.ProfileRepository.FindProfile()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: profile})
}

func (h *handlerProfile) CreateProfile(c echo.Context) error {
	request := new(profiledto.ProfileRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	profile := models.Profile{
		Phone:   request.Phone,
		Address: request.Address,
		UserID:  int(userId),
	}

	data, err := h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseProfile(data)})
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		Id:      u.Id,
		Phone:   u.Phone,
		Address: u.Address,
		User:    u.User,
	}
}
