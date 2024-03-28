package controllers

import (
	"go-api/services/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSevice auth.AuthService
}

type loginPayload struct {
	IdToken string `json:"idToken"`
}

func NewAuthController(authSevice auth.AuthService) *AuthController {
	return &AuthController{authSevice: authSevice}
}

func (ac *AuthController) Login(context echo.Context) error {
	payload := loginPayload{}
	context.Bind(&payload)

	data := ac.authSevice.AuthenticateWithThirdParty(payload.IdToken)
	return context.JSON(http.StatusOK, data)
}
