package controllers

import (
	"go-api/services/auth"
	"go-api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSevice auth.AuthService
}

func NewAuthController(authSevice auth.AuthService) *AuthController {
	return &AuthController{authSevice: authSevice}
}

func (ac *AuthController) Login(context echo.Context) error {
	payload := new(requests.LoginRequest)

	if err := context.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := context.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data := ac.authSevice.AuthenticateWithThirdParty(payload.IdToken)
	return context.NoContent(200)
	return context.JSON(http.StatusOK, data)
}
