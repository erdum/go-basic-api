package controllers

import (
	"go-api/services/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSevice auth.AuthService
}

func NewAuthController(authSevice auth.AuthService) *AuthController {
	return &AuthController{authSevice: authSevice}
}

func (ac *AuthController) Login(context echo.Context) error {
	return nil
}
