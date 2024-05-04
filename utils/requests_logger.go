package utils

import (
	"os"
	"go-api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHTTPRequestsLogger(app *echo.Echo, logFilePath string) error {
	appConfig := config.GetConfig()
	if appConfig.Debug {
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} method=${method}, uri=${uri}, status=${status} ${latency_human}\n",
		}))
		return nil
	}
	
	logFile, err := os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "${time_rfc3339} ${remote_ip} ${status} ${method} ${uri} ${latency_human}\n",
	}))

	return nil
}
