package controllers

import (
	"net/http"
	"go-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HomeController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *HomeController {
	return &HomeController{
		db: db,
	}
}

func (hc *HomeController) Home(context echo.Context) error {

	user := models.User{Name: "Erdum", Email: "erdumadnan@gmail.com"}
	result := hc.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return context.String(http.StatusOK, user.Email)
}
