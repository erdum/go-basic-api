package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
	username string
}

func NewHomeController() *HomeController {
	return &HomeController{username: "Erdum"}
}

func (hc *HomeController) Home(context echo.Context) error {
	return context.String(http.StatusOK, hc.username)
}
