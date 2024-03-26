package main

import (
	"go-api/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New();
	homeController := controllers.NewHomeController()

	router.GET("/", homeController.Home)

	router.Logger.Fatal(router.Start(":8000"))
}
