package controllers

import (
	"net/http"
	"go-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetAllUsers(context echo.Context) error {
	users := []models.User{}
	uc.db.Find(&users)

	return context.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(context echo.Context) error {
	userId := context.Param("id")
	user := models.User{}
	uc.db.First(&user, userId)

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(context echo.Context) error {
	type UserPayload struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}
	payload := UserPayload{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	user := models.User{Name: payload.Name, Email: payload.Email}
	uc.db.Select("name", "email").Create(&user)

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(context echo.Context) error {
	userId := context.Param("id")
	user := models.User{}
	uc.db.First(&user, userId)

	type UserPayload struct {
		Name string `json:"name"`
	}
	payload := UserPayload{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if payload.Name != "" && payload.Name != user.Name {
		user.Name = payload.Name
		uc.db.Save(&user)
	}

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(context echo.Context) error {
	userId := context.Param("id")
	uc.db.Delete(&models.User{}, userId)

	return context.JSON(http.StatusOK, map[string]interface{}{"message": "User successfully deleted"})
}
