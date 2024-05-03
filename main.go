package main

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/models"
	"go-api/services/auth"
	"go-api/validators"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initialMigration() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	app := echo.New()

	db, err := initialMigration()
	if err != nil {
		app.Logger.Fatal(err)
	}

	appConfig, err := config.LoadConfig()
	if err != nil {
		app.Logger.Fatal(err)
	}

	app.Use(middleware.Logger())
	app.Use(middleware.RequestID())
	app.Validator = validators.NewDefaultValidator()

	// Services
	authService := auth.NewFirebaseAuth(db)

	// Inject services into the controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(db)

	app.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
	app.GET("/users/:id", userController.GetUser)
	app.PUT("/users/:id", userController.UpdateUser)
	app.DELETE("/users/:id", userController.DeleteUser)

	app.POST("/login", authController.Login)

	app.Logger.Fatal(app.Start(":" + appConfig.Port))
}
