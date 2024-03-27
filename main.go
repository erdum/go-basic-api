package main

import (
	"go-api/controllers"
	"go-api/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
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
	homeController := controllers.New(db)

	router.GET("/", homeController.Home)

	router.Logger.Fatal(router.Start(":8000"))
}
