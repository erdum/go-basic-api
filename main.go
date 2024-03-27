package main

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initialMigration() *gorm.DB {
	db, error := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if error != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	return db
}

func main() {
	db := initialMigration()
	router := echo.New()
	userController := controllers.NewUserController(db)

	router.GET("/users", userController.GetAllUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	router.Logger.Fatal(router.Start(":8000"))
}
