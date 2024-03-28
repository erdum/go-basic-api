package main

import (
	"go-api/controllers"
	"go-api/models"
	"go-api/services/auth"

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

	// Services
	authService := auth.NewFirebaseAuth(db)

	// Inject services into the controllers
	authController := controllers.NewAuthController(authService)

	// router.GET("/users", userController.GetAllUsers)
	// router.POST("/users", userController.CreateUser)
	// router.GET("/users/:id", userController.GetUser)
	// router.PUT("/users/:id", userController.UpdateUser)
	// router.DELETE("/users/:id", userController.DeleteUser)

	router.GET("/login", authController.Login)

	router.Logger.Fatal(router.Start(":8000"))
}
